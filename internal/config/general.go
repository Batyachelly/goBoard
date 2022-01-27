package config

type General struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"DEBUG"`
}
