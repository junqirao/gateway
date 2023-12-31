package model

import (
	"github.com/junqirao/gateway/component/registry"
	"strings"
)

// ServerGroup ...
type ServerGroup struct {
	Operation   registry.Operation `json:"operation"`
	ServerName  string             `json:"server_name"`
	ServiceName string             `json:"service_name"`
	GroupName   string             `json:"group_name"`
	LB          *LoadBalance       `json:"load_balance"` // updatable
}

// LoadBalance ...
type LoadBalance struct {
	Strategy string `json:"strategy"`
}

// NodeRegisterData ...
type NodeRegisterData struct {
	ServerGroup *ServerGroup       `json:"server_group"`
	Operation   registry.Operation `json:"operation"`           // group operation
	Node        *NodeInfo          `json:"node_info,omitempty"` // node info
}

// NodeInfo ...
type NodeInfo struct {
	Name     string                 `json:"name"`
	Protocol string                 `json:"protocol"`
	Host     string                 `json:"host"`
	Port     int                    `json:"port"`
	Meta     map[string]interface{} `json:"meta"`
}

// RouterPattern returns router pattern in /{group.name}/{service.name}/*
func (i NodeInfo) RouterPattern(groupName, serviceName string) string {
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
	add(groupName)
	add(serviceName)
	builder.WriteString("*")
	return builder.String()
}
