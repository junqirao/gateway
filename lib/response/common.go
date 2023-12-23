package response

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// Option of JSON
type Option func(*JSON)

// WithStatusCode set http status code
func WithStatusCode(code int) Option {
	return func(o *JSON) {
		o.Status = code
	}
}

// WithData set response data
func WithData(data interface{}) Option {
	if resp, ok := data.(JSON); ok {
		return func(o *JSON) {
			o.Code = resp.Code
			o.Message = resp.Message
			o.Status = resp.Status
			o.Data = data
		}
	} else {
		return func(o *JSON) {
			o.Data = data
		}
	}
}

// WithMessage set response message
func WithMessage(msg string) Option {
	return func(o *JSON) {
		o.Message = msg
	}
}

// WithCode set business code
func WithCode(code int) Option {
	return func(o *JSON) {
		o.Code = code
	}
}

// Success response with default status code http.StatusOK
func Success(r *ghttp.Request, opts ...Option) {
	resp := DefaultSuccessResponse
	for i := range opts {
		opts[i](&resp)
	}

	r.Response.WriteHeader(resp.Status)
	r.Response.WriteJsonExit(resp)
}

// Data response with default status code http.StatusOK and data
func Data(r *ghttp.Request, data interface{}, opts ...Option) {
	resp := DefaultSuccessResponse
	for i := range opts {
		opts[i](&resp)
	}
	resp.Data = data
	r.Response.WriteHeader(resp.Status)
	r.Response.WriteJsonExit(resp)
}

// Error response with default status code http.StatusBadRequest
func Error(r *ghttp.Request, err error, opts ...Option) {
	resp := DefaultFailedResponse
	if r, ok := err.(JSON); ok {
		resp = r
	}
	resp.Message = err.Error()

	for i := range opts {
		opts[i](&resp)
	}

	r.Response.WriteHeader(resp.Status)
	r.Response.WriteJsonExit(resp)
}

// Result response both data and error.
// if error not nil it response default status http.StatusOK,else http.StatusBadRequest
func Result(r *ghttp.Request, err error, data interface{}, opts ...Option) {
	resp := DefaultSuccessResponse
	if err != nil {
		resp = DefaultFailedResponse
		resp.Message = err.Error()
	}
	resp.Data = data

	for i := range opts {
		opts[i](&resp)
	}

	r.Response.WriteHeader(resp.Status)
	r.Response.WriteJsonExit(resp)
}
