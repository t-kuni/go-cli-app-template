package system

import (
	"os"

	"github.com/rotisserie/eris"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
	"github.com/t-kuni/go-cli-app-template/domain/infrastructure/system"
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger(i *do.Injector) (system.ILogger, error) {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: false,
	})

	return &Logger{
		logger: logger,
	}, nil
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *Logger) WithField(key string, value interface{}) system.ILogger {
	return &Logger{
		logger: l.logger.WithField(key, value).Logger,
	}
}

func (l *Logger) WithFields(fields map[string]interface{}) system.ILogger {
	return &Logger{
		logger: l.logger.WithFields(fields).Logger,
	}
}

func (l *Logger) WithError(err error) system.ILogger {
	// erisのスタックトレースを取得
	stackTrace := eris.ToString(err, true)
	return &Logger{
		logger: l.logger.WithFields(logrus.Fields{
			"error":      err.Error(),
			"stacktrace": stackTrace,
		}).Logger,
	}
}
