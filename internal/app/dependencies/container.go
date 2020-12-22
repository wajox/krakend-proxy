package dependencies

import (
	krakendcfg "github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
)

// Container is a DI container for application.
type Container struct {
	Cfg    *krakendcfg.ServiceConfig
	Logger logging.Logger
}
