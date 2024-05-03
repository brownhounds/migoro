package utils

import (
	"os"

	"github.com/joho/godotenv"
)


func Env(k string) string {
	err := godotenv.Load()
	if err != nil {
		Warning("Loading .env file", err.Error())
	}
	return os.Getenv(k)
}
