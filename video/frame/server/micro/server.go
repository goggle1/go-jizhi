package server

import (
	"context"
	"log"
	"time"

	"git.tvblack.com/video/frame/proto/p_common"

	"git.tvblack.com/video/frame/core"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/registry/etcdv3"
)

type microComp struct {
	Service micro.Service
}

func NewMicro(name, version string, etcd []string, addr string) *microComp {

	m := &microComp{}

	registry := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = etcd
	})

	// New Service
	m.Service = micro.NewService(
		micro.Name(name),
		micro.Version(version),
		micro.Registry(registry),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Address(addr),
	)
	// graceful
	m.Service.Server().Init(
		server.Wait(true),
	)

	return m
}

func (m *microComp) Start(app core.Appcation, w ...server.HandlerWrapper) {
	// Initialise service
	m.Service.Init(
		// handler wrap
		micro.WrapHandler(w...),
		micro.AfterStart(func() error {
			app.Run()
			return nil
		}),
		micro.AfterStop(func() error {
			app.Stop()
			return nil
		}),
	)

	app.Init()

	// Run service
	if err := m.Service.Run(); err != nil {
		log.Fatal(err)
	}
}

func Process(ctx context.Context, event *p_common.StringMsg) error {
	return nil
}
