package management

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/junqirao/gateway/management/api"
	"github.com/junqirao/gateway/management/authorization"
)

func setupRouter(ctx context.Context, server *ghttp.Server, cfg *Config) {
	group := server.Group("/management")
	middlewares := []ghttp.HandlerFunc{ghttp.MiddlewareCORS}
	if cfg.Secret != "" {
		middlewares = append(middlewares, authorization.VerifySignature(cfg.Secret))
	} else {
		g.Log().Warningf(ctx, "management check signature is disalbed, check management.secret in config.yml")
	}
	if cfg.IpWhitelist != "" {
		middlewares = append(middlewares, authorization.CheckIpWhitelist(cfg.IpWhitelist))
	}
	group.Middleware(middlewares...)

	serverManagementRouters(group)
}

func serverManagementRouters(group *ghttp.RouterGroup) {
	group.Group("/server", func(group *ghttp.RouterGroup) {
		// info
		group.GET("/list", api.ServerMgr.List)
		group.GET("/{name}", api.ServerMgr.Get)
		group.DELETE("/{name}", api.ServerMgr.Delete)
		// config
		group.POST("/config/{name}", api.ServerMgr.UpdateConfig)
	})
}
