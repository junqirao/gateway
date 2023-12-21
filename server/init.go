package server

import (
	"context"
	"github.com/junqirao/gateway/component/registry"
)

func Init() {
	loadAllFromRegistry(context.Background())
	registry.Instance().Subscribe(context.TODO(), registryKey, new(serverConfigWatcher))
}
