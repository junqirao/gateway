package registry

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gateway/component/registry/etcd"
	"github.com/junqirao/gateway/component/registry/event"
)

type Registry interface {
	Subscribe(ctx context.Context, key string, handler event.Watcher)
	Get(ctx context.Context, key string) (m map[string]string, err error)
	Set(ctx context.Context, key string, value interface{}) (err error)
	Delete(ctx context.Context, key string) (err error)
}

type emptyRegistry struct{}

func (emptyRegistry) Subscribe(ctx context.Context, key string, handler event.Watcher) {
	panic("failed to initialize registry")
}

func (emptyRegistry) Get(ctx context.Context, key string) (m map[string]string, err error) {
	return
}

func (emptyRegistry) Set(ctx context.Context, key string, value interface{}) (err error) {
	return
}

func (emptyRegistry) Delete(ctx context.Context, key string) (err error) {
	return
}

const (
	TypeEtcd = "etcd"
)

var (
	ins Registry
)

// Instance of registry
func Instance() Registry {
	if ins != nil {
		return ins
	}

	v, err := g.Cfg().Get(context.TODO(), "registry.type", TypeEtcd)
	if err != nil {
		return &emptyRegistry{}
	}

	switch v.String() {
	case TypeEtcd:
		ins = etcd.Ins
		return ins
	}
	return &emptyRegistry{}
}
