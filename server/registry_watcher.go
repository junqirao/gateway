package server

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type serverConfigWatcher struct{}

// OnChange ...
func (serverConfigWatcher) OnChange(key string, value []byte) {
	registryConfigHandler(context.Background(), key, string(value))
}

// OnClose ...
func (serverConfigWatcher) OnClose(err error) {
	g.Log().Warningf(context.Background(), "server config watcher closed: %v", err)
}
