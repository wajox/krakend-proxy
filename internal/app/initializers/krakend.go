package initializers

import (
	"context"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	"github.com/devopsfaith/krakend/router"
	krakendgin "github.com/devopsfaith/krakend/router/gin"
	"github.com/gin-gonic/gin"
	ginmw "github.com/wajox/gin-ext-lib/middleware"
)

//TODO replace it with env/config value
const maxAllowed = 20

func InitializeKrakenD(
	ctx context.Context,
	serviceConfig *config.ServiceConfig,
	logger logging.Logger,
) router.Router {
	return newGinRouter(
		ctx,
		newMws(serviceConfig),
		logger,
	)
}

func newMws(cfg *config.ServiceConfig) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		limit.MaxAllowed(maxAllowed),
	}
}

func newGinRouter(
	ctx context.Context,
	mws []gin.HandlerFunc,
	logger logging.Logger,
) router.Router {
	engine := gin.New()
	engine.Use(ginmw.Recovery())

	routerFactory := krakendgin.NewFactory(krakendgin.Config{
		Engine:         engine,
		ProxyFactory:   customProxyFactory{logger, proxy.DefaultFactory(logger)},
		Middlewares:    mws,
		Logger:         logger,
		HandlerFactory: krakendgin.EndpointHandler,
		RunServer:      router.RunServer,
	})

	return routerFactory.NewWithContext(ctx)
}

// customProxyFactory adds a logging middleware wrapping the internal factory.
type customProxyFactory struct {
	logger  logging.Logger
	factory proxy.Factory
}

// New implements the Factory interface.
func (cf customProxyFactory) New(cfg *config.EndpointConfig) (p proxy.Proxy, err error) {
	p, err = cf.factory.New(cfg)
	if err == nil {
		p = proxy.NewLoggingMiddleware(cf.logger, cfg.Endpoint)(p)
	}

	return
}
