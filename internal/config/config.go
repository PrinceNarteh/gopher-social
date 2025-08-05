package config

import "time"

type config struct {
	AppConfig  appConfig
	DBConfig   dbConfig
	MailConfig mailConfig
}

type appConfig struct {
	Addr        string
	Version     string
	Env         string
	FrontendURL string
}

type dbConfig struct {
	DBAddr         string
	DBMaxOpenConns int
	DBMaxIdleConns int
	DBMaxIdleTime  string
}

type mailConfig struct {
	FromEmail string
	ApiKey    string
	Exp       time.Duration
}
