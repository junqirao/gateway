package model

import (
	"strings"
)

// ServerGroup ...
type ServerGroup struct {
	ServerName string     `json:"server_name"`
	Group      *GroupInfo `json:"group"`
}

// ServiceRegisterData ...
type ServiceRegisterData struct {
	ServerGroup
	Service ServiceInfo `json:"service_info"` // service info
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
