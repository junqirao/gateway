package service

import (
	"github.com/junqirao/gateway/model"
	"github.com/junqirao/gateway/proxy"
	"github.com/junqirao/gateway/server"
)

// Register service
func Register(reg *model.ServiceRegisterData) {
	h := proxy.NewHandler(reg)
	instance, ok := server.GetInstance(reg.ServerName)
	if !ok || instance.Srv() == nil {
		return
	}
	instance.Srv().BindHandler(reg.RouterPattern(), h.Proxy)
}
