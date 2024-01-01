package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"strings"
)

// Router routes requests to service, by group name and service name
// protocol://host:port/{group}/{service}/{url}
type Router struct {
	serverName string
}

// NewRouter ...
func NewRouter(serverName string) *Router {
	return &Router{
		serverName: serverName,
	}
}

func (r *Router) Route(request *ghttp.Request) {
	paths := strings.Split(strings.TrimPrefix(request.URL.Path, "/"), "/")
	if len(paths) < 2 {
		r.unavailable(request)
		return
	}
	var (
		gName = paths[0]
		sName = paths[1]
	)

	group, ok := findGroup(r.serverName, gName)
	if !ok {
		r.unavailable(request)
		return
	}
	service, ok := group.Service(sName)
	if !ok {
		r.unavailable(request)
		return
	}

	service.Select().Proxy(request)
}

func (r *Router) unavailable(request *ghttp.Request, reason ...string) {
	request.Response.WriteHeader(503)
	body := "unavailable: " + request.URL.Path
	if len(reason) > 0 && reason[0] != "" {
		body = reason[0]
	}
	request.Response.Write(body)
	return
}
