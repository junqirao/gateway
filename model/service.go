package model

import (
	"strings"
)

// ServiceRegisterData ...
type ServiceRegisterData struct {
	ServerName string      `json:"server_name"`  // server name
	Group      GroupInfo   `json:"group_info"`   // group info
	Service    ServiceInfo `json:"service_info"` // service info
	Callback   string      `json:"callback"`     // callback url, pushing register result: JSONResponse, success(code=0), failure(code>0)
}

// RouterPattern returns router pattern in /{group.name}/{service.name}/*
func (d *ServiceRegisterData) RouterPattern() string {
	builder := strings.Builder{}
	builder.WriteString("/")
	add := func(name string) {
		for _, s := range strings.Split(name, "/") {
			if s == "" {
				continue
			}
			builder.WriteString(s + "/")
		}
	}
	add(d.Group.Name)
	add(d.Service.Name)
	builder.WriteString("*")
	return builder.String()
}

// GroupInfo info
type GroupInfo struct {
	Name string `json:"name"`
}

// ServiceInfo ...
type ServiceInfo struct {
	Name     string                 `json:"name"`
	Protocol string                 `json:"protocol"`
	Host     string                 `json:"host"`
	Port     int                    `json:"port"`
	Meta     map[string]interface{} `json:"meta"`
}
