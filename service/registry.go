package service

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gateway/component/registry"
	"github.com/junqirao/gateway/model"
)

var (
	registryKey = g.Cfg().MustGet(context.Background(), "registry.identity", "local.undefined").String() + "/service/registry/"
)

// registryHandler ...
func registryHandler(ctx context.Context, name, cfgStr string, create ...bool) {
	// name : node_name
	var err error
	if cfgStr == "" {
		return
	}

	sc := new(model.NodeRegisterData)
	if err = json.Unmarshal([]byte(cfgStr), &sc); err != nil {
		g.Log().Warningf(ctx, "load service registry data [%s](value=%s) failed: %v", name, cfgStr, err)
		return
	}

	if sc.ServerGroup == nil {
		return
	}

	nodeOp := sc.Operation
	if len(create) > 0 && create[0] {
		nodeOp = registry.OperationUpdate
	}
	CreateOrGetGroup(sc.ServerGroup.ServerName, sc.ServerGroup.GroupName).
		SubmitChanges(sc.ServerGroup).
		UpdateOrCreateNode(sc.Node, nodeOp)
	return
}
