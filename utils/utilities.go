package utils

import (
	"crypto/rand"
	"fmt"
	"migoro/types"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CreateMigration(n string) {
	t := truncateString(strconv.FormatInt(time.Now().UnixNano(), 10), 12)
	CreateDirIfNotExist()

	fn := t + "_" + n + ".sql"
	f := getMigrationsPath() + fn
	createMigrationFile(f)
	Success("Migration Created", fn)
}

func createMigrationFile(f string) {
	t := []byte("/* UP-START */\n\n/* UP-END */\n/* DOWN-START */\n\n/* DOWN-END */")

	err := os.WriteFile(f, t, 0o644)
	if err != nil {
		Error("Creating migration file", err.Error())
		os.Exit(1)
	}
}

func getMigrationsPath() string {
	return Env("MIGRATION_DIR") + "/"
}

func CreateDirIfNotExist() {
	if _, err := os.Stat(Env("MIGRATION_DIR")); os.IsNotExist(err) {
		err = os.MkdirAll(Env("MIGRATION_DIR"), 0o755)
		if err != nil {
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

func GetFileContent(f string) string {
	c, err := os.ReadFile(getMigrationsPath() + f)
	if err != nil {
		Error("Reading migration file", err.Error())
		os.Exit(1)
	}

	return string(c)
}

func GetStringInBetween(str, start, end string) (r string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str, end)
	return str[s:e]
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
