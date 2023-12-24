package response

import "net/http"

// default
var (
	DefaultSuccessResponse = NewJSONResp(http.StatusOK, 0, "success")
	DefaultFailedResponse  = NewJSONResp(http.StatusBadRequest, 1, "unknown error")

	ErrorInvalidParams         = NewJSONResp(http.StatusBadRequest, 1000, "invalid parameters")
	ErrorResourceNotFound      = NewJSONResp(http.StatusNotFound, 1001, "requested resource not found")
	ErrorOperationNotAllowed   = NewJSONResp(http.StatusBadRequest, 1002, "operation not allowed")
	ErrorResourceAlreadyExists = NewJSONResp(http.StatusConflict, 1004, "resource already exists")
)
