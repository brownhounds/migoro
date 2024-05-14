package utils

import (
	"log"

	"github.com/logrusorgru/aurora/v4"
)

func Error(h, m string) {
	hf := aurora.Red(h + ":").Bold()
	mf := aurora.Red(m)
	log.Printf("%s %s\n", hf, mf)
}

func Warning(h, m string) {
	hf := aurora.Yellow(h + ":").Bold()
	mf := aurora.Yellow(m)
	log.Printf("%s %s\n", hf, mf)
}

func Info(h, m string) {
	hf := aurora.Cyan(h + ":").Bold()
	mf := aurora.Cyan(m)
	log.Printf("%s %s\n", hf, mf)
}

func Success(h, m string) {
	hf := aurora.Green(h + ":").Bold()
	mf := aurora.Green(m)
	log.Printf("%s %s\n", hf, mf)
}

func Notice(h, m string) {
	hf := aurora.White(h + ":").Bold()
	mf := aurora.White(m)
	log.Printf("%s %s\n", hf, mf)
}
