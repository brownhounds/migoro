package dispatcher

import (
	"migoro/adapters"
	"migoro/error_context"
	"migoro/utils"
	"strings"
)

func Status() {
	err, adapter := adapters.Init()
	if err != nil {
		error_context.Context.SetError()
		return
	}

	err, files := utils.IOReadDir(utils.Env("MIGRATION_DIR"))
	if err != nil {
		error_context.Context.SetError()
		return
	}

	for _, file := range files {
		if !strings.HasSuffix(file, utils.UP+".sql") {
			continue
		}

		err, contents := utils.GetMigrationFileContent(file)
		if err != nil {
			error_context.Context.SetError()
			return
		}

		c := strings.TrimSpace(contents)

		if c == "" {
			utils.Warning("EMPTY FILE ", file)
			continue
		}
		fileNoSuffix, _ := strings.CutSuffix(file, "_"+utils.UP+".sql")

		err, migrations := adapter.GetMigrationsFromLog()
		if err != nil {
			error_context.Context.SetError()
			return
		}

		if utils.InSliceOfStructs(migrations, fileNoSuffix) {
			utils.Success("APPLIED    ", file)
		} else {
			utils.Error("NOT APPLIED", file)
		}
	}
}
