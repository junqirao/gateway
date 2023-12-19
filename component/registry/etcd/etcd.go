package etcd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	cli *clientv3.Client
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

	cli, err = clientv3.New(*cfg)
	if err != nil {
		glog.DefaultLogger().Panicf(ctx, "etcd client init failed: %v", err)
		return
	}
}
