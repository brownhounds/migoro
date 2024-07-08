package utils

import (
	"crypto/rand"
	"fmt"
	"migoro/types"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	MIGRATION_DIR = "MIGRATION_DIR"
	UP            = "up"
	DOWN          = "down"
)

func CreateMigrationFile(name string) {
	CreateDirIfNotExist()

	t := truncateString(strconv.FormatInt(time.Now().UnixNano(), 10), 12)

	up := getMigrationFileName(t, name, UP)
	down := getMigrationFileName(t, name, DOWN)

	createMigrationFile(up)
	createMigrationFile(down)

	Success("Migration File Created", up)
	Success("Migration File Created", down)
}

func getMigrationFileName(t, name, method string) string {
	return getMigrationsPath() + t + "_" + name + "_" + method + ".sql"
}

func createMigrationFile(f string) {
	if err := os.WriteFile(f, make([]byte, 0), 0o644); err != nil {
		Error("Creating migration file", err.Error())
		os.Exit(1)
	}
}

func getMigrationsPath() string {
	return strings.TrimSuffix(Env(MIGRATION_DIR), "/") + "/"
}

func CreateDirIfNotExist() {
	if _, err := os.Stat(Env(MIGRATION_DIR)); os.IsNotExist(err) {
		if err = os.MkdirAll(Env(MIGRATION_DIR), 0o755); err != nil {
			Error("Creating migration directory", err.Error())
			os.Exit(1)
		}
		Success("Creating migration directory", "Created")
	}
}

func IOReadDir(root string) []string {
	var files []string
	fileInfo, err := os.ReadDir(root)
	if err != nil {
		Error("Reading migrations directory", err.Error())
		os.Exit(1)
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files
}

func Exists(f string) bool {
	if _, err := os.Stat(getMigrationsPath() + f); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func InSliceOfStructs(slice []types.Migration, value string) bool {
	file := types.Migration{MigrationFile: value}
	for _, migration := range slice {
		if migration == file {
			return true
		}
	}

	return false
}

func stripSQLComments(c []byte) string {
	partsOut := make([]string, 0)
	parts := strings.Split(string(c), "\n")

	for _, line := range parts {
		if !strings.HasPrefix(strings.TrimSpace(line), "--") {
			partsOut = append(partsOut, line)
		}
	}

	return strings.Join(partsOut, "\n")
}

func GetMigrationFileContent(f string) string {
	c, err := os.ReadFile(getMigrationsPath() + f)
	if err != nil {
		Error("Reading migration file", err.Error())
		os.Exit(1)
	}

	return stripSQLComments(c)
}

func ValidateStringANU(s string) bool {
	r := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	return r.MatchString(s)
}

func truncateString(str string, num int) string {
	newStr := str
	if len(str) > num {
		newStr = str[0:num]
	}
	return newStr
}

func MakeRandom() string {
	s := make([]byte, 16)
	if _, err := rand.Read(s); err != nil {
		Error("MakeRandom", fmt.Sprintf("No able to generate a random string: %s", err.Error()))
		os.Exit(1)
	}
	return fmt.Sprintf("%x", s)
}

func ValidateEnvVariables(envVars []string) {
	for _, value := range envVars {
		_, defined := os.LookupEnv(value)
		if !defined {
			panic("ENV Variable is not defined: " + value)
		}
	}
}

func GetMigrationFiles() []string {
	dir, err := os.Getwd()
	if err != nil {
		Error("Reading Current Working Directory", err.Error())
		os.Exit(1)
	}
	return IOReadDir(filepath.Join(dir, Env("MIGRATION_DIR")))
}
