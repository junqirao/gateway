package etcd

import (
	"context"
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	testKey := "test/abc"
	_, err := Ins.cli.Put(context.TODO(), testKey, "abc")
	if err != nil {
		t.Fatal(err)
		return
	}
	get, err := Ins.cli.Get(context.TODO(), testKey)
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, kv := range get.Kvs {
		fmt.Println("kv str: ", kv.String())
	}
	_, err = Ins.cli.Delete(context.TODO(), testKey)
	if err != nil {
		t.Fatal(err)
		return
	}
}
