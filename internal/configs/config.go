package configs

type Config struct {
	Addr string
	DB   DBConfig
	Env  string
}

type DBConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}
