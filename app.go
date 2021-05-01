package lifecycle

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type App struct {
	opts 	options
	ctx 	context.Context
	cancel	func()
}

type Service interface {
	Start() error
	Stop() 	error
}

func New(opts ...Option) *App {
	options := options{
		ctx:    context.Background(),
		sigs:   []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}
	for _, o := range opts {
		o(&options)
	}
	ctx, cancel := context.WithCancel(options.ctx)
	return &App{
		opts: 	options,
		ctx: 	ctx,
		cancel: cancel,
	}
}

func (app *App) Run() error {
	g, ctx := errgroup.WithContext(app.ctx)

	for _, s := range app.opts.services {
		srv := s
		g.Go(func() error {
			<-ctx.Done()
			return srv.Stop()
		})
		g.Go(func() error {
			return srv.Start()
		})
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, app.opts.sigs...)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				_ = app.Stop()
			}
		}
	})
	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (app *App) Stop() error {
	if app.cancel != nil {
		app.cancel()
	}
	return nil
}