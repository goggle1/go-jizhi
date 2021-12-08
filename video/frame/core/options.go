package core

type Componenter interface {
	Init()
	Run()
	Stop()
}

func NewCompoent(c Componenter) Option {
	return func(o *Options) {
		o.component = append(o.component, c)
	}
}
