package dispatcher

import (
	"fmt"
	"migoro/adapters"
	"migoro/error_context"
	"migoro/utils"
	"strings"

	"github.com/logrusorgru/aurora/v4"
)

func Status() {
	err, adapter := adapters.Init()
	if err != nil {
		error_context.Context.SetError()
		return
	}

	files := utils.IOReadDir(utils.Env("MIGRATION_DIR"))

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

		err, migrations := adapter.GetMigrationsFromLog()
		if err != nil {
			error_context.Context.SetError()
			return
		}

		if utils.InSliceOfStructs(migrations, fileNoSuffix) {
			fmt.Println(aurora.Green(file + o + " APPLIED"))
		} else {
			fmt.Println(aurora.Red(file + o + " NOT APPLIED"))
		}
	}
}
