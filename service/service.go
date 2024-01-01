package service

import (
	"github.com/junqirao/gateway/component/registry"
	"github.com/junqirao/gateway/model"
	"github.com/junqirao/gateway/service/balancer"
	"github.com/junqirao/gateway/service/node"
	"sync"
)

// Service ...
type Service struct {
	mu    sync.RWMutex
	group *Group
	Name  string
	nodes []*node.Node
	lb    balancer.Balancer
}

// UpdateOrCreateNode ...
func (s *Service) UpdateOrCreateNode(ni *model.NodeInfo, op registry.Operation) {
	if s == nil || ni == nil || op.IsEmpty() {
		return
	}

	defer func() {
		s.lb.Update(s.nodes)
	}()

	for i, n := range s.nodes {
		if n.Name == ni.Name {
			if op.IsDelete() {
				s.nodes = append(s.nodes[:i], s.nodes[i+1:]...)
				return
			} else if op.IsUpdate() {
				s.nodes[i] = node.New(s.group.Name, s.Name, ni)
				return
			} else {
				return
			}
		}
	}

	s.nodes = append(s.nodes, node.New(s.group.Name, s.Name, ni))
}

// Select node
func (s *Service) Select() *node.Node {
	return s.lb.Select()
}
