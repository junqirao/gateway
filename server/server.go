package server

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gateway/model"
	"github.com/junqirao/gateway/proxy"
)

// Register service
func Register(reg *model.ServiceRegisterData) {
	h := proxy.NewHandler(reg)
	g.Server(reg.ServerName).BindHandler(reg.RouterPattern(), h.Proxy)
}
