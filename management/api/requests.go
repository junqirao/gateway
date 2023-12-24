package api

import (
	"github.com/junqirao/gateway/model"
)

// ServerOperationReq params
type ServerOperationReq struct {
	Name string `json:"name" dc:"server name"`
}

// ServerGetReq ...
type ServerGetReq struct {
	ServerOperationReq
}

// ServerDeleteReq ...
type ServerDeleteReq struct {
	ServerOperationReq
}

// ServerUpdateConfigReq ...
type ServerUpdateConfigReq struct {
	Name                string `json:"name" swaggerignore:"true"`
	*model.ServerConfig `json:"config"`
}
