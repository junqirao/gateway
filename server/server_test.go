package server

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/junqirao/gateway/lib/response"
	"github.com/junqirao/gateway/model"
	"testing"
	"time"
)

var (
	srd = model.ServiceRegisterData{
		ServerName: "", // default
		Group: model.GroupInfo{
			Name: "/api/v1",
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
		Callback: "http://127.0.0.1:8989/callback",
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

	RegisterService(&srd)
	runDefaultServer()
}

func TestInit(t *testing.T) {
	Init()
	time.Sleep(time.Second)
}

func TestStopStart(t *testing.T) {
	name := "test"
	srv := g.Server(name)
	config := &model.ServerConfig{
		Address: ":7777",
	}
	err := srv.SetConfig(config.C(name))
	if err != nil {
		t.Fatal(err)
		return
	}

	instance := NewInstance(name, config)
	for i := 0; i < 3; i++ {
		t.Log("i = ", i)
		if err = instance.Start(context.Background()); err != nil {
			t.Fatal(err)
			return
		}

		if err = instance.Stop(context.Background()); err != nil {
			t.Fatal(err)
			return
		}
	}
}

func TestStopStartRaw(t *testing.T) {
	server := g.Server("test")

	for i := 0; i < 3; i++ {
		t.Log("i = ", i)
		err := server.Start()
		if err != nil {
			t.Fatal(err)
			return
		}
		err = server.Shutdown()
		if err != nil {
			t.Fatal(err)
			return
		}
	}
}
