package service

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gateway/model"
	"strings"
)

// nodeRegistryHandler ...
func nodeRegistryHandler(ctx context.Context, name, cfgStr string) {
	// name : server.group.service.node_name
	parts := strings.Split(name, ".")
	if len(parts) < 4 {
		return
	}

	serverName := parts[0]
	groupName := parts[1]
	serviceName := parts[2]
	nodeName := parts[3]
	group := CreateOrGetGroup(serverName, groupName)

	var err error
	if cfgStr == "" {
		// unregister
		if service, ok := group.findService(serviceName); ok {
			service.RemoveNode(nodeName)
		}
		return
	}

	sc := new(model.NodeRegisterData)
	if err = json.Unmarshal([]byte(cfgStr), &sc); err != nil {
		g.Log().Warningf(ctx, "load node registry data [%s](value=%s) failed: %v", name, cfgStr, err)
		return
	}

	if sc.ServerGroup == nil {
		return
	}

	group.CreateOrGetService(sc.ServerGroup).UpdateOrCreateNode(sc.Node)
	return
}
