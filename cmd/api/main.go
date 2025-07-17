package main

import (
	"log"

	"github.com/PrinceNarteh/social/internal/db"
	"github.com/PrinceNarteh/social/internal/handlers"
	"github.com/PrinceNarteh/social/internal/middleware"
	"github.com/PrinceNarteh/social/internal/store"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Panic(err)
	}

	store := store.NewStorage(db)
	handlers := handlers.NewHandlers(store)
	middleware.NewMiddleware(store)

	app := NewApplication()
	r := app.mount()
	app.initRoutes(r, handlers)
	app.run(r)
}
