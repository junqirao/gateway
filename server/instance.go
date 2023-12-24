package server

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/junqirao/gateway/lib/response"
	"github.com/junqirao/gateway/model"
	"sync"
)

var (
	insMap = sync.Map{} // server_name:instance
)

// GetInstance of server
func GetInstance(name string) (i *Instance, ok bool) {
	v, ok := insMap.Load(name)
	if ok && v != nil {
		i, ok = v.(*Instance)
	}
	return
}

// Register of server instance
func Register(ctx context.Context, name string, config *model.ServerConfig) (i *Instance, err error) {
	_, ok := insMap.Load(name)
	if ok {
		err = response.ErrorResourceAlreadyExists.WithDetail(name)
		return
	}

	i = NewInstance(name, config)
	insMap.Store(name, i)
	if config.Enabled {
		err = i.Start(ctx)
	}
	return
}

// DeRegister of server instance
func DeRegister(ctx context.Context, name string) (err error) {
	v, ok := insMap.LoadAndDelete(name)
	if !ok {
		err = response.ErrorResourceNotFound.WithDetail(name)
		return
	}

	ins := v.(*Instance)
	return ins.Stop(ctx)
}

// UpdateConfigOrRegister of server instance
func UpdateConfigOrRegister(ctx context.Context, name string, config *model.ServerConfig) (err error) {
	instance, ok := GetInstance(name)
	if !ok {
		instance, err = Register(ctx, name, config)
		if err != nil {
			g.Log().Warningf(ctx, "server register [%s] failed: %v", name, err)
			return
		}
	} else {
		if err = instance.UpdateConfig(ctx, config); err != nil {
			g.Log().Warningf(ctx, "set server config [%s] failed: %v", name, err)
			return
		}
	}
	return
}

// Instance of server
type Instance struct {
	id   string // change on start
	name string
	srv  *ghttp.Server
	cfg  *model.ServerConfig
}

// NewInstance constructor
func NewInstance(name string, cfg *model.ServerConfig) *Instance {
	return &Instance{
		name: name,
		cfg:  cfg,
	}
}

// Start http server
func (i *Instance) Start(ctx context.Context) (err error) {
	if i.srv == nil {
		i.id = guid.S()
		serverName := fmt.Sprintf("%s_%s", i.name, i.id)
		i.srv = g.Server(serverName)
		if err = i.srv.SetConfig(i.cfg.C(serverName)); err != nil {
			return
		}
	}
	if i.srv.Status() == ghttp.ServerStatusRunning {
		return
	}
	if err = i.srv.Start(); err != nil {
		g.Log().Warningf(ctx, "failed to start server: %v", err)
	} else {
		g.Log().Infof(ctx, "server started: %v", i.name)
	}
	return
}

// Stop http server
func (i *Instance) Stop(ctx context.Context) (err error) {
	if i.srv == nil || i.srv.Status() == ghttp.ServerStatusStopped {
		return
	}
	if err = i.srv.Shutdown(); err != nil {
		g.Log().Warningf(ctx, "failed to stop server: %v", err)
	} else {
		i.srv = nil
		g.Log().Infof(ctx, "server stopped: %v", i.name)
	}
	return
}

// UpdateConfig of http server
func (i *Instance) UpdateConfig(ctx context.Context, cfg *model.ServerConfig) (err error) {
	enabledChanges := i.cfg.Enabled != cfg.Enabled
	propertiesChanges := i.cfg.Properties != nil

	if propertiesChanges {
		i.cfg.Properties = cfg.Properties
	}

	if enabledChanges {
		i.cfg.Enabled = cfg.Enabled
		if i.cfg.Enabled {
			err = i.Start(ctx)
		} else {
			err = i.Stop(ctx)
		}
	}
	return
}

// Srv returns http server
func (i *Instance) Srv() *ghttp.Server {
	return i.srv
}
