package logging

import (
	"fmt"
	"strings"

	krakendlogs "github.com/devopsfaith/krakend/logging"
	"github.com/rs/zerolog/log"
)

var (
	defaultLogger = logger{Level: krakendlogs.LEVEL_CRITICAL}
	logLevels     = map[string]int{
		"DEBUG":    krakendlogs.LEVEL_DEBUG,
		"INFO":     krakendlogs.LEVEL_INFO,
		"WARNING":  krakendlogs.LEVEL_WARNING,
		"ERROR":    krakendlogs.LEVEL_ERROR,
		"CRITICAL": krakendlogs.LEVEL_CRITICAL,
	}
	// NoOp is the NO-OP logger.
	NoOp, _ = NewLogger("CRITICAL")
)

// NewLogger creates and returns a Logger object.
func NewLogger(level string) (krakendlogs.Logger, error) {
	l, ok := logLevels[strings.ToUpper(level)]
	if !ok {
		return defaultLogger, krakendlogs.ErrInvalidLogLevel
	}

	return logger{Level: l}, nil // x return krakendlogs.NewLogger("DEBUG", os.Stdout, "[KRAKEND]")
}

type logger struct {
	Level int
}

// Debug logs a message using DEBUG as log level.
func (l logger) Debug(v ...interface{}) {
	if l.Level > krakendlogs.LEVEL_DEBUG {
		return
	}

	log.Debug().Msg(fmt.Sprintln(v...))
}

// Info logs a message using INFO as log level.
func (l logger) Info(v ...interface{}) {
	if l.Level > krakendlogs.LEVEL_INFO {
		return
	}

	log.Info().Msg(fmt.Sprintln(v...))
}

// Warning logs a message using WARNING as log level.
func (l logger) Warning(v ...interface{}) {
	if l.Level > krakendlogs.LEVEL_WARNING {
		return
	}

	log.Warn().Msg(fmt.Sprintln(v...))
}

// Error logs a message using ERROR as log level.
func (l logger) Error(v ...interface{}) {
	if l.Level > krakendlogs.LEVEL_ERROR {
		return
	}

	log.Error().Msg(fmt.Sprintln(v...))
}

// Critical logs a message using CRITICAL as log level.
func (l logger) Critical(v ...interface{}) {
	log.Error().Msg(fmt.Sprintln(v...))
}

// Fatal is equivalent to l.Critical(fmt.Sprint()) followed by a call to os.Exit(1).
func (l logger) Fatal(v ...interface{}) {
	log.Panic().Msg(fmt.Sprintln(v...))
}
