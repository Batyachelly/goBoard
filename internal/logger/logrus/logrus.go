package logrus

import (
	"github.com/Batyachelly/goBoard/internal/logger"

	log "github.com/sirupsen/logrus"
)

type Logrus struct {
	sl *log.Logger
}

type Config struct {
	Level string
}

func New(cfg Config) (*Logrus, error) {
	logrus := new(Logrus)

	logrus.sl = log.New()

	switch cfg.Level {
	case logger.DebugLevel:
		logrus.sl.SetLevel(log.DebugLevel)
	case logger.InfoLevel:
		logrus.sl.SetLevel(log.InfoLevel)
	case logger.WarnLevel:
		logrus.sl.SetLevel(log.WarnLevel)
	case logger.ErrorLevel:
		logrus.sl.SetLevel(log.ErrorLevel)
	case logger.PanicLevel:
		logrus.sl.SetLevel(log.PanicLevel)
	case logger.FatalLevel:
		logrus.sl.SetLevel(log.FatalLevel)
	}

	return logrus, nil
}

func (l *Logrus) Info(template string, args ...interface{}) {
	l.sl.Infof(template, args...)
}

func (l *Logrus) Debug(template string, args ...interface{}) {
	l.sl.Debugf(template, args...)
}

func (l *Logrus) Error(template string, args ...interface{}) {
	l.sl.Errorf(template, args...)
}

func (l *Logrus) Fatal(template string, args ...interface{}) {
	l.sl.Fatalf(template, args...)
}
