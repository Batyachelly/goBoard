package http

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/Batyachelly/goBoard/generated/swagger" // docs is generated by Swag CLI
	"github.com/Batyachelly/goBoard/internal/logger"
	"github.com/Batyachelly/goBoard/internal/usecase"

	"github.com/gorilla/mux"
	swagger "github.com/swaggo/http-swagger"
)

type Serve interface {
	Serve() error
}

type Server struct {
	server  *http.Server
	usecase usecase.Usecaser
	log     logger.Logger
}

type Config struct {
	Addr         string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	Log          logger.Logger
}

func NewServer(cfg Config, usecase usecase.Usecaser) *Server {
	r := mux.NewRouter()

	s := &Server{
		usecase: usecase,
		server: &http.Server{
			Handler: r,

			Addr:         cfg.Addr,
			WriteTimeout: cfg.WriteTimeout,
			ReadTimeout:  cfg.ReadTimeout,
		},
		log: cfg.Log,
	}

	r.Use(CaptchaVerify)

	sub := r.PathPrefix("/api/v1").Subrouter()

	sub.HandleFunc("/board", s.GetBoards).Methods(http.MethodGet)
	sub.HandleFunc("/board/{board_id}", s.GetBoard).Methods(http.MethodGet)
	sub.HandleFunc("/board/{board_id}/thread/{thread_id}", s.GetThread).Methods(http.MethodGet)

	sub.HandleFunc("/board/{board_id}/thread", s.PostThread).Methods(http.MethodPost)
	sub.HandleFunc("/board/{board_id}/thread/{thread_id}/comment", s.PostMessage).Methods(http.MethodPost)

	r.PathPrefix("/swagger/").Handler(swagger.Handler(
		swagger.URL("doc.json"),
	))

	return s
}

func (s *Server) Serve() error {
	return fmt.Errorf("serve http: %w", s.server.ListenAndServe())
}
