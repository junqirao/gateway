package authorization

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/junqirao/gateway/lib/response"
	"strings"
	"sync"
	"time"
)

var (
	ipWhitelist = sync.Map{}
	debug       bool
)

const (
	defaultTimestampGap = 1000 * 60
)

func init() {
	debug = false
	v, err := g.Cfg().Get(context.Background(), "debug", false)
	if err == nil {
		debug = v.Bool()
	}
}

// VerifySignature
// response.DefaultRequestTimeoutResponse when
// X-Timestamp - time.Now > defaultTimestampGap, disabled when debug = true
// response.DefaultUnauthorizedResponse when signature mismatch
func VerifySignature(secret string) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		timestamp := gconv.Int64(r.GetHeader("X-Timestamp"))
		if !debug {
			if v := time.Now().Unix() - timestamp; v <= 0 || v > defaultTimestampGap {
				response.Error(r, response.DefaultRequestTimeoutResponse)
				return
			}
		}
		if sign(r.GetHeader("X-Nonce"), r.GetHeader("X-Timestamp"), secret) != r.GetHeader("X-Signature") {
			response.Error(r, response.DefaultUnauthorizedResponse.WithDetail("signature mismatch"))
			return
		}
		r.Middleware.Next()
	}
}

// CheckIpWhitelist ...
func CheckIpWhitelist(whitelist string) func(r *ghttp.Request) {
	for _, ip := range strings.Split(whitelist, ",") {
		ipWhitelist.Store(ip, struct{}{})
	}

	return func(r *ghttp.Request) {
		_, ok := ipWhitelist.Load(r.GetClientIp())
		if !ok {
			response.Error(r, response.DefaultBlockedResponse.WithDetail("ip address not allowed"))
			return
		}
		r.Middleware.Next()
	}
}
