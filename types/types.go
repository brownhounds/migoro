package types

import (
	"github.com/jmoiron/sqlx"
)

type Migration struct {
	MigrationFile string `db:"migration_file"`
}

type DbCheck struct {
	Exists bool `db:"test"`
}

type Adapter interface {
	ValidateEnvironment()
	CreateDatabase()
	CreateMigrationsLog()
	Connection() (error, *sqlx.DB)
	GetMigrationTableName() string
	GetDatabaseName() string
	WriteMigrationLog(file string, hash string)
	CleanMigrationLog(file string)
	DatabaseExists() (error, *DbCheck)
	MigrationsLogExists() (error, *DbCheck)
	GetMigrationsFromLog() (error, *[]Migration)
	GetLatestMigrationsFromLog() (error, *[]Migration)
}
