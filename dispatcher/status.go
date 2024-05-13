package dispatcher

import (
	"fmt"
	"migoro/adapters"
	"migoro/utils"
	"strings"

	"github.com/logrusorgru/aurora/v4"
)

func Status() {
	adapter := adapters.Init()

	files := utils.IOReadDir(utils.Env("MIGRATION_DIR"))

	l := 0
	for _, file := range files {
		if len(file) > l {
			l = len(file)
		}
	}

	for _, file := range files {
		o := strings.Repeat(" ", l-len(file))
		m := utils.GetFileContent(file)
		c := strings.TrimSpace(utils.GetStringInBetween(m, "/* UP-START */", "/* UP-END */"))

		if c == "" {
			fmt.Println(aurora.Yellow(file + o + " EMPTY FILE"))
			continue
		}
		if utils.InSliceOfStructs(adapter.GetMigrationsFromLog(), file) {
			fmt.Println(aurora.Green(file + o + " APPLIED"))
		} else {
			fmt.Println(aurora.Red(file + o + " NOT APPLIED"))
		}
	}
}
