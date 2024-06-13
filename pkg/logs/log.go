package logs

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	h      *log.Helper
	Logger log.Logger
)

func InitLogger() {
	Logger = NewLogger()
	h = log.NewHelper(log.With(Logger))
}
func NewLogger() log.Logger {
	return nil
}

func Debugf(template string, args ...interface{}) {
	h.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	h.Infof(template, args...)
}

func Errorf(template string, args ...interface{}) {
	h.Errorf(template, args...)
}

func InfofCtx(ctx context.Context, template string, args ...interface{}) {
	requestId := requestIdFromCtx(ctx)
	if requestId != "" {
		log.NewHelper(log.With(Logger, "", requestId)).Infof(template, args...)
	} else {
		h.Infof(template, args...)
	}
}

func ErrorfCtx(ctx context.Context, template string, args ...interface{}) {
	requestId := requestIdFromCtx(ctx)
	if requestId != "" {
		log.NewHelper(log.With(Logger, "", requestId)).Errorf(template, args...)
	} else {
		h.Errorf(template, args...)
	}
}

func FatalfCtx(ctx context.Context, template string, args ...interface{}) {
	requestId := requestIdFromCtx(ctx)
	if requestId != "" {
		log.NewHelper(log.With(Logger, "", requestId)).Fatalf(template, args...)
	} else {
		h.Fatalf(template, args...)
	}
}

func Error(args ...interface{}) {
	h.Error(args...)
}

func Warn(args ...interface{}) {
	h.Warn(args...)
}

func Debug(args ...interface{}) {
	h.Debug(args...)
}

func Info(args ...interface{}) {
	h.Info(args...)
}

func Fatal(args ...interface{}) {
	h.Fatal(args...)
}

func Warnf(template string, args ...interface{}) {
	h.Warnf(template, args...)
}

func requestIdFromCtx(ctx context.Context) string {
	if requestId, ok := ctx.Value("").(string); ok {
		return requestId
	}
	return ""
}
