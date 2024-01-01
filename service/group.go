package service

import (
	"context"
	"fmt"
	"github.com/junqirao/gateway/model"
	"github.com/junqirao/gateway/proxy/balancer"
	"github.com/junqirao/gateway/proxy/node"
	"strings"
	"sync"
	"sync/atomic"
)

var (
	groups = sync.Map{} // key: group_name, value: *Group
)

// Group ...
type Group struct {
	Name     string
	services sync.Map // server_name.service_name : Service
	count    atomic.Int32
}

// CreateOrGetGroup ...
func CreateOrGetGroup(serverName, groupName string) *Group {
	if g, ok := findGroup(serverName, groupName); ok {
		return g
	}
	g := new(Group)
	g.Name = groupName
	groups.Store(serverGroupName(serverName, groupName), g)
	logger.Infof(context.TODO(), "create group: %s", g.Name)
	return g
}

func findGroup(serverName, groupName string) (*Group, bool) {
	if v, ok := groups.Load(serverGroupName(serverName, groupName)); ok {
		if g, ok := v.(*Group); ok {
			return g, ok
		}
	}
	return nil, false
}

func serverGroupName(s, g string) string {
	if s == "" {
		return g
	}
	var sb strings.Builder
	sb.WriteString(s)
	sb.WriteByte('.')
	sb.WriteString(g)
	return sb.String()
}

// CreateOrGetService ...
func (g *Group) CreateOrGetService(sg *model.ServerGroup) *Service {
	if service, ok := g.findService(sg.ServiceName); ok {
		return service
	}

	service := &Service{
		group: g,
		nodes: make([]*node.Node, 0),
		lb:    balancer.New(sg.LB),
		Name:  sg.ServiceName,
	}
	g.services.Store(sg.ServiceName, service)
	g.count.Add(1)
	g.log("create service: %s", service.Name)
	return service
}

// Service ...
func (g *Group) findService(serviceName string) (*Service, bool) {
	if v, ok := g.services.Load(serviceName); ok {
		if service, ok := v.(*Service); ok {
			return service, ok
		}
	}
	return nil, false
}

func (g *Group) removeService(serviceName string) {
	g.services.Delete(serviceName)
	g.count.Add(-1)
	g.log("remove service: %s", serviceName)

	if g.count.Load() == 0 {
		groups.Delete(g.Name)
		logger.Infof(context.TODO(), "delete group: %s", g.Name)
	}
}

func (g *Group) log(v string, vs ...interface{}) {
	logger.Infof(context.TODO(), "[group:%s][%v] %s", g.Name, g.count.Load(), fmt.Sprintf(v, vs...))
}
