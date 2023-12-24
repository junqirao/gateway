package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/junqirao/gateway/lib/response"
	"github.com/junqirao/gateway/server"
)

var ServerMgr = &serverMgrApi{
	logger: g.Log("management"),
}

type serverMgrApi struct {
	logger *glog.Logger
}

// List server info
// @Tags ServerMgr
// @Summary list server info
// @Success 200 {object} response.JSON{} ""
// @Router /management/server/list[get]
func (a *serverMgrApi) List(r *ghttp.Request) {
	ctx := r.Context()
	data, err := server.ListServerInfo(ctx)
	if err != nil {
		response.Error(r, err)
		return
	}
	response.Data(r, data)
}

// Get
// @Tags ServerMgr
// @Summary get server info by name
// @Param name path string true "server name"
// @Success 200 {object} response.JSON{} ""
// @Router /management/server/{name} [get]
func (a *serverMgrApi) Get(r *ghttp.Request) {
	ctx := r.Context()
	req := new(ServerGetReq)
	err := r.Parse(&req)
	if err != nil {
		a.logger.Warningf(ctx, "serverMgrApi.GetServerInfo params parse error : %s", err)
		response.Error(r, response.ErrorInvalidParams, response.WithMessage(err.Error()))
		return
	}

	data, err := server.GetServerInfo(ctx, req.Name)
	if err != nil {
		response.Error(r, err)
		return
	}
	response.Data(r, data)
}

// Delete
// @Tags ServerMgr
// @Summary delete server
// @Param name path string true "server name"
// @Success 200 {object} response.JSON{} ""
// @Router /management/server/{name} [delete]
func (a *serverMgrApi) Delete(r *ghttp.Request) {
	ctx := r.Context()
	req := new(ServerDeleteReq)
	err := r.Parse(&req)
	if err != nil {
		a.logger.Warningf(ctx, "serverMgrApi.DeleteConfig params parse error : %s", err)
		response.Error(r, response.ErrorInvalidParams, response.WithMessage(err.Error()))
		return
	}

	data, err := server.DeleteConfig(ctx, req.Name)
	if err != nil {
		response.Error(r, err)
		return
	}
	response.Data(r, data)
}

// UpdateConfig
// @Tags ServerMgr
// @Summary update server config
// @Param name path string true "server name"
// @Param entity body ServerConfigUpdateReq true "body"
// @Success 200 {object} response.JSON{} ""
// @Router /management/server/config/{name} [post]
func (a *serverMgrApi) UpdateConfig(r *ghttp.Request) {
	ctx := r.Context()
	req := new(ServerUpdateConfigReq)
	err := r.Parse(&req)
	if err != nil {
		a.logger.Warningf(ctx, "serverMgrApi.UpdateConfig params parse error : %s", err)
		response.Error(r, response.ErrorInvalidParams, response.WithMessage(err.Error()))
		return
	}

	data, err := server.SetConfig(ctx, req.Name, req.ServerConfig)
	if err != nil {
		response.Error(r, err)
		return
	}
	response.Data(r, data)
}
