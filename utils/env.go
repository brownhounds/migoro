package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func Env(k string) string {
	godotenv.Load() //nolint
	return os.Getenv(k)
}
