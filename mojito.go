package mojito

import (
	"context"
	"os"
	"time"

	"github.com/rushteam/mojito/pkg/signals"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

//Service ..
type Service interface {
	Options() ServiceOptions
	Start() error
	Close() error
}

//ServiceOptions ..
type ServiceOptions interface {
	ID() string
	Name() string
}

//App ..
type App struct {
	ctx     context.Context
	stop    chan struct{}
	logger  *zap.Logger
	hooks   map[string][]func(*App)
	service []Service
	//shutdownTimeout timeout will be forces stop
	shutdownTimeout time.Duration
}

//AppOptions ..
type AppOptions func(app *App)

func (app *App) runHooks(stage string) {
	if hooks, ok := app.hooks[stage]; ok {
		for _, h := range hooks {
			h(app)
		}
	}
}

//Init ..
func Init(opts ...AppOptions) *App {
	logger, _ := zap.NewDevelopment()
	app := &App{
		stop:            make(chan struct{}),
		ctx:             context.Background(),
		logger:          logger,
		shutdownTimeout: time.Second * 2,
	}
	for _, opt := range opts {
		opt(app)
	}
	return app
}

// Run ..
func (app *App) Run(service ...Service) error {
	app.waitSignals()
	var eg errgroup.Group
	app.runHooks("before_run")
	app.service = append(app.service, service...)
	for _, srv := range app.service {
		eg.Go(func() error {
			return srv.Start()
		})
		app.logger.Debug("start", zap.String("service", srv.Options().Name()))
	}
	app.runHooks("after_run")
	graceShutdown := make(chan struct{})
	go func() {
		eg.Wait()
		close(graceShutdown)
	}()
	defer app.logger.Sync()
	for {
		select {
		// case <-time.Tick(time.Second):
		// 	app.logger.Debug(".")
		case <-graceShutdown:
			app.logger.Debug("grace shutdown")
			return nil
		case <-app.stop:
			app.logger.Debug("force shutdown")
			return nil
		}
	}
}

// graceShutdown ..
func (app *App) graceShutdown() error {
	// ctx, cancel := context.WithTimeout(app.ctx, app.shutdownTimeout)
	// defer cancel()
	pid := os.Getpid()
	app.logger.Debug("shutdown", zap.Int("pid", pid), zap.String("timeout", app.shutdownTimeout.String()))
	var eg errgroup.Group
	for _, srv := range app.service {
		eg.Go(func() error {
			app.logger.Debug("close", zap.String("service", srv.Options().Name()))
			return srv.Close()
		})
	}
	eg.Go(func() error {
		<-time.After(app.shutdownTimeout)
		close(app.stop)
		return nil
	})
	return eg.Wait()
}

func (app *App) waitSignals() {
	app.logger.Debug("init listen signal")
	signals.Shutdown(func() {
		app.graceShutdown()
	})
	// time.Sleep(time.Microsecond) //sleep 1 micro second for frist listen signal
}

/*
func waitSignals(app *App) {
	sig := make(chan os.Signal)
	signal.Notify(
		sig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
		// syscall.SIGSTOP,
		// syscall.SIGKILL,
	)
	go func() {
		app.logger.Debug("init listen signal")
		select {
		case s := <-sig:
			switch s {
			case syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM: //syscall.SIGHUP,
				app.logger.Debug("listen signal", zap.String("mod", "signal"), zap.String("signal", "SIGQUIT"))
				app.graceShutdown() // grace stop
				//syscall.SIGKILL, syscall.SIGSTOP terminate now
			}
		}
	}()
	time.Sleep(time.Microsecond) //sleep 1 micro second for frist listen signal
}
*/