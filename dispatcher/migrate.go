package dispatcher

import (
	"fmt"
	"migoro/adapters"
	"migoro/error_context"
	"migoro/query"
	"migoro/utils"
	"strings"
)

func Migrate() {
	err, adapter := adapters.Init()
	if err != nil {
		error_context.Context.SetError()
		return
	}

	l := 0
	f := utils.IOReadDir(utils.Env("MIGRATION_DIR"))
	hash := utils.MakeRandom()

	for _, file := range f {
		if !strings.HasSuffix(file, "_"+utils.UP+".sql") {
			continue
		}

		fileNoSuffix, _ := strings.CutSuffix(file, "_"+utils.UP+".sql")

		err, migrations := adapter.GetMigrationsFromLog()
		if err != nil {
			error_context.Context.SetError()
			return
		}

		if utils.InSliceOfStructs(migrations, fileNoSuffix) {
			continue
		}

		migrationContents := strings.TrimSpace(utils.GetMigrationFileContent(file))

		if strings.TrimSpace(migrationContents) == "" {
			utils.Warning("Migration file is empty", file)
			continue
		}

		utils.Info("Applying Migration", "...")
		fmt.Println(migrationContents)

		err, con := adapter.Connection()
		if err != nil {
			error_context.Context.SetError()
			return
		}

		query.ApplyMigration(con, migrationContents)
		adapter.WriteMigrationLog(fileNoSuffix, hash)

		utils.Success("Migration Applied", file)
		fmt.Println("")

		l++
	}

	if l == 0 {
		utils.Success("Migrate", "Nothing to migrate")
	}
}
