package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func Env(k string) string {
	godotenv.Load()
	return os.Getenv(k)
}
