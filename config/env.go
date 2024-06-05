package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUrl  string
	Port   string
	AppEnv string
}

var Env = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		DbUrl:  os.Getenv("DB_URL"),
		Port:   os.Getenv("PORT"),
		AppEnv: os.Getenv(""),
	}
}
