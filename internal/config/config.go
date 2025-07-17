package config

type config struct {
	AppConfig appConfig
	DBConfig  dbConfig
}

type appConfig struct {
	Addr    string
	Version string
	Env     string
}

type dbConfig struct {
	DBAddr         string
	DBMaxOpenConns int
	DBMaxIdleConns int
	DBMaxIdleTime  string
}
