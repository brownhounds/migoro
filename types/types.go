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
	Connection() *sqlx.DB
	ValidateEnvironment()
	GetMigrationTableName() string
	GetDatabaseName() string
	DatabaseExists() DbCheck
	CreateDatabase()
	MigrationsLogExists() DbCheck
	CreateMigrationsLog()
	GetMigrationsFromLog() []Migration
	WriteMigrationLog(file string, hash string)
	GetLatestMigrationsFromLog() []Migration
	CleanMigrationLog(file string)
}
