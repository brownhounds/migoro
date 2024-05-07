package query

import (
	"migoro/types"
	"migoro/utils"
	"os"

	"github.com/jmoiron/sqlx"
)

func Exists(connection *sqlx.DB, query string) types.DbCheck {
	r := types.DbCheck{}
	err := connection.Get(&r, query)
	if err != nil {
		utils.Error("Executing query", err.Error())
		connection.Close()
		os.Exit(1)
	}
	connection.Close()
	return r
}

func Query(connection *sqlx.DB, query string) {
	_, err := connection.Exec(query)
	if err != nil {
		utils.Error("Executing query", err.Error())
		connection.Close()
		os.Exit(1)
	}
	connection.Close()
}

func GetMigrations(connection *sqlx.DB, query string) []types.Migration {
	m := []types.Migration{}
	err := connection.Select(&m, query)
	if err != nil {
		utils.Error("Getting migration log results", err.Error())
		connection.Close()
		os.Exit(1)
	}
	connection.Close()
	return m
}

func WriteMigrationLog(connection *sqlx.DB, query string, file string, hash string) {
	_, err := connection.Exec(query, file, hash)
	if err != nil {
		utils.Error("Insert into migration table", err.Error())
		connection.Close()
		os.Exit(1)
	}
	connection.Close()
}

func CleanMigrationLog(connection *sqlx.DB, query string, file string) {
	_, err := connection.Exec(query, file)
	if err != nil {
		utils.Error("Remove from migration table", err.Error())
		connection.Close()
		os.Exit(1)
	}
	connection.Close()
}

func ApplyMigration(connection *sqlx.DB, query string) {
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
