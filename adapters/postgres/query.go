package postgres

import (
	"fmt"
	"migoro/utils"
)

type Query struct {}

func (q Query) DatabaseExistsQuery() string {
	return fmt.Sprintf(`SELECT EXISTS (SELECT datname FROM pg_database WHERE datname = '%s') as test;`, utils.Env("SQL_DB"))
}

func (q Query) CreateDatabaseQuery() string {
	return fmt.Sprintf(`CREATE DATABASE %s;`, utils.Env("SQL_DB"))
}

func (q Query) TableLogExistsQuery() string {
	return fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = '%s' AND table_name = '%s') as test;", utils.Env("MIGRATION_SCHEMA"), utils.Env("MIGRATION_TABLE"))
}

func (q Query) CreateLogTableQuery() string {
	return fmt.Sprintf(`
		CREATE SCHEMA %s;
		CREATE SCHEMA %s;
		CREATE TABLE %s.%s (
			id serial PRIMARY KEY NOT NULL,
			migration_file character varying(512) NOT NULL,
			migration_hash TEXT NOT NULL,
			created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
	);`, utils.Env("MIGRATION_SCHEMA"), utils.Env("SQL_DB_SCHEMA"), utils.Env("MIGRATION_SCHEMA"), utils.Env("MIGRATION_TABLE"))
}

func (q Query) GetMigrationsQuery() string {
	return fmt.Sprintf("SELECT migration_file FROM %s.%s", utils.Env("MIGRATION_SCHEMA"), utils.Env("MIGRATION_TABLE"))
}

func (q Query) WriteMigrationLogQuery() string {
	return fmt.Sprintf("INSERT INTO %s.%s (migration_file, migration_hash) VALUES ($1, $2);", utils.Env("MIGRATION_SCHEMA"), utils.Env("MIGRATION_TABLE"))
}

func (q Query) GetLatestMigrationsQuery() string {
	return fmt.Sprintf("SELECT migration_file FROM %s.%s WHERE migration_hash = (SELECT migration_hash FROM %s.%s ORDER BY created_at DESC LIMIT 1) ORDER BY migration_file DESC", utils.Env("MIGRATION_SCHEMA"), utils.Env("MIGRATION_TABLE"), utils.Env("MIGRATION_SCHEMA"), utils.Env("MIGRATION_TABLE"))
}

func (q Query) CleanMigrationLogQuery() string {
	return fmt.Sprintf("DELETE FROM %s.%s WHERE migration_file = $1", utils.Env("MIGRATION_SCHEMA"), utils.Env("MIGRATION_TABLE"))
}
