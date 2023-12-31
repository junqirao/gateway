package balancer

import (
	"github.com/junqirao/gateway/service/node"
	"math/rand"
	"time"
)

type random struct {
	nodes []*node.Node
	*rand.Rand
}

func newRandom() *random {
	return &random{
		nodes: nil,
		Rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (r *random) Select() *node.Node {
	switch len(r.nodes) {
	case 0:
		return nil
	case 1:
		return r.nodes[0]
	default:
		return r.nodes[r.Intn(len(r.nodes))]
	}
}

func (r *random) Update(nodes []*node.Node) {
	r.nodes = nodes
}
