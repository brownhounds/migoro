package migoro_test

import (
	"fmt"
	"migoro/adapters"
	"migoro/dispatcher"
	"testing"

	gt "github.com/brownhounds/go-testing"
)

func TestShouldPanicWithoutSQLDriverSet(t *testing.T) {
	const UNSUPPORTED_ADAPTER = "unsupported_driver"

	testCases := []struct {
		name     string
		value    string
		expected string
	}{
		{
			name:     fmt.Sprintf("%s is set to empty string", adapters.SQL_DRIVER),
			value:    "",
			expected: adapters.UnsetAdapterErrorMessage(adapters.SQL_DRIVER),
		},
		{
			name:     fmt.Sprintf("%s is set to unsupported driver", adapters.SQL_DRIVER),
			value:    "unsupported_driver",
			expected: adapters.UnsupportedAdapterErrorMessage(UNSUPPORTED_ADAPTER),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Setenv(adapters.SQL_DRIVER, testCase.value)
			gt.AssertPanic(t, testCase.expected, dispatcher.Init)
		})
	}
}
