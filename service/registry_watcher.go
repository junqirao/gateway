package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type serviceRegistryWatcher struct{}

// OnChange ...
func (serviceRegistryWatcher) OnChange(key string, value []byte) {
	registryConfigHandler(context.Background(), key, string(value))
}

// OnClose ...
func (serviceRegistryWatcher) OnClose(err error) {
	g.Log().Warningf(context.Background(), "service config watcher closed: %v", err)
}
