package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/junqirao/gateway/component/registry"
	"github.com/junqirao/gateway/lib/response"
	"github.com/junqirao/gateway/model"
	"testing"
)

var (
	srd = model.NodeRegisterData{
		ServerGroup: &model.ServerGroup{
			ServerName:  "",
			ServiceName: "v1",
			GroupName:   "api",
		},
		Node: &model.NodeInfo{
			Name:     "test",
			Protocol: "http",
			Host:     "127.0.0.1",
			Port:     8989,
			Meta: map[string]interface{}{
				"test1": 123,
				"test2": map[string]interface{}{
					"a": 1,
					"b": 2,
				},
				"test3": []string{"a", "b"},
			},
		},
	}
)

func runDefaultServer() {
	server := g.Server()
	server.SetPort(8000)
	server.Run()
}

func runTestServer() {
	server := g.Server("test")
	server.SetPort(srd.Node.Port)
	server.Group("/", func(group *ghttp.RouterGroup) {
		group.GET(buildHandler("/a", "123", 200))
		group.GET(buildHandler("/b", "456", 200))
		group.ALL("/callback", func(r *ghttp.Request) {
			fmt.Println("recv callback body: " + string(r.GetBody()))
		})
	})
	server.Run()
}

func buildHandler(path string, responseData interface{}, status int) (string, ghttp.HandlerFunc) {
	return path, func(r *ghttp.Request) {
		r.Response.WriteHeader(status)
		r.Response.WriteJson(response.JSON{
			Data:    responseData,
			Message: fmt.Sprintf("router: %v", path),
		})
	}
}

func TestRegister(t *testing.T) {
	go runTestServer()

	// SubmitChanges(&srd)
	runDefaultServer()
}

func TestInit(t *testing.T) {
	marshal, _ := json.Marshal(srd)
	err := registry.Instance().Set(context.TODO(), fmt.Sprintf("%s%s", registryKey, "test"), string(marshal))
	if err != nil {
		t.Fatal(err)
		return
	}
	Init()
	groups.Range(func(key, value any) bool {
		fmt.Println("key = ", key)
		group := value.(*Group)
		fmt.Println("group.Name = ", group.Name)
		group.services.Range(func(k, v any) bool {
			service := v.(*Service)
			fmt.Println("service.Name = ", service.Name)
			fmt.Println("service.lb = ", service.lb)
			fmt.Println("service.nodes.length = ", len(service.nodes))
			for i, node := range service.nodes {
				marshal, _ := json.Marshal(node)
				fmt.Println("service.nodes[", i, "] = ", string(marshal))
			}
			return true
		})
		return true
	})
}
