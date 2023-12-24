package etcd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/junqirao/gateway/component/registry/event"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
)

var (
	Ins *instance
)

type Config = clientv3.Config

func init() {
	ctx := context.Background()
	v, err := g.Cfg().Get(ctx, "registry.etcd")
	if err != nil {
		g.Log().Panicf(ctx, "etcd read config failed: %v", err)
		return
	}

	cfg := new(Config)
	err = v.Struct(&cfg)
	if err != nil {
		g.Log().Panicf(ctx, "etcd parse config failed: %v", err)
		return
	}

	cli, err := clientv3.New(*cfg)
	if err != nil {
		g.Log().Panicf(ctx, "etcd client init failed: %v", err)
		return
	}
	Ins = &instance{cli: cli}
}

type instance struct {
	cli *clientv3.Client
}

// Subscribe a key
func (i *instance) Subscribe(ctx context.Context, key string, handler event.Watcher) {
	var ops []clientv3.OpOption
	hasPrefix := strings.HasSuffix(key, "/")
	if hasPrefix {
		ops = append(ops, clientv3.WithPrefix())
	}
	go func() {
		g.Log().Infof(ctx, "subscribe value change of key(prefix:%v): %v", hasPrefix, key)
		watchChan := i.cli.Watch(ctx, key, ops...)
		for {
			select {
			case resp := <-watchChan:
				if resp.Canceled {
					handler.OnClose(resp.Err())
					return
				}
				for _, ev := range resp.Events {
					k := string(ev.Kv.Key)
					if hasPrefix {
						k = strings.TrimPrefix(k, key)
					}
					handler.OnChange(k, ev.Kv.Value)
				}
			case <-ctx.Done():
				handler.OnClose(nil)
				return
			}
		}
	}()
	return
}

// Set key and value
func (i *instance) Set(ctx context.Context, key string, value interface{}) (err error) {
	_, err = i.cli.Put(ctx, key, gconv.String(value))
	return
}

// Get by key
func (i *instance) Get(ctx context.Context, key string) (m map[string]string, err error) {
	var ops []clientv3.OpOption
	hasPrefix := strings.HasSuffix(key, "/")
	if hasPrefix {
		ops = append(ops, clientv3.WithPrefix())
	}

	resp, err := i.cli.Get(ctx, key, ops...)
	if err != nil {
		return
	}
	m = make(map[string]string)
	for _, value := range resp.Kvs {
		k := string(value.Key)
		if hasPrefix {
			k = strings.TrimPrefix(k, key)
		}
		m[k] = string(value.Value)
	}
	return
}

// Delete by key
func (i *instance) Delete(ctx context.Context, key string) (err error) {
	_, err = i.cli.Delete(ctx, key)
	return
}
