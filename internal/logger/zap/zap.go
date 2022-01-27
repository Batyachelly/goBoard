package zap

import (
	"fmt"

	"github.com/Batyachelly/goBoard/internal/logger"

	"go.uber.org/zap"
)

type Zaplog struct {
	sl *zap.SugaredLogger
}

type Config struct {
	Level string
}

func New(cfg Config) (*Zaplog, error) {
	zaplog := new(Zaplog)

	switch cfg.Level {
	case logger.DebugLevel:
		devLog, err := zap.NewDevelopment()
		if err != nil {
			return nil, fmt.Errorf("zap new development: %w", err)
		}

		zaplog.sl = devLog.Sugar()
	case logger.InfoLevel:
		prodLog, err := zap.NewProduction()
		if err != nil {
			return nil, fmt.Errorf("zap new development: %w", err)
		}

		zaplog.sl = prodLog.Sugar()
	default:
		zaplog.sl = zap.NewExample().Sugar()
	}

	return zaplog, nil
}

func (z *Zaplog) Info(template string, args ...interface{}) {
	z.sl.Infof(template, args...)
}

func (z *Zaplog) Debug(template string, args ...interface{}) {
	z.sl.Debugf(template, args...)
}

func (z *Zaplog) Error(template string, args ...interface{}) {
	z.sl.Errorf(template, args...)
}

func (z *Zaplog) Fatal(template string, args ...interface{}) {
	z.sl.Fatalf(template, args...)
}
