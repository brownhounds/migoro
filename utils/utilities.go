package utils

import (
	"crypto/rand"
	"fmt"
	"migoro/error_context"
	"migoro/types"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var TimeUnix func() int64 = time.Now().UnixNano

const (
	MIGRATION_DIR = "MIGRATION_DIR"
	UP            = "up"
	DOWN          = "down"
)

func CreateMigrationFile(name string) {
	CreateDirIfNotExist()

	t := truncateString(strconv.FormatInt(TimeUnix(), 10), 12)

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
		error_context.Context.SetError()
		Error("Creating migration file", err.Error())
		return
	}
}

func getMigrationsPath() string {
	return strings.TrimSuffix(Env(MIGRATION_DIR), "/") + "/"
}

func CreateDirIfNotExist() {
	fmt.Println("ENV VAR: " + Env(MIGRATION_DIR))
	if _, err := os.Stat(Env(MIGRATION_DIR)); os.IsNotExist(err) {
		if err = os.MkdirAll(Env(MIGRATION_DIR), 0o755); err != nil {
			error_context.Context.SetError()
			Error("Creating migration directory", err.Error())
			return
		}
		Success("Creating migration directory", "Created")
	}
}

func IOReadDir(root string) (err error, f []string) {
	var files []string
	fileInfo, err := os.ReadDir(root)
	if err != nil {
		error_context.Context.SetError()
		Error("Reading migrations directory", err.Error())
		return err, nil
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return nil, files
}

func Exists(f string) bool {
	if _, err := os.Stat(getMigrationsPath() + f); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func InSliceOfStructs(slice *[]types.Migration, value string) bool {
	file := types.Migration{MigrationFile: value}
	for _, migration := range *slice {
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

func GetMigrationFileContent(f string) (err error, contents string) {
	c, err := os.ReadFile(getMigrationsPath() + f)
	if err != nil {
		error_context.Context.SetError()
		Error("Reading migration file", err.Error())
		return err, ""
	}

	out := stripSQLComments(c)
	return nil, out
}

func ValidateStringANU(s string) bool {
	r := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	return r.MatchString(s)
}

func truncateString(str string, num int) string {
	newStr := str
	if len(str) > num {
		newStr = str[:num]
	}
	return newStr
}

func MakeRandom() (err error, random string) {
	s := make([]byte, 16)
	if _, err := rand.Read(s); err != nil {
		Error("MakeRandom", fmt.Sprintf("No able to generate a random string: %s", err.Error()))
		return err, ""
	}
	return nil, fmt.Sprintf("%x", s)
}

func EnvVarNotDefinedErrorMessage(value string) string {
	return fmt.Sprintf("ENV Variable is not defined: %s", value)
}

func ValidateEnvVariables(envVars []string) {
	for _, value := range envVars {
		_, defined := os.LookupEnv(value)
		if !defined {
			panic(EnvVarNotDefinedErrorMessage(value))
		}
	}
}
