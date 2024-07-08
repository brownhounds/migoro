package utils

import (
	"os"
	"path"

	"github.com/joho/godotenv"
)

func Env(k string) string {
	dir, err := os.Getwd()
	if err != nil {
		Error("Reading Current Working Directory", err.Error())
		os.Exit(1)
	}
	godotenv.Load(path.Join(dir, ".env")) //nolint
	return os.Getenv(k)
}
