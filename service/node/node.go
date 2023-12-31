package node

import (
	"github.com/junqirao/gateway/model"
	"github.com/junqirao/gateway/proxy"
	"strings"
)

// Node ...
type Node struct {
	// proxy
	*proxy.Handler

	// healthy check

	// runtime
	model.NodeInfo
}

// New ...
func New(groupName, serviceName string, info *model.NodeInfo) *Node {
	return &Node{
		Handler:  proxy.NewHandler(info, strings.TrimSuffix(info.RouterPattern(groupName, serviceName), "/*")),
		NodeInfo: *info,
	}
}
