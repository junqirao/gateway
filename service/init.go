package service

import (
	"context"
	"github.com/junqirao/gateway/component/registry"
)

func Init() {
	// load
	loadFromRegistry(context.Background())

	// watcher
	registry.Instance().Subscribe(context.TODO(), registryKey, new(serviceRegistryWatcher))
}

func loadFromRegistry(ctx context.Context) {
	cfgMap, err := registry.Instance().Get(ctx, registryKey)
	if err != nil {
		return
	}

	for name, s := range cfgMap {
		registryHandler(ctx, name, s, true)
	}
}
