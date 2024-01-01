package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type nodeRegistryWatcher struct{}

// OnChange ...
func (nodeRegistryWatcher) OnChange(key string, value []byte) {
	nodeRegistryHandler(context.Background(), key, string(value))
}

// OnClose ...
func (nodeRegistryWatcher) OnClose(err error) {
	g.Log().Warningf(context.Background(), "service config watcher closed: %v", err)
}
