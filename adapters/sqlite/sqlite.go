package sqlite

import (
	"migoro/error_context"
	"migoro/query"
	"migoro/types"
	"migoro/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	SQL_FILE        = "SQL_FILE"
	MIGRATION_DIR   = "MIGRATION_DIR"
	MIGRATION_TABLE = "MIGRATION_TABLE"
)

type Sqlite struct{}

func (adapter Sqlite) Connection() (error, *sqlx.DB) {
	connection, err := sqlx.Open(utils.Env("SQL_DRIVER"), utils.Env("SQL_FILE"))
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

func (adapter Sqlite) DatabaseExists() (error, *types.DbCheck) {
	return nil, &types.DbCheck{Exists: true}
}

func (adapter Sqlite) CreateDatabase() {}

func (adapter Sqlite) MigrationsLogExists() (error, *types.DbCheck) {
	err, con := adapter.Connection()
	defer con.Close()
	if err != nil {
		error_context.Context.SetError()
		return err, nil
	}
	return nil, query.Exists(con, TableLogExistsQuery())
}

func (adapter Sqlite) CreateMigrationsLog() {
	err, con := adapter.Connection()
	defer con.Close()
	if err != nil {
		error_context.Context.SetError()
		return
	}
	query.Query(con, CreateLogTableQuery())
}

func (adapter Sqlite) GetMigrationsFromLog() (error, *[]types.Migration) {
	err, con := adapter.Connection()
	defer con.Close()
	if err != nil {
		error_context.Context.SetError()
		return err, nil
	}
	return nil, query.GetMigrations(con, GetMigrationsQuery())
}

func (adapter Sqlite) WriteMigrationLog(file, hash string) {
	err, con := adapter.Connection()
	defer con.Close()
	if err != nil {
		error_context.Context.SetError()
		return
	}
	query.WriteMigrationLog(con, WriteMigrationLogQuery(), file, hash)
}

func (adapter Sqlite) GetLatestMigrationsFromLog() (error, *[]types.Migration) {
	err, con := adapter.Connection()
	defer con.Close()
	if err != nil {
		error_context.Context.SetError()
		return err, nil
	}
	return nil, query.GetMigrations(con, GetLatestMigrationsQuery())
}

func (adapter Sqlite) CleanMigrationLog(file string) {
	err, con := adapter.Connection()
	defer con.Close()
	if err != nil {
		error_context.Context.SetError()
		return
	}
	query.CleanMigrationLog(con, CleanMigrationLogQuery(), file)
}
