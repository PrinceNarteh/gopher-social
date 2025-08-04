package main

import (
	"fmt"

	chi "github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/PrinceNarteh/social/internal/handlers"
)

func (app *application) initRoutes(r *chi.Mux, h *handlers.Handler) {
	r.Route("/api/v1", func(r chi.Router) {
		// health check
		r.Get("/health", app.healthCheckHandler)

		// swagger
		docsURL := fmt.Sprintf("%s/swagger/doc.json", "4000")
		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))

		// auth routes
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", h.Auth.LoginHandler)
			r.Post("/register", h.Auth.RegisterHandler)
		})

		// users routes
		r.Route("/users", func(r chi.Router) {
			r.Get("/activate/{token}", h.User.ActivateUserHandler)

			r.Route("/{userId}", func(r chi.Router) {
				r.Use(h.Middleware.UserContextMiddleware)

				r.Get("/", h.User.GetUserHandler)
				r.Patch("/", h.User.UpdateUserHandler)
				r.Delete("/", h.User.DeleteUserHandler)

				// follow and unfollow
				r.Put("/follow", h.User.FollowHandler)
				r.Put("/unfollow", h.User.UnfollowHandler)
			})

			r.Group(func(r chi.Router) {
				r.Get("/feed", h.Feed.GetUserFeedHandler)
			})
		})

		// posts routes
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", h.Post.GetAllPostsHandler)
			r.Post("/", h.Post.CreatePostHandler)

			r.Route("/{postId}", func(r chi.Router) {
				r.Use(h.Middleware.PostContextMiddlware)

				r.Get("/", h.Post.GetPostHandler)
				r.Patch("/", h.Post.UpdatePostHandler)
				r.Delete("/", h.Post.DeletePostHandler)
			})
		})
	})
}
