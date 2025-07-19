package main

import (
	"log"

	"github.com/PrinceNarteh/social/internal/db"
	"github.com/PrinceNarteh/social/internal/store"
)

func main() {
	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	store := store.NewStorage(dbConn)

	db.Seed(store)
}
