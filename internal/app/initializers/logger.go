package initializers

import (
	krakendlogs "github.com/devopsfaith/krakend/logging"
	"github.com/gobuffalo/envy"
	"github.com/wajox/krakend-proxy/internal/logging"
)

func InitializeLogger() (krakendlogs.Logger, error) {
	return logging.NewLogger(envy.Get("KRAKEND_LOG_LEVEL", "DEBUG"))
}
