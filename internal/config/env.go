package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

func initConfig() *config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("could not load .env file")
	}

	return &config{
		AppConfig: appConfig{
			Env:         getEnvString("APP_ENV", "developement"),
			Addr:        fmt.Sprintf(":%s", getEnvString("APP_HOST", "4000")),
			Version:     getEnvString("APP_VER", "0.0.2"),
			FrontendURL: getEnvString("FRONTEND_URL", "http://localhost:4000"),
		},
		DBConfig: dbConfig{
			DBAddr: getEnvString(
				"DB_ADDR",
				"postgres://admin:admin_secret@localhost:5432/social?sslmode=disable",
			),
			DBMaxOpenConns: getEnvInt("DB_MAX_OPEN_CONNS", 10),
			DBMaxIdleConns: getEnvInt("DB_MAX_IDLE_CONNS", 10),
			DBMaxIdleTime:  getEnvString("DB_MAX_IDLE_TIME", "10m"),
		},
		MailConfig: mailConfig{
			FromEmail: getEnvString("MAILER_FROM_EMAIL", "donprinart@gmail.com"),
			ApiKey:    getEnvString("MAILER_API_KEY", "1234567890"),
			Exp:       time.Hour * 24 * 3, // 3 days
		},
	}
}

func getEnvString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return valueInt
}
