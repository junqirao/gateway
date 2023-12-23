package response

import "net/http"

// default
var (
	DefaultSuccessResponse = NewJSONResp(http.StatusOK, 0, "success")
	DefaultFailedResponse  = NewJSONResp(http.StatusBadRequest, 1, "unknown error")

	ErrorInvalidParams    = NewJSONResp(http.StatusBadRequest, 1000, "invalid parameters")
	ErrorResourceNotFound = NewJSONResp(http.StatusNotFound, 1001, "requested resource not found")
)
