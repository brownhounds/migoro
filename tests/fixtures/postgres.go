package test_fixtures

import (
	"fmt"
	"migoro/adapters"
	"migoro/adapters/postgres"
	"testing"
)

type PostgresEnv struct {
	Key   string
	Value string
}

type PostgresFixture struct {
	TEST_DATABASE string
	MIGRATION_DIR string
	ENV           []struct {
		Key   string
		Value string
	}
}

func (p *PostgresFixture) New() *PostgresFixture {
	p.TEST_DATABASE = "test_database"
	p.MIGRATION_DIR = "migrations"
	p.ENV = []struct {
		Key   string
		Value string
	}{
		{Key: adapters.SQL_DRIVER, Value: adapters.POSTGRES},
		{Key: postgres.SQL_HOST, Value: "localhost"},
		{Key: postgres.SQL_PORT, Value: "5432"},
		{Key: postgres.SQL_USER, Value: "admin"},
		{Key: postgres.SQL_PASSWORD, Value: "admin"},
		{Key: postgres.SQL_DB, Value: p.TEST_DATABASE},
		{Key: postgres.SQL_DB_SCHEMA, Value: "test_schema"},
		{Key: postgres.SQL_SSL, Value: "disable"},
		{Key: postgres.MIGRATION_DIR, Value: p.MIGRATION_DIR},
		{Key: postgres.MIGRATION_TABLE, Value: "migration_log"},
		{Key: postgres.MIGRATION_SCHEMA, Value: "test_platform"},
	}

	return p
}

func (p *PostgresFixture) InitEnv(t *testing.T) {
	for _, item := range p.ENV {
		t.Setenv(item.Key, item.Value)
	}
}

func (p *PostgresFixture) RemoveDatabase(t *testing.T) {
	database := postgres.Postgres{}
	err, con := database.ConnectionWithoutDB()
	if err != nil {
		t.Fatal(err)
	}
	defer con.Close()

	query := fmt.Sprintf(`DROP DATABASE %s;`, p.TEST_DATABASE)

	if _, err := con.Exec(query); err != nil {
		t.Fatal(err)
	}
}
