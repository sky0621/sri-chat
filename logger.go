package chat

import (
	"os"

	"github.com/Sirupsen/logrus"
)

// Logger ...
type Logger struct {
	Entry   *logrus.Entry
	logfile *os.File
}

// NewLogger ...
func NewLogger(config *Config) (*Logger, error) {
	logrusEntry := logrus.WithFields(logrus.Fields{
		"host":   config.Host,
		"system": "sri-chat",
	})
	logrusEntry.Logger.Formatter = new(logrus.TextFormatter)
	logrusEntry.Logger.Formatter = new(logrus.JSONFormatter) // default

	_, err := os.Stat(config.Filepath)
	var logfile *os.File
	if err == nil {
		logfile, err = os.OpenFile(config.Filepath, os.O_APPEND, 0666)
	} else {
		logfile, err = os.Create(config.Filepath)
	}
	if err != nil {
		return nil, err
	}
	logrusEntry.Logger.Out = logfile

	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return nil, err
	}
	logrusEntry.Logger.Level = level

	return &Logger{Entry: logrusEntry, logfile: logfile}, nil
}

// Close ...
func (l *Logger) Close() {
	l.logfile.Close()
}
