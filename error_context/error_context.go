package error_context

import (
	"fmt"
	"os"
)

var Context = ErrorContext{hasError: false}

type ErrorContext struct {
	hasError bool
}

func (e *ErrorContext) SetError() {
	e.hasError = true
}

func (e *ErrorContext) Exit() {
	if e.hasError {
		fmt.Println("!!!Exit triggered by context")
		os.Exit(1)
	}
}
