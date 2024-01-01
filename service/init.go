package service

import (
	"context"
	"github.com/junqirao/gateway/component/registry"
)

func Init() {
	// load
	loadFromRegistry(context.Background())

	// watcher
	registry.Instance().Subscribe(context.TODO(), registry.NodeRegPath(), new(nodeRegistryWatcher))
}

func loadFromRegistry(ctx context.Context) {
	cfgMap, err := registry.Instance().Get(ctx, registry.NodeRegPath())
	if err != nil {
		return
	}

	for name, s := range cfgMap {
		nodeRegistryHandler(ctx, name, s)
	}
}
