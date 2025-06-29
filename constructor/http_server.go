package constructor

import (
	"context"
	"net/http"
	"time"
)

type Server interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type apiServer struct {
	srv *http.Server
}

type Config struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Handler      http.Handler
}

func NewServer(cfg Config) Server {
	srv := &http.Server{
		Addr:         cfg.Addr,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		Handler:      cfg.Handler,
	}

	return &apiServer{srv: srv}
}

func (a *apiServer) ListenAndServe() error {
	return a.srv.ListenAndServe()
}

func (a *apiServer) Shutdown(ctx context.Context) error {
	return a.srv.Shutdown(ctx)
}
