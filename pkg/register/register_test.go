package register

import (
	"context"
	"github.com/junqirao/gateway/pkg/model"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
)

var (
	identity = "local.test"
	etcdCfg  = clientv3.Config{
		Endpoints: []string{"10.11.21.50:12379", "10.11.21.50:12381", "10.11.21.50:12383"},
	}
	serverName = "test_server_1"
	nrd        = model.NodeRegisterData{
		ServerGroup: model.ServerGroup{
			ServerName:  serverName,
			ServiceName: "v1",
			GroupName:   "api",
		},
		Node: model.NodeInfo{
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
	register Register
)

func newRegister(t *testing.T) {
	var err error
	register, err = New(TypeEtcd, func() interface{} {
		return etcdCfg
	}, WithLogger(nil), WithIdentity(identity))
	if err != nil {
		t.Fatal(err)
		return
	}
}

func TestNew(t *testing.T) {
	newRegister(t)
}

func TestRegister(t *testing.T) {
	newRegister(t)
	if err := register.Register(context.Background(), &nrd); err != nil {
		t.Fatal(err)
	}
}

func TestUnRegister(t *testing.T) {
	newRegister(t)
	if err := register.Unregister(context.Background(), &nrd); err != nil {
		t.Fatal(err)
	}
}
