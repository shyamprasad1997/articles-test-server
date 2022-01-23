package infrastructure

import (
	"articles-test-server/shared/utils"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	OutputStdout = "stdout"
	OutputFile   = "file"
	FormatText   = "text"
	FormatJSON   = "json"
)

// Logger struct.
type Logger struct {
	Log     *logrus.Logger
	Logfile *os.File
}

// NewLogger returns new Logger.
// repository: https://github.com/sirupsen/logrus
func NewLogger() (*Logger, error) {
	var err error
	var file *os.File

	log := logrus.New()
	// To set log level
	log.Level, err = logrus.ParseLevel(utils.Env.AppLoggerLevel)
	if err != nil {
		return nil, utils.ErrorsWrap(err, "Cannot set level")
	}
	// To set log formatter
	switch utils.Env.AppLoggerFormat {
	case FormatText:
		log.Formatter = &logrus.TextFormatter{}
	case FormatJSON:
		log.Formatter = &logrus.JSONFormatter{}
	}
	// To set output
	switch utils.Env.AppLoggerOutput {
	case OutputStdout: // output: stdout
		log.Out = os.Stdout
	case OutputFile: // output: file
		err = os.MkdirAll("logs", os.ModePerm)
		if err != nil {
			return nil, utils.ErrorsWrap(err, "Cannot make directory")
		}
		logfile := "logs/console.log"
		file, err = os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return nil, utils.ErrorsWrap(err, "can't open file")
		}
		log.Out = file

	}
	return &Logger{Log: log, Logfile: file}, nil
}
