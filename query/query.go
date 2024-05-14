package query

import (
	"migoro/error_context"
	"migoro/types"
	"migoro/utils"

	"github.com/jmoiron/sqlx"
)

func Exists(connection *sqlx.DB, query string) *types.DbCheck {
	r := types.DbCheck{}
	if err := connection.Get(&r, query); err != nil {
		error_context.Context.SetError()
		utils.Error("Executing query", err.Error())
		return nil
	}
	return &r
}

func Query(connection *sqlx.DB, query string) {
	if _, err := connection.Exec(query); err != nil {
		error_context.Context.SetError()
		utils.Error("Executing query", err.Error())
	}
}

func GetMigrations(connection *sqlx.DB, query string) *[]types.Migration {
	m := []types.Migration{}
	if err := connection.Select(&m, query); err != nil {
		error_context.Context.SetError()
		utils.Error("Getting migration log results", err.Error())
		return nil
	}
	return &m
}

func WriteMigrationLog(connection *sqlx.DB, query, file, hash string) {
	if _, err := connection.Exec(query, file, hash); err != nil {
		error_context.Context.SetError()
		utils.Error("Insert into migration table", err.Error())
	}
}

func CleanMigrationLog(connection *sqlx.DB, query, file string) {
	if _, err := connection.Exec(query, file); err != nil {
		error_context.Context.SetError()
		utils.Error("Remove from migration table", err.Error())
	}
}

func ApplyMigration(connection *sqlx.DB, query string) {
	tx, err := connection.Begin()
	if err != nil {
		error_context.Context.SetError()
		utils.Error("Applying migration - Begin", err.Error())
		return
	}

	if _, err := connection.Exec(query); err != nil {
		if err := tx.Rollback(); err != nil {
			error_context.Context.SetError()
			utils.Error("Applying migration - Rollback", err.Error())
			return
		}
		error_context.Context.SetError()
		utils.Error("Applying migration", err.Error())
		return
	}

	if err := tx.Commit(); err != nil {
		error_context.Context.SetError()
		utils.Error("Applying migration - Commit", err.Error())
		return
	}
}
