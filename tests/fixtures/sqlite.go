package test_fixtures

import (
	"migoro/adapters"
	"migoro/adapters/sqlite"
	"os"
	"testing"
)

type SqliteEnv struct {
	Key   string
	Value string
}

type SqliteFixture struct {
	TEST_DATABASE string
	ENV           []struct {
		Key   string
		Value string
	}
}

func (s *SqliteFixture) New() *SqliteFixture {
	s.TEST_DATABASE = "test_database.db"
	s.ENV = []struct {
		Key   string
		Value string
	}{
		{Key: adapters.SQL_DRIVER, Value: adapters.SQLITE3},
		{Key: sqlite.SQL_FILE, Value: s.TEST_DATABASE},
		{Key: sqlite.MIGRATION_DIR, Value: "migrations"},
		{Key: sqlite.MIGRATION_TABLE, Value: "migration_log"},
	}

	return s
}

func (s *SqliteFixture) InitEnv(t *testing.T) {
	for _, item := range s.ENV {
		t.Setenv(item.Key, item.Value)
	}
}

func (s *SqliteFixture) RemoveDatabase(t *testing.T) {
	if err := os.Remove(s.TEST_DATABASE); err != nil {
		t.Fatal(err)
	}
}
