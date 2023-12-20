package registry

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestSubscribe(t *testing.T) {
	Instance().Subscribe(context.TODO(), "/test", func(key string, value []byte) {
		fmt.Println("key   --> " + key)
		fmt.Println("value --> " + string(value))
	})
	time.Sleep(time.Minute * 10)
}
