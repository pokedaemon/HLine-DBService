package api

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    "8080",
		LogLevel:    "debug",
		DatabaseURL: "host=localhost dbname=reverseshoptest sslmode=disable",
	}
}
