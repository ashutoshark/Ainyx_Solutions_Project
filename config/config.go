package config

import "os"

// config struct for app settings
type Config struct {
	Port        string
	DatabaseURL string
}

// this function loads the config
func Load() *Config {
	// get port
	p := os.Getenv("SERVER_PORT")
	if p == "" {
		p = "3000"
	}

	// get db url
	db_url := os.Getenv("DATABASE_URL")
	if db_url == "" {
		db_url = "postgres://postgres:ashutosh@localhost:5432/userdb?sslmode=disable"
	}

	return &Config{
		Port:        p,
		DatabaseURL: db_url,
	}
}
