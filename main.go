package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gateway/management"
	"github.com/junqirao/gateway/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server.Init()
	management.Init()

	graceExit()
}

func graceExit() {
	osc := make(chan os.Signal, 1)
	signal.Notify(osc, syscall.SIGTERM, syscall.SIGINT)
	s := <-osc
	g.Log().Infof(context.TODO(), "recv stop sig: %s", s)
	os.Exit(0)
}
