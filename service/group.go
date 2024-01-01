package service

import (
	"github.com/junqirao/gateway/model"
	"github.com/junqirao/gateway/service/balancer"
	"github.com/junqirao/gateway/service/node"
	"strings"
	"sync"
)

var (
	groups = sync.Map{} // key: group_name, value: *Group
)

// Group ...
type Group struct {
	Name     string
	services sync.Map // service_name : Service
}

// CreateOrGetGroup ...
func CreateOrGetGroup(serverName, groupName string) *Group {
	if g, ok := findGroup(serverName, groupName); ok {
		return g
	}
	g := new(Group)
	g.Name = groupName
	groups.Store(serverGroupName(serverName, groupName), g)
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

// SubmitChanges ...
func (g *Group) SubmitChanges(sg *model.ServerGroup) *Service {
	if sg.Operation.IsDelete() {
		g.services.Delete(sg.ServiceName)
		return nil
	}

	var service *Service
	v, ok := g.services.Load(sg.ServiceName)
	if !ok {
		service = g.CreateAndSaveService(sg)
	} else {
		service = v.(*Service)
		if sg.Operation.IsUpdate() {
			service.lb = balancer.New(sg.LB)
		}
	}

	return service
}

// CreateAndSaveService ...
func (g *Group) CreateAndSaveService(sg *model.ServerGroup) *Service {
	service := &Service{
		group: g,
		nodes: make([]*node.Node, 0),
		lb:    balancer.New(sg.LB),
		Name:  sg.ServiceName,
	}
	g.services.Store(sg.ServiceName, service)
	return service
}

// Service ...
func (g *Group) Service(serviceName string) (*Service, bool) {
	if v, ok := g.services.Load(serviceName); ok {
		if service, ok := v.(*Service); ok {
			return service, ok
		}
	}
	return nil, false
}
