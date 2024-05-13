package dispatcher

import (
	"fmt"
	"migoro/adapters"
	"migoro/query"
	"migoro/utils"
	"strings"
)

func Migrate() {
	adapter := adapters.Init()

	l := 0
	f := utils.IOReadDir(utils.Env("MIGRATION_DIR"))
	hash := utils.MakeRandom()

	for _, file := range f {
		if !strings.HasSuffix(file, "_"+utils.UP+".sql") {
			continue
		}

		fileNoSuffix, _ := strings.CutSuffix(file, "_"+utils.UP+".sql")

		if utils.InSliceOfStructs(adapter.GetMigrationsFromLog(), fileNoSuffix) {
			continue
		}

		migrationContents := strings.TrimSpace(utils.GetMigrationFileContent(file))

		if strings.TrimSpace(migrationContents) == "" {
			utils.Warning("Migration file is empty", file)
			continue
		}

		utils.Info("Applying Migration", "...")
		fmt.Println(migrationContents)

		query.ApplyMigration(adapter.Connection(), migrationContents)
		adapter.WriteMigrationLog(fileNoSuffix, hash)

		utils.Success("Migration Applied", file)
		fmt.Println("")

		l++
	}

	if l == 0 {
		utils.Success("Migrate", "Nothing to migrate")
	}
}
