package main

import (
	"migoro/cmd"
	"migoro/error_context"
)

func main() {
	cmd.Execute()
	error_context.Context.ExitWithError()
}
