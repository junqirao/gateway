package service

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gateway/model"
)

var (
	registryKey = g.Cfg().MustGet(context.Background(), "registry.identity", "local.undefined").String() + "/service/registry/"
)

func registryConfigHandler(ctx context.Context, name, cfgStr string) {
	var err error
	if cfgStr == "" {
		// delete
		// if err = DeRegister(ctx, name); err != nil {
		// 	g.Log().Warningf(ctx, "deregister service(%s) failed: %v", name, err)
		// } else {
		// 	g.Log().Infof(ctx, "deregister service(%s) success", name)
		// }
		// return
	}

	sc := new(model.ServiceRegisterData)
	if err = json.Unmarshal([]byte(cfgStr), &sc); err != nil {
		g.Log().Warningf(ctx, "load service config [%s](value=%s) failed: %v", name, cfgStr, err)
		return
	}

	// if err = UpdateConfigOrRegister(ctx, name, sc); err != nil {
	// 	g.Log().Warningf(ctx, "update or register service(%s) failed: %v", name, err)
	// } else {
	// 	g.Log().Infof(ctx, "update or register service(%s) success", name)
	// }
	return
}
