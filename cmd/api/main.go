package main

import (
	"fmt"
	"log"

	"github.com/PrinceNarteh/gopher-social/internal/configs"
	"github.com/PrinceNarteh/gopher-social/internal/db"
	"github.com/PrinceNarteh/gopher-social/internal/env"
	"github.com/PrinceNarteh/gopher-social/internal/repositories"
)

const version = "0.0.1"

func main() {
	// load enviroment variables
	env.LoadEnv()

	// configure enviroment variables
	cfg := configs.Config{
		Addr: fmt.Sprintf(":%s", env.GetStringEnv("PORT", "8080")),
		DB: configs.DBConfig{
			Addr: env.GetStringEnv(
				"DB_URI",
				"postgres://admin:secret_password@localhost/social?sslmode=disable",
			),
			MaxOpenConns: env.GetIntEnv("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: env.GetIntEnv("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  env.GetStringEnv("DB_MAX_IDLE_TIME", "15m"),
		},
		Env: env.GetStringEnv("ENV", "development"),
	}

	// connect to database
	db, err := db.New(cfg.DB.Addr, cfg.DB.MaxOpenConns, cfg.DB.MaxIdleConns, cfg.DB.MaxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	// initiate repositories
	repo := repositories.NewRepositories(db)
	app := &application{
		config: cfg,
		repo:   repo,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
