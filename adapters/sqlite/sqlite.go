package sqlite

import (
	"os"

	"github.com/brownhounds/migoro/query"
	"github.com/brownhounds/migoro/types"
	"github.com/brownhounds/migoro/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	SQL_FILE        = "SQL_FILE"
	MIGRATION_DIR   = "MIGRATION_DIR"
	MIGRATION_TABLE = "MIGRATION_TABLE"
)

type Sqlite struct{}

func (adapter Sqlite) Connection() *sqlx.DB {
	connection, err := sqlx.Open(utils.Env("SQL_DRIVER"), utils.Env("SQL_FILE"))
	if err != nil {
		utils.Error("Database connection", err.Error())
		os.Exit(1)
	}

	{
		// Driver doesn't log error on initial connection
		// Ping is necessary to evaluate early
		err := connection.Ping()
		if err != nil {
			connection.Close()
			utils.Error("Database connection", err.Error())
			os.Exit(1)
		}
	}

	return connection
}

func (adapter Sqlite) ValidateEnvironment() {
	utils.ValidateEnvVariables([]string{
		SQL_FILE,
		MIGRATION_DIR,
		MIGRATION_TABLE,
	})
}

func (adapter Sqlite) GetMigrationTableName() string {
	return utils.Env(MIGRATION_TABLE)
}

func (adapter Sqlite) GetDatabaseName() string {
	return utils.Env(SQL_FILE)
}

func (adapter Sqlite) DatabaseExists() types.DbCheck {
	return types.DbCheck{Exists: true}
}

func (adapter Sqlite) CreateDatabase() {}

func (adapter Sqlite) MigrationsLogExists() types.DbCheck {
	return query.Exists(adapter.Connection(), TableLogExistsQuery())
}

func (adapter Sqlite) CreateMigrationsLog() {
	query.Query(adapter.Connection(), CreateLogTableQuery())
}

func (adapter Sqlite) GetMigrationsFromLog() []types.Migration {
	return query.GetMigrations(adapter.Connection(), GetMigrationsQuery())
}

func (adapter Sqlite) WriteMigrationLog(file, hash string) {
	query.WriteMigrationLog(adapter.Connection(), WriteMigrationLogQuery(), file, hash)
}

func (adapter Sqlite) GetLatestMigrationsFromLog() []types.Migration {
	return query.GetMigrations(adapter.Connection(), GetLatestMigrationsQuery())
}

func (adapter Sqlite) CleanMigrationLog(file string) {
	query.CleanMigrationLog(adapter.Connection(), CleanMigrationLogQuery(), file)
}
