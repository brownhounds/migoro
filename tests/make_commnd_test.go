package migoro_test

import (
	"bytes"
	"fmt"
	"log"
	"migoro/dispatcher"
	test_fixtures "migoro/tests/fixtures"
	test_utils "migoro/tests/utils"
	"path/filepath"
	"testing"

	gt "github.com/brownhounds/go-testing"
)

func TestMakeCommandCreatesAMigrationFiles(t *testing.T) {
	fixture := test_fixtures.SqliteFixture{}
	fixture.New()
	fixture.InitEnv(t)

	t.Cleanup(func() {
		test_utils.RestoreUtilsTimestamp()
		gt.RemoveFileDir(t, fixture.MIGRATION_DIR)
	})

	testCases := []struct {
		filename  string
		timestamp int64
	}{
		{
			filename:  "create_test_table",
			timestamp: int64(1),
		},
		{
			filename:  "create_test_table_2",
			timestamp: int64(2),
		},
	}

	for index, item := range testCases {
		t.Run(item.filename, func(t *testing.T) {
			buffer := new(bytes.Buffer)
			log.SetFlags(0)
			log.SetOutput(buffer)
			test_utils.MockUtilsTimestamp(int64(index + 1))

			dispatcher.Make(item.filename)

			gt.AssertFileDirExists(t, filepath.Join(fixture.MIGRATION_DIR, fmt.Sprintf("%d_%s_%s.sql", index+1, item.filename, "up")))
			gt.AssertFileDirExists(t, filepath.Join(fixture.MIGRATION_DIR, fmt.Sprintf("%d_%s_%s.sql", index+1, item.filename, "down")))

			gt.ToMatchSnapshot(t, buffer.String(), item.filename)
		})
	}
}
