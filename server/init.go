package server

import (
	"context"
	"github.com/junqirao/gateway/component/registry"
)

func Init() {
	// load
	loadFromRegistry(context.Background())

	// watcher
	registry.Instance().Subscribe(context.TODO(), configRegistryKey, new(serverConfigWatcher))
	registry.Instance().Subscribe(context.TODO(), statusRegistryKey, new(serverStatusWatcher))
}

func loadFromRegistry(ctx context.Context) {
	cfgMap, err := registry.Instance().Get(ctx, configRegistryKey)
	if err != nil {
		return
	}

	for name, s := range cfgMap {
		registryConfigHandler(ctx, name, s)
	}

	statusMap, err := registry.Instance().Get(ctx, statusRegistryKey)
	if err != nil {
		return
	}
	for name, s := range statusMap {
		registryStatusHandler(ctx, name, s)
	}
}
