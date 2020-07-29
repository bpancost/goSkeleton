package logging

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var once sync.Once

type LogrusLogger struct {
	logger *logrus.Entry
}

// While this will ensure this is always *some* logger, you are expected to call GetLoggerWithBaseFields at least once
// in the application setup.
var log = &LogrusLogger{
	logger: logrus.New().WithFields(logrus.Fields{}),
}

func GetLoggerWithBaseFields(projectName string, hostName string) Logger {
	fields := logrus.Fields{
		"project_name": projectName,
		"host_name":    hostName,
	}
	once.Do(func() {
		log = &LogrusLogger{
			logger: logrus.New().WithFields(fields),
		}
	})
	return log
}

func Trace(args ...interface{}) {
	log.Trace(args)
}
func (log *LogrusLogger) Trace(args ...interface{}) {
	log.logger.Trace(args)
}

func Tracef(format string, args ...interface{}) {
	log.Tracef(format, args)
}
func (log *LogrusLogger) Tracef(format string, args ...interface{}) {
	log.logger.Tracef(format, args)
}

func Debug(args ...interface{}) {
	log.Debug(args)
}
func (log *LogrusLogger) Debug(args ...interface{}) {
	log.logger.Debug(args)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args)
}
func (log *LogrusLogger) Debugf(format string, args ...interface{}) {
	log.logger.Debugf(format, args)
}

func Info(args ...interface{}) {
	log.Info(args)
}
func (log *LogrusLogger) Info(args ...interface{}) {
	log.logger.Info(args)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args)
}
func (log *LogrusLogger) Infof(format string, args ...interface{}) {
	log.logger.Infof(format, args)
}

func Warn(args ...interface{}) {
	log.Warn(args)
}
func (log *LogrusLogger) Warn(args ...interface{}) {
	log.logger.Warn(args)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args)
}
func (log *LogrusLogger) Warnf(format string, args ...interface{}) {
	log.logger.Warnf(format, args)
}

func Error(args ...interface{}) {
	log.Error(args)
}
func (log *LogrusLogger) Error(args ...interface{}) {
	log.logger.Error(args)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args)
}
func (log *LogrusLogger) Errorf(format string, args ...interface{}) {
	log.logger.Errorf(format, args)
}

func Fatal(args ...interface{}) {
	log.Fatal(args)
}
func (log *LogrusLogger) Fatal(args ...interface{}) {
	log.logger.Fatal(args)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args)
}
func (log *LogrusLogger) Fatalf(format string, args ...interface{}) {
	log.logger.Fatalf(format, args)
}

func Panic(args ...interface{}) {
	log.Panic(args)
}
func (log *LogrusLogger) Panic(args ...interface{}) {
	log.logger.Panic(args)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args)
}
func (log *LogrusLogger) Panicf(format string, args ...interface{}) {
	log.logger.Panicf(format, args)
}
