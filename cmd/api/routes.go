package main

import (
	chi "github.com/go-chi/chi/v5"

	"github.com/PrinceNarteh/social/internal/handlers"
)

func (app *application) initRoutes(r *chi.Mux, h *handlers.Handler) {
	r.Route("/api/v1", func(r chi.Router) {
		// health check
		r.Get("/health", app.healthCheckHandler)

		// auth routes
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", h.Auth.LoginHandler)
			r.Post("/register", h.Auth.RegisterHandler)
		})

		// posts routes
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", h.Post.GetAllPostsHandler)
			r.Post("/", h.Post.CreatePostHandler)

			r.Route("/{postId}", func(r chi.Router) {
				r.Use(h.Middleware.PostsContextMiddlware)

				r.Get("/", h.Post.GetPostHandler)
				r.Patch("/", h.Post.UpdatePostHandler)
				r.Delete("/", h.Post.DeletePostHandler)
			})
		})
	})
}
