package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}
type AuthConfig struct {
	Secret string
}
type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
}
