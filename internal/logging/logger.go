package logging

import (
	"context"
	"net/http"
)

type Logger interface {
	AddRestRequestLogger(req *http.Request) *http.Request
	GetRestRequestLogger(req *http.Request) Logger
	AddGrpcContextLogger(ctx context.Context, fullMethod string) context.Context
	GetGrpcContextLogger(ctx context.Context) Logger
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
}
