package command

import (
	"fmt"

	// "git.tvblack.com/video/frame/message"
	"github.com/goggle1/go-jizhi/video/frame/message"

	"github.com/sirupsen/logrus"
)

var (
	DefaultCommand = newCommand()
)

type Commander interface {
	DoRequest(Commander, *message.ReqContext) *message.Result
	OnExecute(*message.ReqContext, message.IContext) *message.Result
	IsLogin() bool
}

type CmdManager interface {
	Options() Options
	Register(cmd string, ser Commander)
	Get(sType string) Commander
}

func newCommand(opts ...Option) *command {
	opt := newOptions(opts...)
	return &command{
		opts: opt,
	}
}

type Options struct {
	ctls map[string]Commander
}

type Option func(*Options)

type command struct {
	opts Options
}

func newOptions(opts ...Option) Options {
	opt := Options{
		ctls: make(map[string]Commander),
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func (s *command) Register(cmd string, ser Commander) {
	if s.opts.ctls[cmd] != nil {
		panic(fmt.Errorf("[service] error:%s", cmd))
	}
	s.opts.ctls[cmd] = ser
}

func (s *command) Get(sType string) Commander {
	ier, ok := s.opts.ctls[sType]
	if !ok {
		logrus.Errorln("[service] get error:", sType)
		return nil
	}
	return ier
}
func (s *command) Options() Options {
	return s.opts
}
