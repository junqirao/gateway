package registry

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gateway/component/registry/etcd"
	"github.com/junqirao/gateway/component/registry/event"
)

type Registry interface {
	Subscribe(ctx context.Context, key string, handler event.Handler)
}

type emptyRegistry struct{}

func (emptyRegistry) Subscribe(ctx context.Context, key string, handler event.Handler) {
	panic("failed to initialize registry")
}

const (
	TypeEtcd = "etcd"
)

var (
	ins Registry
)

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
