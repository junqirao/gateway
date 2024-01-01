package register

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/junqirao/gateway/pkg/model"
)

// RegistryType ...
type RegistryType string

const (
	TypeEtcd RegistryType = "etcd"
)

const (
	path = "/node/"
)

// Register of gateway nodes
type Register interface {
	Register(ctx context.Context, nrd *model.NodeRegisterData) error
	Unregister(ctx context.Context, nrd *model.NodeRegisterData) error
}

// Config fetcher
type Config func() interface{}

func (c Config) get(ptr interface{}) error {
	if c == nil {
		return fmt.Errorf("config is nil")
	}
	v := c()
	if v == nil {
		return fmt.Errorf("config is nil")
	}

	marshal, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return json.Unmarshal(marshal, &ptr)
}

// New register with specific registry type and config fetcher
func New(t RegistryType, cfg Config, opts ...Option) (register Register, err error) {
	switch t {
	case TypeEtcd:
		register, err = newEtcdNodeRegister(cfg, opts...)
	default:
		return nil, fmt.Errorf("unknown registry type: %s", t)
	}
	return register, nil
}
