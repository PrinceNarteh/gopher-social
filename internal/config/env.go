package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

func initConfig() *config {
	godotenv.Load()

	return &config{
		AppConfig: appConfig{
			Env:     getEnvString("APP_ENV", "developement"),
			Addr:    fmt.Sprintf(":%s", getEnvString("APP_HOST", "4000")),
			Version: getEnvString("APP_VER", "0.0.2"),
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
