package config

type Postgres struct {
	Host         string `env:"POSTGRES_HOST"`
	Port         int    `env:"POSTGRES_PORT"`
	User         string `env:"POSTGRES_USER"`
	Password     string `env:"POSTGRES_PASSWORD"`
	DB           string `env:"POSTGRES_DB"`
	SSLMode      string `env:"POSTGRESS_SSL"`
	VersionTable string `env:"POSTGRESS_VERSION_TABLE" envDefault:"version"`
}
