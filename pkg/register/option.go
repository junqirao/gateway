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

// WithIdentity set identity, it will be used in registry path like "/<identity>/node/<name>"
func WithIdentity(identity string) Option {
	return func(nr *etcdNodeRegister) {
		nr.identity = identity
		if !strings.HasSuffix(nr.identity, "/") {
			nr.identity += "/"
		}
	}
}
