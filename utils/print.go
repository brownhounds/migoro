package utils

import (
	"fmt"

	"github.com/logrusorgru/aurora/v4"
)

func Error(h, m string) {
	hf := aurora.Red(h + ":").Bold()
	mf := aurora.Red(m)
	fmt.Printf("%s %s\n", hf, mf)
}

func Warning(h, m string) {
	hf := aurora.Yellow(h + ":").Bold()
	mf := aurora.Yellow(m)
	fmt.Printf("%s %s\n", hf, mf)
}

func Info(h, m string) {
	hf := aurora.Cyan(h + ":").Bold()
	mf := aurora.Cyan(m)
	fmt.Printf("%s %s\n", hf, mf)
}

func Success(h, m string) {
	hf := aurora.Green(h + ":").Bold()
	mf := aurora.Green(m)
	fmt.Printf("%s %s\n", hf, mf)
}
