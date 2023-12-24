package server

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gateway/model"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	Init()
	time.Sleep(time.Second)
}

func TestStopStart(t *testing.T) {
	name := "test"
	srv := g.Server(name)
	config := &model.ServerConfig{
		Properties: &model.ServerProperties{
			Address: ":7777",
		},
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
