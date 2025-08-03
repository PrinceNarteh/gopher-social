package main

import (
	"log"

	"github.com/PrinceNarteh/social/internal/db"
	"github.com/PrinceNarteh/social/internal/handlers"
	"github.com/PrinceNarteh/social/internal/middleware"
	"github.com/PrinceNarteh/social/internal/store"
	"github.com/PrinceNarteh/social/internal/utils"
)

//	@title			GopherSocial API
//	@description	API for GopherSocial, a social network gophers
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath	/v1
func main() {
	// Initialize the database
	db, err := db.InitDB()
	if err != nil {
		log.Panic(err)
	}

	// Initialize logger
	utils.NewLogger()

	// Initialize the store, middleware and handlers
	store := store.NewStorage(db)
	middleware.NewMiddleware(store)
	handlers := handlers.NewHandlers(store)

	// Initialize the application and mount routes
	app := NewApplication()
	r := app.mount()
	app.initRoutes(r, handlers)
	utils.Logger.Fatal(app.run(r))
}
