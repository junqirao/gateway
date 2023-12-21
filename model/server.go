package model

import "github.com/gogf/gf/v2/net/ghttp"

type ServerConfig struct {
	Up bool `json:"up"`
	*ghttp.ServerConfig
}
