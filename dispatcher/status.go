package dispatcher

import (
	"fmt"
	"strings"

	"github.com/brownhounds/migoro/adapters"
	"github.com/brownhounds/migoro/utils"

	"github.com/logrusorgru/aurora/v4"
)

func Status() {
	adapter := adapters.Init()

	files := utils.GetMigrationFiles()

	l := 0
	for _, file := range files {
		if !strings.HasSuffix(file, utils.UP+".sql") {
			continue
		}

		if len(file) > l {
			l = len(file)
		}
	}

	for _, file := range files {
		if !strings.HasSuffix(file, utils.UP+".sql") {
			continue
		}

		o := strings.Repeat(" ", l-len(file))
		c := strings.TrimSpace(utils.GetMigrationFileContent(file))

		if c == "" {
			fmt.Println(aurora.Yellow(file + o + " EMPTY FILE"))
			continue
		}
		fileNoSuffix, _ := strings.CutSuffix(file, "_"+utils.UP+".sql")
		if utils.InSliceOfStructs(adapter.GetMigrationsFromLog(), fileNoSuffix) {
			fmt.Println(aurora.Green(file + o + " APPLIED"))
		} else {
			fmt.Println(aurora.Red(file + o + " NOT APPLIED"))
		}
	}
}
