package adapters

import (
	"errors"
	"fmt"
	"os"

	"github.com/brownhounds/migoro/adapters/postgres"
	"github.com/brownhounds/migoro/adapters/sqlite"
	"github.com/brownhounds/migoro/types"
	"github.com/brownhounds/migoro/utils"
)

const (
	SQL_DRIVER = "SQL_DRIVER"
	POSTGRES   = "postgres"
	SQLITE3    = "sqlite3"
)

var unresolvedAdapterErrorMessage = errors.New("could not resolve the adapter")

func unresolvedAdapterError(adapter string) error {
	return fmt.Errorf("%w : %s", unresolvedAdapterErrorMessage, adapter)
}

func Init() types.Adapter {
	adapter, err := resolveAdapter()
	if err != nil {
		utils.Error("Resolving Adapter", err.Error())
		os.Exit(1)
	}

	adapter.ValidateEnvironment()

	return adapter
}

func resolveAdapter() (types.Adapter, error) {
	switch utils.Env(SQL_DRIVER) {
	case POSTGRES:
		return &postgres.Postgres{}, nil
	case SQLITE3:
		return &sqlite.Sqlite{}, nil
	default:
		return nil, unresolvedAdapterError(utils.Env(SQL_DRIVER))
	}
}
