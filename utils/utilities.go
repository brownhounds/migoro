package utils

import (
	"crypto/rand"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// CreateMigration makes migration
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

	err := os.WriteFile(f, t, 0644)
	if err != nil {
		Error("Creating migration file", err.Error())
		os.Exit(1)
	}
}

func getMigrationsPath() string {
	return Env("MIGRATION_DIR") + "/"
}

// CreateDirIfNotExist creates directory if dosent exist
func CreateDirIfNotExist() {
	if _, err := os.Stat(Env("MIGRATION_DIR")); os.IsNotExist(err) {
		err = os.MkdirAll(Env("MIGRATION_DIR"), 0755)
		if err != nil {
			Error("Creating migration directory", err.Error())
			os.Exit(1)
		}
		Success("Creating migration directory", "Created")
	}
}

// IOReadDir reads contents of the directory and returns all the files
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

// Exists checks if given file or directory exists
func Exists(f string) bool {
	if _, err := os.Stat(getMigrationsPath() + f); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// InSliceOfStructs checks if provided value exists in slice of structs
func InSliceOfStructs(slice interface{}, field string, value string) bool {
	r := reflect.ValueOf(slice)
	for i := 0; i < r.Len(); i++ {
		s := r.Index(i)
		f := s.FieldByName(field)
		if f.IsValid() {
			if f.Interface() == value {
				return true
			}
		}
	}
	return false
}

// GetFileContent get contents of the file
func GetFileContent(f string) string {
	c, err := os.ReadFile(getMigrationsPath() + f)
	if err != nil {
		Error("Reading migration file", err.Error())
		os.Exit(1)
	}

	return string(c)
}

// GetStringInBetween gets string in between two given characters from the string
func GetStringInBetween(str string, start string, end string) (r string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str, end)
	return str[s:e]
}

// ValidateStringANU validates string against alpha numerical value plus underscores
func ValidateStringANU(s string) bool {
	r := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	return r.MatchString(s)
}

func truncateString(str string, num int) string {
	new := str
	if len(str) > num {
		new = str[0:num]
	}
	return new
}

// MakeRandom Token
func MakeRandom() string {
	s := make([]byte, 16)
	rand.Read(s)
	return fmt.Sprintf("%x", s)
}
