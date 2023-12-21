package event

// Watcher of key
type Watcher interface {
	// OnChange of key
	OnChange(key string, value []byte)
	// OnClose watcher
	OnClose(err error)
}
