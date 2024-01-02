package register

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/junqirao/gateway/pkg/model"
	"os"
	"os/signal"
	"syscall"
)

// RegistryType ...
type RegistryType string

const (
	TypeEtcd RegistryType = "etcd"
)

const (
	path = "node/"
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

// Automatic register and unregister, by listening os exist signal, it will block
func Automatic(ctx context.Context, t RegistryType, cfg Config, nrd *model.NodeRegisterData, opts ...Option) (err error) {
	register, err := New(t, cfg, opts...)
	if err != nil {
		return
	}

	if err = register.Register(ctx, nrd); err != nil {
		return
	}

	osc := make(chan os.Signal, 1)
	signal.Notify(osc, syscall.SIGTERM, syscall.SIGINT)
	<-osc
	return register.Unregister(ctx, nrd)
}
