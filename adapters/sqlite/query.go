package sqlite

import (
	"fmt"
	"migoro/utils"
)

func TableLogExistsQuery() string {
	return fmt.Sprintf("SELECT EXISTS (SELECT name FROM sqlite_master WHERE type='table' AND name='%s') as test;", utils.Env(MIGRATION_TABLE))
}

func CreateLogTableQuery() string {
	return fmt.Sprintf(`
	CREATE TABLE %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		migration_file TEXT NOT NULL,
		migration_hash TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`, utils.Env(MIGRATION_TABLE))
}

func GetMigrationsQuery() string {
	return fmt.Sprintf("SELECT migration_file FROM %s", utils.Env(MIGRATION_TABLE))
}

func WriteMigrationLogQuery() string {
	return fmt.Sprintf("INSERT INTO %s (migration_file, migration_hash) VALUES ($1, $2);", utils.Env(MIGRATION_TABLE))
}

func GetLatestMigrationsQuery() string {
	return fmt.Sprintf(`
	SELECT migration_file
	FROM %s
	WHERE migration_hash = (
		SELECT migration_hash
		FROM %s
		ORDER BY created_at DESC
		LIMIT 1
	)
	ORDER BY migration_file DESC;`, utils.Env(MIGRATION_TABLE), utils.Env(MIGRATION_TABLE))
}

func CleanMigrationLogQuery() string {
	return fmt.Sprintf("DELETE FROM %s WHERE migration_file = $1", utils.Env(MIGRATION_TABLE))
}
