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
		if utils.InSliceOfStructs(adapter.GetMigrationsFromLog(), file) {
			continue
		}

		m := utils.GetFileContent(file)
		migrationContents := strings.TrimSpace(utils.GetStringInBetween(m, "/* UP-START */", "/* UP-END */"))

		if migrationContents == "" {
			utils.Warning("Migration file is empty", file)
			continue
		}

		utils.Info("Applying Migration", "...")
		fmt.Println(migrationContents)

		query.ApplyMigration(adapter.Connection(), migrationContents)
		adapter.WriteMigrationLog(file, hash)
		utils.Success("Migration Applied", file)
		fmt.Println("")

		l++
	}

	if l == 0 {
		utils.Success("Migrate", "Nothing to migrate")
	}
}
