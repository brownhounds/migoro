package query

import (
	"migoro/types"
	"migoro/utils"
	"os"

	"github.com/jmoiron/sqlx"
)

func Exists(connection *sqlx.DB, query string) types.DbCheck {
	r := types.DbCheck{}
	if err := connection.Get(&r, query); err != nil {
		utils.Error("Executing query", err.Error())
		os.Exit(1)
	}
	return r
}

func Query(connection *sqlx.DB, query string) {
	if _, err := connection.Exec(query); err != nil {
		utils.Error("Executing query", err.Error())
		os.Exit(1)
	}
}

func GetMigrations(connection *sqlx.DB, query string) []types.Migration {
	m := []types.Migration{}
	if err := connection.Select(&m, query); err != nil {
		utils.Error("Getting migration log results", err.Error())
		os.Exit(1)
	}
	return m
}

func WriteMigrationLog(connection *sqlx.DB, query, file, hash string) {
	if _, err := connection.Exec(query, file, hash); err != nil {
		utils.Error("Insert into migration table", err.Error())
		os.Exit(1)
	}
}

func CleanMigrationLog(connection *sqlx.DB, query, file string) {
	if _, err := connection.Exec(query, file); err != nil {
		utils.Error("Remove from migration table", err.Error())
		os.Exit(1)
	}
}

func ApplyMigration(connection *sqlx.DB, query string) {
	tx, err := connection.Begin()
	if err != nil {
		utils.Error("Applying migration - Begin", err.Error())
		os.Exit(1)
	}

	if _, err := connection.Exec(query); err != nil {
		if err := tx.Rollback(); err != nil {
			utils.Error("Applying migration - Rollback", err.Error())
			os.Exit(1)
		}
		utils.Error("Applying migration", err.Error())
		os.Exit(1)
	}

	if err := tx.Commit(); err != nil {
		utils.Error("Applying migration - Commit", err.Error())
		os.Exit(1)
	}
}
