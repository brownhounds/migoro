package utils

import (
	"fmt"

	"github.com/logrusorgru/aurora/v4"
)

func Error(h string, m string) {
	hf := aurora.Red(h + ":").Bold()
	mf := aurora.Red(m)
	fmt.Printf("%s %s\n", hf, mf)
}

func Warning(h string, m string) {
	hf := aurora.Yellow(h + ":").Bold()
	mf := aurora.Yellow(m)
	fmt.Printf("%s %s\n", hf, mf)
}

func Success(h string, m string) {
	hf := aurora.Green(h + ":").Bold()
	mf := aurora.Green(m)
	fmt.Printf("%s %s\n", hf, mf)
}
