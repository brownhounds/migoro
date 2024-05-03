package adapters

import (
	"fmt"
	"os"

	"migoro/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	SQL_HOST = "SQL_HOST"
	SQL_PORT = "SQL_PORT"
	SQL_USER = "SQL_USER"
	SQL_PASSWORD = "SQL_PASSWORD"
	SQL_DB = "SQL_DB"
	SQL_SSL = "SQL_SSL"
)

type Postgres struct {}

func (a Postgres) Connection() *sqlx.DB {
	connection, err := sqlx.Open(utils.Env("SQL_DRIVER"), fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	utils.Env(SQL_HOST),
	utils.Env(SQL_PORT),
	utils.Env(SQL_USER),
	utils.Env(SQL_PASSWORD),
	utils.Env(SQL_DB),
	utils.Env(SQL_SSL)))

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

func (a Postgres) ValidateEnvironment() {
    ValidateEnvVariables([]string{
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

func (a Postgres) TableLogExistsQuery() string {
	return fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = '%s' AND table_name = '%s') as test;", utils.Env("MIGRATION_SCHEMA"), utils.Env("MIGRATION_TABLE"))
}

func (a Postgres) CreateLogTableQuery() string {
	return fmt.Sprintf(`CREATE SCHEMA %s;CREATE TABLE %s.%s (
		id serial PRIMARY KEY NOT NULL,
		migration_file character varying(512) NOT NULL,
		migration_hash TEXT NOT NULL,
		created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
	);`, utils.Env("MIGRATION_SCHEMA"), utils.Env("MIGRATION_SCHEMA"), utils.Env("MIGRATION_TABLE"))
}

func (a Postgres) GetMigrationsQuery() string {
	return fmt.Sprintf("SELECT migration_file FROM %s.%s", utils.Env("MIGRATION_SCHEMA"), utils.Env("MIGRATION_TABLE"))
}

func (a Postgres) WriteMigrationLogQuery() string {
	return fmt.Sprintf("INSERT INTO %s.%s (migration_file, migration_hash) VALUES ($1, $2);", utils.Env("MIGRATION_SCHEMA"), utils.Env("MIGRATION_TABLE"))
}

func (a Postgres) GetLatestMigrationsQuery() string {
	return fmt.Sprintf("SELECT migration_file FROM %s.%s WHERE migration_hash = (SELECT migration_hash FROM %s.%s ORDER BY created_at DESC LIMIT 1) ORDER BY migration_file DESC", utils.Env("MIGRATION_SCHEMA"), utils.Env("MIGRATION_TABLE"), utils.Env("MIGRATION_SCHEMA"), utils.Env("MIGRATION_TABLE"))
}

func (a Postgres) RollbackMigrationLogQuery() string {
	return fmt.Sprintf("DELETE FROM %s.%s WHERE migration_file = $1", utils.Env("MIGRATION_SCHEMA"), utils.Env("MIGRATION_TABLE"))
}
