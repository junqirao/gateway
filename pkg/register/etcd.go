package register

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/junqirao/gateway/pkg/model"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
)

// etcdNodeRegister ...
type etcdNodeRegister struct {
	cli      *clientv3.Client
	logger   Logger
	identity string
	mu       sync.RWMutex

	rs map[string]*model.NodeRegisterData
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
		mu:     sync.RWMutex{},
		rs:     make(map[string]*model.NodeRegisterData),
	}
	for i := range opts {
		opts[i](nr)
	}
	return nr, nil
}

// Register implements Register interface
func (r *etcdNodeRegister) Register(ctx context.Context, nrd *model.NodeRegisterData) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	name := nrd.RegistryKey()
	if v := r.rs[name]; v != nil {
		return fmt.Errorf("node already registered: %s", name)
	}

	marshal, err := json.Marshal(nrd)
	if err != nil {
		return err
	}
	if _, err = r.cli.Put(ctx, r.key(name), string(marshal)); err != nil {
		return err
	}

	r.rs[name] = nrd
	r.logger.Printf("node registered: %s", name)
	return nil
}

// Unregister implements Register interface
func (r *etcdNodeRegister) Unregister(ctx context.Context, nrd *model.NodeRegisterData) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	name := nrd.RegistryKey()

	if _, err := r.cli.Delete(ctx, r.key(name)); err != nil {
		return err
	}
	delete(r.rs, name)
	r.logger.Printf("node unregistered: %s", name)
	return nil
}

func (r *etcdNodeRegister) key(name string) string {
	return fmt.Sprintf("%s%s%s", r.identity, path, name)
}
