package server

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/junqirao/gateway/component/registry"
	"github.com/junqirao/gateway/model"
)

var (
	registryKey = g.Cfg().MustGet(context.Background(), "registry.identity", "local.undefined").String() + "/server/config/"
)

func loadAllFromRegistry(ctx context.Context) {
	m, err := registry.Instance().Get(ctx, registryKey)
	if err != nil {
		g.Log().Warningf(ctx, "load server config from registry failed: %v", err)
		return
	}

	for name, s := range m {
		err = setupServer(ctx, name, s)
		if err != nil {
			g.Log().Warningf(ctx, "setup server failed: %v", err)
		}
	}
}

func setupServer(ctx context.Context, name string, cfgStr string) (err error) {
	sc := new(model.ServerConfig)
	if err = json.Unmarshal([]byte(cfgStr), &sc); err != nil {
		g.Log().Warningf(ctx, "load server config [%s](value=%s) failed: %v", name, cfgStr, err)
		return
	}

	srv := g.Server(name)

	if sc.ServerConfig != nil {
		if err = srv.SetConfig(*sc.ServerConfig); err != nil {
			g.Log().Warningf(ctx, "set server config [%s] failed: %v", name, err)
			return
		}
	}

	if sc.Up {
		if srv.Status() == ghttp.ServerStatusStopped {
			go srv.Run()
			g.Log().Infof(ctx, "server running: %s", name)
		}
	} else {
		if srv.Status() == ghttp.ServerStatusRunning {
			err = srv.Shutdown()
			g.Log().Infof(ctx, "server stopped: %s", name)
		}
	}

	return
}

type serverConfigWatcher struct{}

func (serverConfigWatcher) OnChange(key string, value []byte) {
	_ = setupServer(context.Background(), key, string(value))
}

func (serverConfigWatcher) OnClose(err error) {
	g.Log().Warningf(context.Background(), "server config watcher closed: %v", err)
}
