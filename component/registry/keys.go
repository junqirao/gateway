package registry

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	identity = g.Cfg().MustGet(context.Background(), "registry.identity", "local.undefined").String()

	// path define, end with "/"
	pathNode         = identity + "/node/"
	pathServerConfig = identity + "/server/config/"
)

// NodeRegKey returns node registry key with specific name
func NodeRegKey(name string) string {
	return pathNode + name
}

// NodeRegPath returns node registry path
func NodeRegPath() string {
	return pathNode
}

// ServerConfigRegKey returns server config registry key with specific name
func ServerConfigRegKey(name string) string {
	return pathNode + name
}

// ServerConfigRegPath returns server config registry path
func ServerConfigRegPath() string {
	return pathServerConfig
}
