package main

import (
	"net/http"
	"time"

	"github.com/aver343/blog/pkg/config"
	"github.com/aver343/blog/pkg/db/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type application struct {
	Config     *config.Config
	Repository repository.Repository
	Router     http.Handler
	Logger     *zap.SugaredLogger
}

// Initialize the application
func NewApplication(cfg *config.Config, repo repository.Repository, logger *zap.SugaredLogger) *application {
	return &application{
		Config:     cfg,
		Repository: repo,
		Logger:     logger,
	}
}

func (app *application) Run(handler http.Handler) error {
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      handler,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	return server.ListenAndServe()
}

func (app *application) Mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/v1", func(r chi.Router) {
		r.Route("/posts", postHandler(app))
		r.Route("/users", userHandler(app))
	})
	return r
}
