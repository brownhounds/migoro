package query

import (
	"migoro/adapters"
	"migoro/utils"
	"os"
)

type Migration struct {
	MigrationFile string `db:"migration_file"`
}

type DbCheck struct {
	Exists bool `db:"test"`
}

func Exists(query string) DbCheck {
	connection := adapters.Init().Connection()
	r := DbCheck{}
	err := connection.Get(&r, query)
	if err != nil {
		utils.Error("Executing query", err.Error())
		connection.Close()
		os.Exit(1)
	}
	connection.Close()
	return r
}

func Query(query string) {
	connection := adapters.Init().Connection()
	_, err := connection.Exec(query)
	if err != nil {
		utils.Error("Executing query", err.Error())
		connection.Close()
		os.Exit(1)
	}
	connection.Close()
}

func GetMigrations(query string) []Migration {
	connection := adapters.Init().Connection()
	m := []Migration{}
	err := connection.Select(&m, query)
	if err != nil {
		utils.Error("Getting migration log results", err.Error())
		connection.Close()
		os.Exit(1)
	}
	connection.Close()
	return m
}

func WriteMigrationLog(query string, file string, hash string) {
	connection := adapters.Init().Connection()
	_, err := connection.Exec(query, file, hash)
	if err != nil {
		utils.Error("Insert into migration table", err.Error())
		connection.Close()
		os.Exit(1)
	}
	connection.Close()
}

func CleanMigrationLog(query string, file string) {
	connection := adapters.Init().Connection()
	_, err := connection.Exec(query, file)
	if err != nil {
		utils.Error("Remove from migration table", err.Error())
		connection.Close()
		os.Exit(1)
	}
	connection.Close()
}

func ApplyMigration(query string) {
	connection := adapters.Init().Connection()
	tx, err := connection.Begin()
	if err != nil {
		utils.Error("Applying migration", err.Error())
		connection.Close()
		os.Exit(1)
	}
	{
		_, err := connection.Exec(query)
		if err != nil {
			tx.Rollback()
			utils.Error("Applying migration", err.Error())
			connection.Close()
			os.Exit(1)
		}
	}

	defer tx.Commit()
	defer connection.Close()
}
