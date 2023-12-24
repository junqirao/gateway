package server

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gateway/model"
	"github.com/junqirao/gateway/proxy"
)

// RegisterService service
func RegisterService(reg *model.ServiceRegisterData) {
	h := proxy.NewHandler(reg)
	g.Server(reg.ServerName).BindHandler(reg.RouterPattern(), h.Proxy)
}
