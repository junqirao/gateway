package service

import (
	"context"
	"fmt"
	"github.com/junqirao/gateway/model"
	"github.com/junqirao/gateway/proxy/balancer"
	"github.com/junqirao/gateway/proxy/node"
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
func (s *Service) UpdateOrCreateNode(ni *model.NodeInfo) {
	if s == nil || ni == nil {
		return
	}

	s.mu.Lock()
	defer func() {
		s.lb.Update(s.nodes)
		s.mu.Unlock()
	}()

	for i, n := range s.nodes {
		if n.Name == ni.Name {
			// update
			s.nodes[i] = node.New(s.group.Name, s.Name, ni)
			s.log("update node: %s", ni.Name)
			return
		}
	}

	// create
	s.nodes = append(s.nodes, node.New(s.group.Name, s.Name, ni))
	s.log("create node: %s", ni.Name)
}

// RemoveNode ...
func (s *Service) RemoveNode(name string) {
	s.mu.Lock()
	defer func() {
		s.lb.Update(s.nodes)
		s.mu.Unlock()
	}()

	for i, n := range s.nodes {
		if n.Name == name {
			s.nodes = append(s.nodes[:i], s.nodes[i+1:]...)
		}
	}
	s.log("remove node: %s", name)

	if len(s.nodes) == 0 {
		s.group.removeService(s.Name)
	}
}

// Select node
func (s *Service) Select() *node.Node {
	return s.lb.Select()
}

func (s *Service) log(v string, vs ...interface{}) {
	logger.Infof(context.TODO(), "[group:%s][service:%s][%v] %s", s.group.Name, s.Name, len(s.nodes), fmt.Sprintf(v, vs...))
}
