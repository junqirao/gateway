package node

import (
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
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

func (n *Node) Proxy(r *ghttp.Request) {
	if n == nil {
		r.Response.WriteHeader(502)
		r.Response.Write(fmt.Sprintf("unavailable: %s", r.URL.Path))
		return
	}
	n.Handler.Proxy(r)
}
