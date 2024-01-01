package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/junqirao/gateway/component/registry"
	"github.com/junqirao/gateway/lib/response"
	"github.com/junqirao/gateway/management"
	"github.com/junqirao/gateway/model"
	"github.com/junqirao/gateway/server"
	"github.com/junqirao/gateway/service"
	"testing"
)

var (
	serverName = "test_server_1"
	sc         = model.ServerConfig{
		Enabled: true,
		Properties: &model.ServerProperties{
			Address: ":80",
		},
	}
	nrd = model.NodeRegisterData{
		ServerGroup: &model.ServerGroup{
			ServerName:  serverName,
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
	srv := g.Server()
	srv.SetPort(8000)
	srv.Run()
}

func runTestServer(t *testing.T) {
	srv := g.Server("test")
	srv.SetPort(nrd.Node.Port)
	srv.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(func(r *ghttp.Request) {
			fmt.Println("[test_server.middleware] request url :", r.Request.URL.String())
			r.Middleware.Next()
		})
		group.GET(buildHandler("/a", "123", 200))
		group.GET(buildHandler("/b", "456", 200))
		group.ALL("/callback", func(r *ghttp.Request) {
			fmt.Println("recv callback body: " + string(r.GetBody()))
		})
	})
	t.Log("test server started at: ", nrd.Node.Port)
	srv.Run()
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
	go runTestServer(t)

	// SubmitChanges(&nrd)
	runDefaultServer()
}

func TestInit(t *testing.T) {
	server.Init()
	service.Init()
	management.Init()

	go runTestServer(t)

	graceExit()
	// groups.Range(func(key, value any) bool {
	// 	fmt.Println("key = ", key)
	// 	group := value.(*Group)
	// 	fmt.Println("group.Name = ", group.Name)
	// 	group.services.Range(func(k, v any) bool {
	// 		service := v.(*findService)
	// 		fmt.Println("service.Name = ", service.Name)
	// 		fmt.Println("service.lb = ", service.lb)
	// 		fmt.Println("service.nodes.length = ", len(service.nodes))
	// 		for i, node := range service.nodes {
	// 			marshal, _ := json.Marshal(node)
	// 			fmt.Println("service.nodes[", i, "] = ", string(marshal))
	// 		}
	// 		return true
	// 	})
	// 	return true
	// })
}

func TestRegisterNode(t *testing.T) {
	marshal, _ := json.Marshal(nrd)
	err := registry.Instance().Set(context.TODO(), registry.NodeRegKey(nrd.RegistryKey()), string(marshal))
	if err != nil {
		t.Fatal(err)
		return
	}
}

func TestUnRegisterNode(t *testing.T) {
	err := registry.Instance().Delete(context.TODO(), registry.NodeRegKey(nrd.RegistryKey()))
	if err != nil {
		t.Fatal(err)
		return
	}
}
