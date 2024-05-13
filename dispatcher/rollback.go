package dispatcher

import (
	"fmt"
	"migoro/adapters"
	"migoro/query"
	"migoro/utils"
	"strings"
)

func Rollback() {
	adapter := adapters.Init()

	md := adapter.GetLatestMigrationsFromLog()
	var dmi []string

	for _, m := range md {
		f := m.MigrationFile + "_" + utils.DOWN + ".sql"
		if !utils.Exists(f) {
			utils.Error("Rollback", "Migration file is missing: "+f)
			continue
		}

		dmi = append(dmi, f)
	}

	if len(dmi) != 0 {
		for _, file := range dmi {
			migrationContents := strings.TrimSpace(utils.GetMigrationFileContent(file))

			if migrationContents == "" {
				utils.Warning("Rollback", fmt.Sprintf("Script not defined in: %s", file))
				continue
			}

			utils.Info("Rolling Back Migration", "...")
			fmt.Println(migrationContents)

			query.ApplyMigration(adapter.Connection(), migrationContents)
			fileNoSuffix, _ := strings.CutSuffix(file, "_"+utils.DOWN+".sql")
			adapter.CleanMigrationLog(fileNoSuffix)
			utils.Success("Migration Rolled Back", file)
			fmt.Println("")
		}
	} else {
		utils.Success("Rollback", "Nothing to rollback")
	}
}
