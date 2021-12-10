package server

import (
	"log"
	"time"

	// "git.tvblack.com/video/frame/core"
	"github.com/goggle1/go-jizhi/video/frame/core"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	web "github.com/micro/go-web"
)

type microwebComp struct {
	Service web.Service
}

func NewMicroWeb(name, version, ip string, etcd []string) *microwebComp {

	m := &microwebComp{}

	registry := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = etcd
	})

	// New Service
	m.Service = web.NewService(
		web.Name(name),
		web.Registry(registry),
		web.Version(version),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Address(ip),
	)
	return m
}

func (m *microwebComp) Start(app core.Appcation) {
	// Initialise service
	m.Service.Init(
		web.AfterStart(func() error {
			app.Run()
			return nil
		}),
		web.AfterStop(func() error {
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
