// Package slog provides a simple wrapper around log.Printf that is only executed when
// debug is set in configuration
package slog

import (
	"fmt"
	"github.com/alexandre-normand/slackscot/v2/config"
	"github.com/spf13/viper"
	"log"
)

// Debugf logs a debug line after checking if the configuration is in debug mode
func Debugf(l *log.Logger, format string, v ...interface{}) {
	if viper.GetBool(config.DebugKey) {
		l.Output(2, fmt.Sprintf(format, v...))
	}
}
