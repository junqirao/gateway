package model

import (
	"fmt"
	"strings"
)

// ServerGroup ...
type ServerGroup struct {
	// position
	ServerName  string `json:"server_name"`
	GroupName   string `json:"group_name"`
	ServiceName string `json:"service_name"`

	// config
	LB *LoadBalance `json:"load_balance"` // updatable
}

// LoadBalance ...
type LoadBalance struct {
	Strategy string `json:"strategy"`
}

// NodeRegisterData ...
type NodeRegisterData struct {
	ServerGroup *ServerGroup `json:"server_group"`
	Node        *NodeInfo    `json:"node_info,omitempty"` // node info
}

// RegistryKey ...
func (n *NodeRegisterData) RegistryKey() string {
	return fmt.Sprintf("%s.%s.%s.%s", n.ServerGroup.ServerName, n.ServerGroup.GroupName, n.ServerGroup.ServiceName, n.Node.Name)
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
