package goboard

import (
	"log"

	"github.com/Batyachelly/goBoard/internal/config"
	"github.com/Batyachelly/goBoard/internal/database/pg"
	"github.com/Batyachelly/goBoard/internal/logger/logrus"
	"github.com/Batyachelly/goBoard/internal/transport/http"
	"github.com/Batyachelly/goBoard/internal/usecase"
)

type App struct {
	httpServer http.Serve
}

// @title           GoBoard API
// @version         0.1
// @description     This is a goBoard server.

// @host      localhost:8080
// @BasePath  /api/v1
func Serve(migrate bool) {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	logLib, err := logrus.New(logrus.Config{
		Level: cfg.General.LogLevel,
	})
	if err != nil {
		log.Fatal(err)
	}

	databaseService, err := pg.NewDatabaseService(pg.Config{
		Host:         cfg.Postgres.Host,
		Port:         cfg.Postgres.Port,
		User:         cfg.Postgres.User,
		Password:     cfg.Postgres.Password,
		DB:           cfg.Postgres.DB,
		SSLMode:      cfg.Postgres.SSLMode,
		VersionTable: cfg.Postgres.VersionTable,
	})
	if err != nil {
		logLib.Fatal("%v", err)
	}

	if migrate {
		if err := databaseService.Migrate(); err != nil {
			logLib.Fatal("%v", err)
		}
	}

	uc := usecase.NewUsecase(databaseService)
	hs := http.NewServer(http.Config{
		Addr:         cfg.HTTP.Addr,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		Log:          logLib,
	}, uc)
	app := App{httpServer: hs}

	logLib.Fatal("%v", app.httpServer.Serve())
}

func Migrate() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	logLib, err := logrus.New(logrus.Config{
		Level: cfg.General.LogLevel,
	})
	if err != nil {
		log.Fatal(err)
	}

	databaseService, err := pg.NewDatabaseService(pg.Config{
		Host:         cfg.Postgres.Host,
		Port:         cfg.Postgres.Port,
		User:         cfg.Postgres.User,
		Password:     cfg.Postgres.Password,
		DB:           cfg.Postgres.DB,
		SSLMode:      cfg.Postgres.SSLMode,
		VersionTable: cfg.Postgres.VersionTable,
	})
	if err != nil {
		logLib.Fatal("%v", err)
	}

	if err := databaseService.Migrate(); err != nil {
		logLib.Fatal("%v", err)
	}
}
