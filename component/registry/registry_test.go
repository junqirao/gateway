package registry

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

type watcher struct {
	wg  sync.WaitGroup
	sig chan struct{}
}

func (n *watcher) Add() {
	n.wg.Add(1)
}

func (n *watcher) OnChange(key string, value []byte) {
	fmt.Println("key   --> " + key)
	fmt.Println("value --> " + string(value))
}

func (n *watcher) OnClose(err error) {
	n.wg.Done()
	fmt.Println(fmt.Sprintf("closed with error: %v", err))
}

func TestSubscribe(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	w := new(watcher)
	w.Add()
	Instance().Subscribe(ctx, "test", w)

	err := Instance().Set(context.TODO(), "test", "123")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		t.Log("cancel after 2 seconds")
		time.Sleep(time.Second * 2)
		cancelFunc()
	}()
	w.wg.Wait()
}

func TestSubscribePrefix(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	w := new(watcher)
	w.Add()
	Instance().Subscribe(ctx, "test/", w)

	err := Instance().Set(context.TODO(), "test/aaa", "123")
	if err != nil {
		t.Fatal(err)
	}
	err = Instance().Set(context.TODO(), "test/bbb", "123")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		t.Log("cancel after 2 seconds")
		time.Sleep(time.Second * 2)
		cancelFunc()
	}()
	w.wg.Wait()
}
