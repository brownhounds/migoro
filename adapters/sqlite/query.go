package sqlite

import (
	"fmt"
	"migoro/utils"
)

type Query struct {}

func (q Query) TableLogExistsQuery() string {
	return fmt.Sprintf("SELECT EXISTS (SELECT name FROM sqlite_master WHERE type='table' AND name='%s') as test;", utils.Env("MIGRATION_TABLE"))
}

func (q Query) CreateLogTableQuery() string {
	return fmt.Sprintf(`
	CREATE TABLE %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		migration_file TEXT NOT NULL,
		migration_hash TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`, utils.Env("MIGRATION_TABLE"))
}

func (q Query) GetMigrationsQuery() string {
	return fmt.Sprintf("SELECT migration_file FROM %s", utils.Env("MIGRATION_TABLE"))
}

func (q Query) WriteMigrationLogQuery() string {
	return fmt.Sprintf("INSERT INTO %s (migration_file, migration_hash) VALUES ($1, $2);", utils.Env("MIGRATION_TABLE"))
}

func (q Query) GetLatestMigrationsQuery() string {
	return fmt.Sprintf(`
	SELECT migration_file
	FROM %s
	WHERE migration_hash = (
		SELECT migration_hash
		FROM %s
		ORDER BY created_at DESC
		LIMIT 1
	)
	ORDER BY migration_file DESC;`, utils.Env("MIGRATION_TABLE"), utils.Env("MIGRATION_TABLE"))
}

func (q Query) CleanMigrationLogQuery() string {
	return fmt.Sprintf("DELETE FROM %s WHERE migration_file = $1", utils.Env("MIGRATION_TABLE"))
}
