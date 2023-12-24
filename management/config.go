package management

// Config of management interfaces
type Config struct {
	Enabled     bool   `json:"enabled"`
	Address     string `json:"address"`
	Port        int    `json:"port"`
	Secret      string `json:"secret"`
	IpWhitelist string `json:"ip_whitelist"`
}
