package postgres

import (
	"fmt"
	"migoro/query"
	"migoro/types"
	"migoro/utils"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	SQL_HOST         = "SQL_HOST"
	SQL_PORT         = "SQL_PORT"
	SQL_USER         = "SQL_USER"
	SQL_PASSWORD     = "SQL_PASSWORD"
	SQL_DB           = "SQL_DB"
	SQL_DB_SCHEMA    = "SQL_DB_SCHEMA"
	SQL_SSL          = "SQL_SSL"
	MIGRATION_DIR    = "MIGRATION_DIR"
	MIGRATION_TABLE  = "MIGRATION_TABLE"
	MIGRATION_SCHEMA = "MIGRATION_SCHEMA"
)

type Postgres struct{}

func (adapter *Postgres) Connection() *sqlx.DB {
	connection, err := sqlx.Open(utils.Env("SQL_DRIVER"), fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		utils.Env(SQL_HOST),
		utils.Env(SQL_PORT),
		utils.Env(SQL_USER),
		utils.Env(SQL_PASSWORD),
		utils.Env(SQL_DB),
		utils.Env(SQL_SSL),
		utils.Env(SQL_DB_SCHEMA),
	))
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

func (adapter *Postgres) ConnectionWithoutDB() *sqlx.DB {
	connection, err := sqlx.Open(utils.Env("SQL_DRIVER"), fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s",
		utils.Env(SQL_HOST),
		utils.Env(SQL_PORT),
		utils.Env(SQL_USER),
		utils.Env(SQL_PASSWORD),
		utils.Env(SQL_SSL),
	))
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

func (adapter *Postgres) ValidateEnvironment() {
	utils.ValidateEnvVariables([]string{
		SQL_HOST,
		SQL_PORT,
		SQL_USER,
		SQL_PASSWORD,
		SQL_DB,
		SQL_SSL,
		MIGRATION_DIR,
		MIGRATION_TABLE,
	})
}

func (adapter *Postgres) GetMigrationTableName() string {
	return utils.Env(MIGRATION_TABLE)
}

func (adapter *Postgres) GetDatabaseName() string {
	return utils.Env(SQL_DB)
}

func (adapter *Postgres) DatabaseExists() types.DbCheck {
	return query.Exists(adapter.ConnectionWithoutDB(), DatabaseExistsQuery())
}

func (adapter *Postgres) CreateDatabase() {
	query.Query(adapter.ConnectionWithoutDB(), CreateDatabaseQuery())
}

func (adapter *Postgres) MigrationsLogExists() types.DbCheck {
	return query.Exists(adapter.Connection(), TableLogExistsQuery())
}

func (adapter *Postgres) CreateMigrationsLog() {
	query.Query(adapter.Connection(), CreateLogTableQuery())
}

func (adapter *Postgres) GetMigrationsFromLog() []types.Migration {
	return query.GetMigrations(adapter.Connection(), GetMigrationsQuery())
}

func (adapter *Postgres) WriteMigrationLog(file string, hash string) {
	query.WriteMigrationLog(adapter.Connection(), WriteMigrationLogQuery(), file, hash)
}

func (adapter *Postgres) GetLatestMigrationsFromLog() []types.Migration {
	return query.GetMigrations(adapter.Connection(), GetLatestMigrationsQuery())
}

func (adapter *Postgres) CleanMigrationLog(file string) {
	query.CleanMigrationLog(adapter.Connection(), CleanMigrationLogQuery(), file)
}
