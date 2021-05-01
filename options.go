package lifecycle

import (
	"context"
	"os"
)

type options struct {
	name 		string

	ctx 		context.Context
	sigs 		[]os.Signal
	services 	[]Service
}

type Option func(o *options)

func Name(name string) Option {
	return func(o *options) { o.name = name }
}

func Context(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

func Signal(sigs []os.Signal) Option {
	return func(o *options) { o.sigs = sigs }
}

func Services(services ...Service) Option {
	return func(o *options) { o.services = services }
}