package log

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

// Logger is a struct of custom logger.
type Logger struct {
	logrus.Logger
	logger *logrus.Logger
}

// New creates a new logger.
func New(ctx context.Context) *Logger {
	return &Logger{
		logger: logrus.New(),
	}
}

// Debug is a standard output for debug.
func (l *Logger) Debug(msg string) {
	fmt.Println(msg)
	logrus.Debug(msg)
}
