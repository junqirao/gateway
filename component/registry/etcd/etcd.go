package etcd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/junqirao/gateway/component/registry/event"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	Ins *instance
)

type Config = clientv3.Config

func init() {
	ctx := context.Background()
	v, err := g.Cfg().Get(ctx, "registry.etcd")
	if err != nil {
		glog.DefaultLogger().Panicf(ctx, "etcd read config failed: %v", err)
		return
	}

	cfg := new(Config)
	err = v.Struct(&cfg)
	if err != nil {
		glog.DefaultLogger().Panicf(ctx, "etcd parse config failed: %v", err)
		return
	}

	cli, err := clientv3.New(*cfg)
	if err != nil {
		glog.DefaultLogger().Panicf(ctx, "etcd client init failed: %v", err)
		return
	}
	Ins = &instance{cli: cli, watch: make(map[string]clientv3.WatchChan)}
}

type instance struct {
	cli   *clientv3.Client
	watch map[string]clientv3.WatchChan
}

func (i *instance) watchHandler(watchChan clientv3.WatchChan, handler event.Handler) {
	for {
		select {
		case resp := <-watchChan:
			if resp.Canceled {
				return
			}
			for _, ev := range resp.Events {
				handler(string(ev.Kv.Key), ev.Kv.Value)
			}
		default:

		}
	}
}

func (i *instance) Subscribe(ctx context.Context, key string, handler event.Handler) {
	go i.watchHandler(i.cli.Watch(ctx, key), handler)
	return
}
