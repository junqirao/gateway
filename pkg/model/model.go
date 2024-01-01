package model

import "fmt"

// ServerGroup ...
type ServerGroup struct {
	ServerName  string `json:"server_name"`
	GroupName   string `json:"group_name"`
	ServiceName string `json:"service_name"`
}

type NodeRegisterData struct {
	ServerGroup ServerGroup `json:"server_group"`
	Node        NodeInfo    `json:"node_info,omitempty"` // node info
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
