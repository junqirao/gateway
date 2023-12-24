package api

import (
	"github.com/junqirao/gateway/model"
)

// ServerOperationReq params
type ServerOperationReq struct {
	Name string `json:"name" dc:"server name"`
}

// ServerCreateReq ...
type ServerCreateReq struct {
	Config *model.ServerConfig `json:"config"`
	Status *model.ServerStatus `json:"status"`
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
	Name   string              `json:"name" swaggerignore:"true"`
	Config *model.ServerConfig `json:"config"`
}

// ServerUpdateStatusReq ...
type ServerUpdateStatusReq struct {
	Name   string              `json:"name" swaggerignore:"true"`
	Status *model.ServerStatus `json:"status"`
}
