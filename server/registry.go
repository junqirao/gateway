package server

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gateway/model"
)

var (
	configRegistryKey = g.Cfg().MustGet(context.Background(), "registry.identity", "local.undefined").String() + "/server/config/"
	statusRegistryKey = g.Cfg().MustGet(context.Background(), "registry.identity", "local.undefined").String() + "/server/status/"
)

func registryStatusHandler(ctx context.Context, name, str string) {
	if str == "" {
		return
	}
	var err error

	st := new(model.ServerStatus)
	if err = json.Unmarshal([]byte(str), &st); err != nil {
		g.Log().Warningf(ctx, "load server status [%s](value=%s) failed: %v", name, str, err)
		return
	}

	ins, ok := GetInstance(name)
	if !ok {
		return
	}

	if err = ins.UpdateStatus(ctx, st); err != nil {
		g.Log().Warningf(ctx, "update server(%s) status failed: %v", name, err)
	} else {
		g.Log().Infof(ctx, "update server(%s) status success: enabled=%v", name, st.Enabled)
	}
	return
}

func registryConfigHandler(ctx context.Context, name, cfgStr string) {
	var err error
	if cfgStr == "" {
		// delete
		if err = DeRegister(ctx, name); err != nil {
			g.Log().Warningf(ctx, "deregister server(%s) failed: %v", name, err)
		} else {
			g.Log().Infof(ctx, "deregister server(%s) success", name)
		}
		return
	}

	sc := new(model.ServerConfig)
	if err = json.Unmarshal([]byte(cfgStr), &sc); err != nil {
		g.Log().Warningf(ctx, "load server config [%s](value=%s) failed: %v", name, cfgStr, err)
		return
	}

	if err = UpdateConfigOrRegister(ctx, name, sc); err != nil {
		g.Log().Warningf(ctx, "update or register server(%s) failed: %v", name, err)
	} else {
		g.Log().Infof(ctx, "update or register server(%s) success", name)
	}
	return
}
