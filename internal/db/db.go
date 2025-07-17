package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/PrinceNarteh/social/internal/config"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.Envs.DBConfig.DBAddr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.Envs.DBConfig.DBMaxOpenConns)
	db.SetMaxIdleConns(config.Envs.DBConfig.DBMaxIdleConns)
	
	duration, err := time.ParseDuration(config.Envs.DBConfig.DBMaxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	log.Println("database connection pool established successfully!")
	return db, nil
}
