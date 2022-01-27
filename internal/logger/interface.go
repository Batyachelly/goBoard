package logger

const (
	DebugLevel = "DEBUG"
	InfoLevel  = "INFO"
	WarnLevel  = "WARN"
	ErrorLevel = "ERROR"
	PanicLevel = "PANIC"
	FatalLevel = "FATAL"
)

type Logger interface {
	Info(string, ...interface{})
	Debug(string, ...interface{})
	Error(string, ...interface{})
	Fatal(string, ...interface{})
}

type TestLogger struct{}

func (t TestLogger) Info(_ string, _ ...interface{}) {
}

func (t TestLogger) Debug(_ string, _ ...interface{}) {
}

func (t TestLogger) Error(_ string, _ ...interface{}) {
}

func (t TestLogger) Fatal(_ string, _ ...interface{}) {
}
