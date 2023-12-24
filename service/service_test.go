package service

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/junqirao/gateway/lib/response"
	"github.com/junqirao/gateway/model"
	"testing"
)

var (
	srd = model.ServiceRegisterData{
		ServerGroup: model.ServerGroup{
			ServerName: "",
			Group: &model.GroupInfo{
				Name: "/api/v1",
			},
		},
		Service: model.ServiceInfo{
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
	server.SetPort(srd.Service.Port)
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

	Register(&srd)
	runDefaultServer()
}
