package postgres

import (
	"fmt"

	"github.com/brownhounds/migoro/utils"
)

func DatabaseExistsQuery() string {
	return fmt.Sprintf(`SELECT EXISTS (SELECT datname FROM pg_database WHERE datname = '%s') as test;`, utils.Env(SQL_DB))
}

func CreateDatabaseQuery() string {
	return fmt.Sprintf(`CREATE DATABASE %s;`, utils.Env(SQL_DB))
}

func TableLogExistsQuery() string {
	return fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = '%s' AND table_name = '%s') as test;", utils.Env(MIGRATION_SCHEMA), utils.Env(MIGRATION_TABLE))
}

func CreateLogTableQuery() string {
	return fmt.Sprintf(`
		CREATE SCHEMA IF NOT EXISTS %s;
		CREATE SCHEMA IF NOT EXISTS %s;
		CREATE TABLE IF NOT EXISTS %s.%s (
			id serial PRIMARY KEY NOT NULL,
			migration_file character varying(512) NOT NULL,
			migration_hash TEXT NOT NULL,
			created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
	);`, utils.Env("MIGRATION_SCHEMA"), utils.Env("SQL_DB_SCHEMA"), utils.Env(MIGRATION_SCHEMA), utils.Env(MIGRATION_TABLE))
}

func GetMigrationsQuery() string {
	return fmt.Sprintf("SELECT migration_file FROM %s.%s", utils.Env(MIGRATION_SCHEMA), utils.Env(MIGRATION_TABLE))
}

func WriteMigrationLogQuery() string {
	return fmt.Sprintf("INSERT INTO %s.%s (migration_file, migration_hash) VALUES ($1, $2);", utils.Env(MIGRATION_SCHEMA), utils.Env(MIGRATION_TABLE))
}

func GetLatestMigrationsQuery() string {
	return fmt.Sprintf("SELECT migration_file FROM %s.%s WHERE migration_hash = (SELECT migration_hash FROM %s.%s ORDER BY created_at DESC LIMIT 1) ORDER BY migration_file DESC", utils.Env(MIGRATION_SCHEMA), utils.Env(MIGRATION_TABLE), utils.Env(MIGRATION_SCHEMA), utils.Env(MIGRATION_TABLE))
}

func CleanMigrationLogQuery() string {
	return fmt.Sprintf("DELETE FROM %s.%s WHERE migration_file = $1", utils.Env(MIGRATION_SCHEMA), utils.Env(MIGRATION_TABLE))
}
