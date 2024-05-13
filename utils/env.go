package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func Env(k string) string {
	if err := godotenv.Load(); err != nil {
		Info("Loading Env File", "File not provided: .env")
	}
	return os.Getenv(k)
}
