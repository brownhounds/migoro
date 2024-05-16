package postgres_test

import (
	"bytes"
	"log"
	"migoro/dispatcher"
	test_fixtures "migoro/tests/fixtures"
	"migoro/utils"
	"testing"

	gt "github.com/brownhounds/go-testing"
)

func TestSqliteValidateEnvVariables(t *testing.T) {
	fixture := &test_fixtures.SqliteFixture{}
	fixture.New()

	for index, item := range fixture.ENV {
		t.Run(item.Key, func(t *testing.T) {
			for _, value := range fixture.ENV[:index+1] {
				t.Setenv(value.Key, value.Value)
			}

			if index != len(fixture.ENV)-1 {
				gt.AssertPanic(t, utils.EnvVarNotDefinedErrorMessage(
					fixture.ENV[index+1].Key),
					dispatcher.Init,
				)
			}
		})
	}
}

func TestSqliteInitCommand(t *testing.T) {
	fixture := &test_fixtures.SqliteFixture{}
	fixture.New()
	fixture.InitEnv(t)

	t.Cleanup(func() {
		fixture.InitEnv(t)
		fixture.RemoveDatabase(t)
	})

	testCases := []struct {
		name string
	}{
		{name: "InitialRun"},
		{name: "FollowingRun"},
	}

	for _, item := range testCases {
		t.Run(item.name, func(t *testing.T) {
			buffer := new(bytes.Buffer)
			log.SetFlags(0)
			log.SetOutput(buffer)

			dispatcher.Init()
			gt.ToMatchSnapshot(t, buffer.String(), item.name)
		})
	}
}
