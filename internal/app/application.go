package app

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wajox/krakend-proxy/internal/app/dependencies"
	"github.com/wajox/krakend-proxy/internal/app/initializers"

	//nolint
	_ "github.com/gobuffalo/envy"
)

type Application struct {
	ReloadCh  chan struct{}
	Container *dependencies.Container
}

func InitializeApplication() (*Application, error) {
	if err := initializers.InitializeLogs(); err != nil {
		return nil, err
	}

	application, err := BuildApplication()
	if err != nil {
		return nil, err
	}

	application.ReloadCh = make(chan struct{})

	return application, nil
}

func (a *Application) Start(ctx context.Context, cli bool) {
	if cli {
		return
	}

	a.StartKrakenD(ctx)
}

func (a *Application) StartKrakenD(ctx context.Context) {
	r := initializers.InitializeKrakenD(ctx, a.Container.Cfg, a.Container.Logger)

	go func() {
		r.Run(*a.Container.Cfg)
	}()
}

func (a *Application) Stop(reloaded bool) (err error) {
	return nil
}

// Reload - reload app and return pointer to new app instance.
func (a *Application) Reload(ctx context.Context, cli bool) (*Application, error) {
	if err := a.Stop(true); err != nil {
		return nil, err
	}

	a, err := InitializeApplication()
	if err != nil {
		log.Fatal().Err(err).Msg("can not initialize application")
	}

	a.Start(ctx, cli)

	return a, nil
}
