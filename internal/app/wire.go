//+build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/wajox/krakend-proxy/internal/app/dependencies"
	"github.com/wajox/krakend-proxy/internal/app/initializers"
)

func BuildApplication() (*Application, error) {
	wire.Build(
		initializers.InitializeLogger,
		initializers.InitializeConfig,
		wire.Struct(new(dependencies.Container), "Cfg", "Logger"),
		wire.Struct(new(Application), "Container"),
	)

	return &Application{}, nil
}
