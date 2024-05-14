package adapters

import (
	"errors"
	"fmt"
	"migoro/adapters/postgres"
	"migoro/adapters/sqlite"
	"migoro/error_context"
	"migoro/types"
	"migoro/utils"
)

const (
	SQL_DRIVER = "SQL_DRIVER"
	POSTGRES   = "postgres"
	SQLITE3    = "sqlite3"
)

var unresolvedAdapterErrorMessage = errors.New("could not resolve the adapter")

func unresolvedAdapterError(adapter string) error {
	return fmt.Errorf("%w: %s", unresolvedAdapterErrorMessage, adapter)
}

func Init() (error, types.Adapter) {
	adapter, err := resolveAdapter()
	if err != nil {
		error_context.Context.SetError()
		utils.Error("Resolving Adapter", err.Error())
		return err, nil
	}

	adapter.ValidateEnvironment()

	return nil, adapter
}

func resolveAdapter() (types.Adapter, error) {
	switch utils.Env(SQL_DRIVER) {
	case POSTGRES:
		return &postgres.Postgres{}, nil
	case SQLITE3:
		return &sqlite.Sqlite{}, nil
	default:
		var adapter string
		if utils.Env(SQL_DRIVER) == "" {
			adapter = "ADAPTER NOT SET"
		} else {
			adapter = utils.Env(SQL_DRIVER)
		}

		return nil, unresolvedAdapterError(adapter)
	}
}
