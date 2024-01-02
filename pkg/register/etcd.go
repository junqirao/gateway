package register

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/junqirao/gateway/pkg/model"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"
)

// etcdNodeRegister ...
type etcdNodeRegister struct {
	cli              *clientv3.Client
	logger           Logger
	registryIdentity string
	nodeIdentity     string
}

func newEtcdNodeRegister(cfg Config, opts ...Option) (*etcdNodeRegister, error) {
	config := clientv3.Config{}
	if err := cfg.get(&config); err != nil {
		return nil, err
	}

	cli, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}

	nr := &etcdNodeRegister{
		cli:    cli,
		logger: defaultLogger,
	}
	for i := range opts {
		opts[i](nr)
	}
	if nr.nodeIdentity == "" {
		hostname, _ := os.Hostname()
		os.Getpid()
		nr.nodeIdentity = fmt.Sprintf("%s_%d", hostname, os.Getpid())
	}
	return nr, nil
}

// Register implements Register interface
func (r *etcdNodeRegister) Register(ctx context.Context, nrd *model.NodeRegisterData) error {
	name := nrd.RegistryKey()
	key := r.key(name)

	marshal, err := json.Marshal(nrd)
	if err != nil {
		return err
	}
	if _, err = r.cli.Put(ctx, key, string(marshal)); err != nil {
		return err
	}

	r.logger.Printf("node registered: %s", name)
	return nil
}

// Unregister implements Register interface
func (r *etcdNodeRegister) Unregister(ctx context.Context, nrd *model.NodeRegisterData) error {
	name := nrd.RegistryKey()

	if _, err := r.cli.Delete(ctx, r.key(name)); err != nil {
		return err
	}
	r.logger.Printf("node unregistered: %s", name)
	return nil
}

func (r *etcdNodeRegister) key(name string) string {
	return fmt.Sprintf("%s%s%s.%s", r.registryIdentity, path, name, r.nodeIdentity)
}
