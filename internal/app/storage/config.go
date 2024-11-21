package storage

type Config struct {
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		DatabaseURL: "host=localhost dbname=reverseshop sslmode=disable",
	}
}
