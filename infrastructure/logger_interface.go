package infrastructure

import (
	"github.com/sirupsen/logrus"
)

// LoggerInterface is main interface for logger
type LoggerInterface interface {
	// // Warnf creates a string formatted log at the debug level.
	// Warnf(format string, args ...interface{})

	// // Infof creates a string formatted log at the debug level.
	// Infof(format string, args ...interface{})

	// // Debug calls debug on the package level logger.
	// Debug(args ...interface{})

	// Info calls Info on the package level logger.
	Info(args ...interface{})

	// // Print calls Print on the package level logger.
	// Print(args ...interface{})

	// Warn calls Warn on the package level logger.
	Warn(args ...interface{})
}
type Log struct {
	*logrus.Logger
}

// Info calls Info on the package level logger.
func (log *Log) Info(args ...interface{}) {
	log.Info(args...)
}

// Info calls Info on the package level logger.
func (log *Log) Warn(args ...interface{}) {
	log.Warn(args...)
}
