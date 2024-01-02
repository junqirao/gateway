package register

import "strings"

// Option ...
type Option func(nr *etcdNodeRegister)

// WithLogger set logger
func WithLogger(logger Logger) Option {
	return func(nr *etcdNodeRegister) {
		if logger != nil {
			nr.logger = logger
		}
	}
}

// WithRegistryIdentity set registryIdentity, it will be used in registry path like "/<registryIdentity>/node/<name>.<nodeIdentity>"
func WithRegistryIdentity(identity string) Option {
	return func(nr *etcdNodeRegister) {
		nr.registryIdentity = identity
		if !strings.HasSuffix(nr.registryIdentity, "/") {
			nr.registryIdentity += "/"
		}
	}
}

// WithNodeIdentity set nodeIdentity it will be used in node path like "/<registryIdentity>/node/<name>.<nodeIdentity>"
func WithNodeIdentity(identity string) Option {
	return func(nr *etcdNodeRegister) {
		nr.nodeIdentity = identity
	}
}
