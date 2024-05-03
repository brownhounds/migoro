package adapters

import (
	"fmt"
	"migoro/utils"
	"os"

	"github.com/jmoiron/sqlx"
)

const (
	SQL_DRIVER = "SQL_DRIVER"
	MIGRATION_DIR = "MIGRATION_DIR"
	MIGRATION_TABLE = "MIGRATION_TABLE"
)

const (
	POSTGRES = "postgres"
	SQLITE = "sqlite"
)

type Adapter interface {
	Connection() *sqlx.DB
	ValidateEnvironment()
	TableLogExistsQuery() string
	CreateLogTableQuery() string
	GetMigrationsQuery() string
	WriteMigrationLogQuery() string
	GetLatestMigrationsQuery() string
	RollbackMigrationLogQuery() string
}


func Init() Adapter {
	adapter, err := resolveAdapter();
	if err != nil {
		utils.Error("Resolving Adapter", err.Error())
		os.Exit(1)
	}

	adapter.ValidateEnvironment()

	return adapter
}

func ValidateEnvVariables(envVars []string) {
    for _, value := range envVars {
		_, defined := os.LookupEnv(value)
		if !defined {
			panic("ENV Variable is not defined: " + value);
		}
    }
}

func resolveAdapter() (Adapter, error) {
	switch utils.Env(SQL_DRIVER) {
	case POSTGRES:
		return Postgres{}, nil
	case SQLITE:
		return nil, fmt.Errorf("could not resolve the adapter: %s", utils.Env(SQL_DRIVER))
	default:
		return nil, fmt.Errorf("could not resolve the adapter: %s", utils.Env(SQL_DRIVER))
	}
}
