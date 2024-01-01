package balancer

import (
	"github.com/junqirao/gateway/model"
	"github.com/junqirao/gateway/proxy/node"
)

const (
	StrategyRandom         = "random"
	StrategyRandomWeighted = "random-weighted"
	StrategySequential     = "sequential"
	StrategyRoundRobin     = "round-robin"
)

var (
	defaultConfig = &model.LoadBalance{
		Strategy: StrategyRandom,
	}
)

type Balancer interface {
	Select() *node.Node
	Update(nodes []*node.Node)
}

func New(cfg *model.LoadBalance) Balancer {
	if cfg == nil {
		cfg = defaultConfig
	}
	if cfg.Strategy == "" {
		cfg.Strategy = StrategyRandom
	}

	var balancer Balancer
	switch cfg.Strategy {
	case StrategyRandom:
		balancer = newRandom()
	case StrategyRandomWeighted:
	case StrategySequential:
	case StrategyRoundRobin:

	}
	return balancer
}
