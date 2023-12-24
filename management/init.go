package management

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

const (
	serverIdentity = "management.http.server"
)

func Init() {
	ctx := context.Background()
	v, err := g.Cfg().Get(ctx, "management")
	if err != nil {
		g.Log().Warningf(ctx, "failed to get management config: %v", err)
		return
	}

	cfg := new(Config)
	if err = v.Struct(&cfg); err != nil {
		g.Log().Warningf(ctx, "failed to load management config: %v", err)
		return
	}

	go setup(ctx, cfg)
}

func setup(ctx context.Context, cfg *Config) {
	if !cfg.Enabled {
		g.Log().Info(ctx, "module management disabled.")
		return
	}

	debug := false
	v, err := g.Cfg().Get(ctx, "debug", false)
	if err == nil {
		debug = v.Bool()
	}

	sName := fmt.Sprintf("%s.%s", serverIdentity, guid.S())
	server := g.Server(sName)
	if cfg.Address != "" {
		server.SetAddr(cfg.Address)
	}
	if cfg.Port > 0 && cfg.Port < 65535 {
		server.SetPort(cfg.Port)
	}
	setupRouter(server.Group("/management"))

	server.SetDumpRouterMap(debug)
	g.Log().Infof(ctx, "%s loaded on %s:%d", serverIdentity, cfg.Address, cfg.Port)
	server.Run()
}
