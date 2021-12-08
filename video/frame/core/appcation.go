package core

import (
	"git.tvblack.com/video/frame/command"
	"git.tvblack.com/video/frame/librarys"
	"git.tvblack.com/video/frame/service"
)

type Appcation interface {
	Init()
	Run()
	Stop()
	NewComponent(c Componenter)
	GetService() service.SerManager
	GetCommand() command.CmdManager
}

type Options struct {
	component []Componenter
	service   service.SerManager
	command   command.CmdManager
}

type Option func(*Options)

type appcation struct {
	opts Options
}

var App = NewApp()
var TimeTask = librarys.NewTimingTask()

func NewApp(opts ...Option) Appcation {

	opt := newApp(opts...)

	app := &appcation{
		opts: opt,
	}
	return app
}

func newApp(opts ...Option) Options {
	opt := Options{
		service: service.DefaultService,
		command: command.DefaultCommand,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func (a *appcation) NewComponent(c Componenter) {
	a.opts.component = append(a.opts.component, c)
}

func (a *appcation) Init() {
	for _, c := range a.opts.component {
		c.Init()
	}
}

func (a *appcation) Run() {
	for _, c := range a.opts.component {
		c.Run()
	}

	services := a.opts.service.GetAll()
	for _, s := range services {
		s.OnCreate()
	}
}

func (a *appcation) Stop() {
	for _, c := range a.opts.component {
		c.Stop()
	}

	services := a.opts.service.GetAll()
	for _, s := range services {
		s.OnDestroy()
	}

}

func (a *appcation) GetService() service.SerManager {
	return a.opts.service
}

func (a *appcation) GetCommand() command.CmdManager {
	return a.opts.command
}
