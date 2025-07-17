package main

import (
	"log"
	"net/http"
	"time"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/PrinceNarteh/social/internal/config"
)

type application struct{}

func NewApplication() *application {
	return &application{}
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	return r
}

func (app *application) run(handler *chi.Mux) error {
	srv := http.Server{
		Addr:         config.Envs.AppConfig.Addr,
		Handler:      handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server running on %s", config.Envs.AppConfig.Addr)
	return srv.ListenAndServe()
}
