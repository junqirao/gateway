package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"time"
)

// ServerInfo ...
type ServerInfo struct {
	ServerConfig

	Name string `json:"name"`
}

// ServerConfig ...
type ServerConfig struct {
	Enabled    bool              `json:"enabled"`
	Properties *ServerProperties `json:"properties"`
}

type ServerProperties struct {
	// Address specifies the server listening address like "port" or ":port",
	// multiple addresses joined using ','.
	Address string `json:"address"`
	// HTTPSAddr specifies the HTTPS addresses, multiple addresses joined using char ','.
	HTTPSAddr string `json:"https_addr"`
	// Endpoints are custom endpoints for service register, it uses Address if empty.
	Endpoints []string `json:"endpoints"`
	// HTTPSCertPath specifies certification file path for HTTPS service.
	HTTPSCertPath string `json:"https_cert_path"`
	// HTTPSKeyPath specifies the key file path for HTTPS service.
	HTTPSKeyPath string `json:"https_key_path"`
	// ReadTimeout is the maximum duration for reading the entire
	// request, including the body.
	//
	// Because ReadTimeout does not let Handlers make per-request
	// decisions on each request body's acceptable deadline or
	// upload rate, most users will prefer to use
	// ReadHeaderTimeout. It is valid to use them both.
	ReadTimeout int64 `json:"read_timeout"`
	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
	WriteTimeout int64 `json:"write_timeout"`
	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alive are enabled. If IdleTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	IdleTimeout int64 `json:"idle_timeout"`
	// MaxHeaderBytes controls the maximum number of bytes the
	// server will read parsing the request header's keys and
	// values, including the request line. It does not limit the
	// size of the request body.
	//
	// It can be configured in configuration file using string like: 1m, 10m, 500kb etc.
	// It's 10240 bytes in default.
	MaxHeaderBytes int `json:"max_header_bytes"`
	// KeepAlive enables HTTP keep-alive.
	KeepAlive bool `json:"keep_alive"`
	// ServerAgent specifies the server agent information, which is wrote to
	// HTTP response header as "Server".
	ServerAgent string `json:"server_agent"`
	// LogPath specifies the directory for storing logging files.
	LogPath string `json:"log_path"`
	// LogLevel specifies the logging level for logger.
	LogLevel string `json:"log_level"`
	// LogStdout specifies whether printing logging content to stdout.
	LogStdout bool `json:"log_stdout"`
	// ErrorStack specifies whether logging stack information when error.
	ErrorStack bool `json:"error_stack"`
	// ErrorLogEnabled enables error logging content to files.
	ErrorLogEnabled bool `json:"error_log_enabled"`
	// ErrorLogPattern specifies the error log file pattern like: error-{Ymd}.log
	ErrorLogPattern string `json:"error_log_pattern"`
	// AccessLogEnabled enables access logging content to files.
	AccessLogEnabled bool `json:"access_log_enabled"`
	// AccessLogPattern specifies the error log file pattern like: access-{Ymd}.log
	AccessLogPattern string `json:"access_log_pattern"`
	// ClientMaxBodySize specifies the max body size limit in bytes for client request.
	// It can be configured in configuration file using string like: 1m, 10m, 500kb etc.
	// It's `8MB` in default.
	ClientMaxBodySize string `json:"client_max_body_size"`
}

func (c *ServerConfig) FillDefault() {
	if c.Properties == nil {
		return
	}
	if c.Properties.ClientMaxBodySize == "" {
		c.Properties.ClientMaxBodySize = "8m"
	}
}

func (c *ServerConfig) C(name string) ghttp.ServerConfig {
	config := ghttp.NewConfig()
	config.Name = name
	if c.Properties.ReadTimeout > 0 {
		config.ReadTimeout = time.Duration(c.Properties.ReadTimeout)
	}
	if c.Properties.IdleTimeout > 0 {
		config.IdleTimeout = time.Duration(c.Properties.IdleTimeout)
	}
	if c.Properties.WriteTimeout > 0 {
		config.WriteTimeout = time.Duration(c.Properties.WriteTimeout)
	}

	if config.KeepAlive != c.Properties.KeepAlive {
		config.KeepAlive = c.Properties.KeepAlive
	}
	if c.Properties.MaxHeaderBytes > 0 {
		config.MaxHeaderBytes = c.Properties.MaxHeaderBytes
	}

	config.Address = c.Properties.Address
	config.HTTPSAddr = c.Properties.HTTPSAddr
	config.Endpoints = c.Properties.Endpoints
	config.HTTPSCertPath = c.Properties.HTTPSCertPath
	config.HTTPSKeyPath = c.Properties.HTTPSKeyPath

	config.ServerAgent = c.Properties.ServerAgent
	config.LogPath = c.Properties.LogPath
	config.LogLevel = c.Properties.LogLevel
	config.ErrorStack = c.Properties.ErrorStack
	config.ErrorLogEnabled = c.Properties.ErrorLogEnabled
	config.ErrorLogPattern = c.Properties.ErrorLogPattern
	config.AccessLogEnabled = c.Properties.AccessLogEnabled
	config.AccessLogPattern = c.Properties.AccessLogPattern
	config.ClientMaxBodySize = gfile.StrToSize(c.Properties.ClientMaxBodySize)

	// default config
	config.DumpRouterMap = false
	return config
}
