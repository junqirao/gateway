package proxy

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/junqirao/gateway/lib/response"
	"github.com/junqirao/gateway/model"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// Handler of proxy
type Handler struct {
	srvInfo model.ServiceInfo
	proxy   httputil.ReverseProxy
	prefix  string
	dialer  *net.Dialer
}

// NewHandler of proxy
func NewHandler(srv *model.ServiceRegisterData) *Handler {
	h := &Handler{
		srvInfo: srv.Service,
		proxy:   httputil.ReverseProxy{},
		prefix:  strings.TrimSuffix(srv.RouterPattern(), "/*"),
		dialer: &net.Dialer{
			Timeout:   3 * time.Second,
			KeepAlive: 30 * time.Second,
		},
	}
	h.proxy.Director = h.director
	h.proxy.Transport = &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           h.dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   3 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	h.proxy.ModifyResponse = h.modifyHandler
	h.proxy.ErrorHandler = h.errorHandler
	return h
}

// Proxy ...
func (h *Handler) Proxy(r *ghttp.Request) {
	h.proxy.ServeHTTP(r.Response.ResponseWriter, r.Request)
}

func (h *Handler) director(req *http.Request) {
	targetHost := h.srvInfo.Host
	if h.srvInfo.Port > 0 {
		targetHost += fmt.Sprintf(":%d", h.srvInfo.Port)
	}

	req.URL.Scheme = h.srvInfo.Protocol
	req.URL.Host = targetHost
	req.URL.Path = parseProxyPath(req.URL.Path, h.prefix)
	req.URL.RawPath = parseProxyPath(req.URL.RawPath, h.prefix)

	if _, ok := req.Header["User-Agent"]; !ok {
		req.Header.Set("User-Agent", "")
	}
}

func (h *Handler) errorHandler(writer http.ResponseWriter, request *http.Request, err error) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadGateway)

	if b, err := json.Marshal(response.JSON{
		Code:    http.StatusBadGateway,
		Data:    request.URL.String(),
		Message: err.Error(),
	}); err != nil {
		panic(gerror.Wrap(err, `WriteJson failed`))
	} else {
		_, _ = writer.Write(b)
	}
}

func (h *Handler) modifyHandler(response *http.Response) error {
	if response.StatusCode == http.StatusNotFound {
		return errors.New(fmt.Sprintf("not found: %v", response.Request.URL.Path))
	}
	return nil
}

func parseProxyPath(path string, prefix string) string {
	if prefix == "" || path == "" {
		return path
	}
	return strings.Replace(path, prefix, "", 1)
}
