package postgres

import (
	"fmt"
	"migoro/error_context"
	"migoro/query"
	"migoro/types"
	"migoro/utils"

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

func (adapter Postgres) Connection() (error, *sqlx.DB) {
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
		error_context.Context.SetError()
		return err, nil
	}

	{
		// Driver doesn't log error on initial connection
		// Ping is necessary to evaluate early
		err := connection.Ping()
		if err != nil {
			connection.Close()
			utils.Error("Database connection", err.Error())
			error_context.Context.SetError()
			return err, nil
		}
	}
	return nil, connection
}

func (adapter Postgres) ConnectionWithoutDB() (error, *sqlx.DB) {
	connection, err := sqlx.Open(utils.Env("SQL_DRIVER"), fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s",
		utils.Env(SQL_HOST),
		utils.Env(SQL_PORT),
		utils.Env(SQL_USER),
		utils.Env(SQL_PASSWORD),
		utils.Env(SQL_SSL),
	))
	if err != nil {
		utils.Error("Database connection", err.Error())
		error_context.Context.SetError()
		return err, nil
	}

	{
		// Driver doesn't log error on initial connection
		// Ping is necessary to evaluate early
		err := connection.Ping()
		if err != nil {
			connection.Close()
			utils.Error("Database connection", err.Error())
			error_context.Context.SetError()
			return err, nil
		}
	}
	return nil, connection
}

func (adapter Postgres) ValidateEnvironment() {
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

func (adapter Postgres) GetMigrationTableName() string {
	return utils.Env(MIGRATION_TABLE)
}

func (adapter Postgres) GetDatabaseName() string {
	return utils.Env(SQL_DB)
}

func (adapter Postgres) DatabaseExists() (error, *types.DbCheck) {
	err, con := adapter.ConnectionWithoutDB()
	if err != nil {
		error_context.Context.SetError()
		return err, nil
	}
	return nil, query.Exists(con, DatabaseExistsQuery())
}

func (adapter Postgres) CreateDatabase() {
	err, con := adapter.ConnectionWithoutDB()
	if err != nil {
		error_context.Context.SetError()
		return
	}
	query.Query(con, CreateDatabaseQuery())
}

func (adapter Postgres) MigrationsLogExists() (error, *types.DbCheck) {
	err, con := adapter.Connection()
	if err != nil {
		error_context.Context.SetError()
		return err, nil
	}
	return nil, query.Exists(con, TableLogExistsQuery())
}

func (adapter Postgres) CreateMigrationsLog() {
	err, con := adapter.Connection()
	if err != nil {
		error_context.Context.SetError()
		return
	}
	query.Query(con, CreateLogTableQuery())
}

func (adapter Postgres) GetMigrationsFromLog() (error, *[]types.Migration) {
	err, con := adapter.Connection()
	if err != nil {
		error_context.Context.SetError()
		return err, nil
	}
	return nil, query.GetMigrations(con, GetMigrationsQuery())
}

func (adapter Postgres) WriteMigrationLog(file, hash string) {
	err, con := adapter.Connection()
	if err != nil {
		error_context.Context.SetError()
		return
	}
	query.WriteMigrationLog(con, WriteMigrationLogQuery(), file, hash)
}

func (adapter Postgres) GetLatestMigrationsFromLog() (error, *[]types.Migration) {
	err, con := adapter.Connection()
	if err != nil {
		error_context.Context.SetError()
		return err, nil
	}
	return nil, query.GetMigrations(con, GetLatestMigrationsQuery())
}

func (adapter Postgres) CleanMigrationLog(file string) {
	err, con := adapter.Connection()
	if err != nil {
		error_context.Context.SetError()
		return
	}
	query.CleanMigrationLog(con, CleanMigrationLogQuery(), file)
}
