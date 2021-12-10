package service

import (
	"fmt"

	// "git.tvblack.com/video/frame/proto/p_common"
	"github.com/goggle1/go-jizhi/video/frame/proto/p_common"

	"github.com/sirupsen/logrus"
)

type ServiceType int

var (
	DefaultService = newService()
)

type Servicer interface {
	OnCreate()
	OnReload(msg *p_common.NoticeMsg)
	OnDestroy()
	GetType() ServiceType
}

type SerManager interface {
	Options() Options
	Register(ser Servicer)
	Get(sType ServiceType) Servicer
	GetAll() map[ServiceType]Servicer
	OnReload(msg *p_common.NoticeMsg)
	OnDestroy()
}

func newService(opts ...Option) *service {
	opt := newOptions(opts...)
	return &service{
		opts: opt,
	}
}

type Options struct {
	ctls map[ServiceType]Servicer
}

type Option func(*Options)

type service struct {
	opts Options
}

func newOptions(opts ...Option) Options {
	opt := Options{
		ctls: make(map[ServiceType]Servicer),
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func (s *service) Register(ser Servicer) {
	if s.opts.ctls[ser.GetType()] != nil {
		panic(fmt.Errorf("[service] error:%s", ser.GetType()))
	}
	s.opts.ctls[ser.GetType()] = ser
}

func (s *service) OnReload(msg *p_common.NoticeMsg) {
	for _, v := range s.opts.ctls {
		v.OnReload(msg)
	}
}

func (s *service) OnDestroy() {
	for _, v := range s.opts.ctls {
		v.OnDestroy()
	}
}

func (s *service) Get(sType ServiceType) Servicer {
	ier, ok := s.opts.ctls[sType]
	if !ok {
		logrus.Errorln("[service] get error:", sType)
		return nil
	}
	return ier
}

func (s *service) Options() Options {
	return s.opts
}

func (s *service) GetAll() map[ServiceType]Servicer {
	return s.opts.ctls
}
