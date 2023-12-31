package proxy

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// Router routes requests to service, by group name and service name
// protocol://host:port/{group}/{service}/{url}
type Router struct {
}

// NewRouter ...
func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Route(request *ghttp.Request) {
	// paths := strings.Split(request.URL.Path, "/")
	// if len(paths) < 2 {
	// 	r.unavailable(request)
	// 	return
	// }
	// var (
	// 	gName = paths[0]
	// 	sName = paths[1]
	// )
	//
	// gr, ok := group.Find(gName)
	// if !ok {
	// 	r.unavailable(request)
	// 	return
	// }
	//
	// h, ok := gr.Select(sName)
	// if !ok {
	// 	r.unavailable(request)
	// 	return
	// }
	//
	// h.Proxy(request)
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
