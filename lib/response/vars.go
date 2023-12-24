package response

import "net/http"

// default code define
var (
	CodeRequestTimeoutResponse = NewJSONResp(http.StatusRequestTimeout, -4, "request timeout")
	CodeBlockedResponse        = NewJSONResp(http.StatusForbidden, -3, "blocked")
	CodeUnauthorizedResponse   = NewJSONResp(http.StatusForbidden, -2, "unauthorized")
	CodeDefaultFailedResponse  = NewJSONResp(http.StatusBadRequest, -1, "unknown error")
	CodeDefaultSuccessResponse = NewJSONResp(http.StatusOK, 0, "success")
)

// business code define
var (
	ErrorInvalidParams         = NewJSONResp(http.StatusBadRequest, 1000, "invalid parameters")
	ErrorResourceNotFound      = NewJSONResp(http.StatusNotFound, 1001, "requested resource not found")
	ErrorOperationNotAllowed   = NewJSONResp(http.StatusBadRequest, 1002, "operation not allowed")
	ErrorResourceAlreadyExists = NewJSONResp(http.StatusConflict, 1004, "resource already exists")
)
