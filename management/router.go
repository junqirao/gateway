package management

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/junqirao/gateway/management/api"
)

func setupRouter(group *ghttp.RouterGroup) {
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
