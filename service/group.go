package service

import (
	"github.com/junqirao/gateway/model"
	"github.com/junqirao/gateway/service/balancer"
	"github.com/junqirao/gateway/service/node"
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
func CreateOrGetGroup(groupName string) *Group {
	if g, ok := findGroup(groupName); ok {
		return g
	}
	g := new(Group)
	g.Name = groupName
	groups.Store(groupName, g)
	return g
}

func findGroup(groupName string) (*Group, bool) {
	if v, ok := groups.Load(groupName); ok {
		if g, ok := v.(*Group); ok {
			return g, ok
		}
	}
	return nil, false
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
