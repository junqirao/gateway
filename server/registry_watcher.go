package server

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type serverConfigWatcher struct{}

func (serverConfigWatcher) OnChange(key string, value []byte) {
	registryConfigHandler(context.Background(), key, string(value))
}

func (serverConfigWatcher) OnClose(err error) {
	g.Log().Warningf(context.Background(), "server config watcher closed: %v", err)
}

type serverStatusWatcher struct{}

func (serverStatusWatcher) OnChange(key string, value []byte) {
	registryStatusHandler(context.Background(), key, string(value))
}

func (serverStatusWatcher) OnClose(err error) {
	g.Log().Warningf(context.Background(), "server status watcher closed: %v", err)
}
