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
		f := m.MigrationFile
		if !utils.Exists(f) {
			utils.Error("Rollback", "Migration file is missing: "+f)
			continue
		}

		m := utils.GetFileContent(f)
		migrationContents := strings.TrimSpace(utils.GetStringInBetween(m, "/* DOWN-START */", "/* DOWN-END */"))

		if len(migrationContents) == 0 {
			utils.Warning("Rollback", "Script not defined in: "+f)
			continue
		}

		dmi = append(dmi, f)
	}

	if len(dmi) != 0 {
		for _, file := range dmi {
			m := utils.GetFileContent(file)
			migrationContents := strings.TrimSpace(utils.GetStringInBetween(m, "/* DOWN-START */", "/* DOWN-END */"))

			if len(migrationContents) == 0 {
				utils.Warning("Rollback", fmt.Sprintf("Script not defined in: %s", file))
				continue
			}

			utils.Info("Rolling Back Migration", "...")
			fmt.Println(migrationContents)

			query.ApplyMigration(adapter.Connection(), migrationContents)
			adapter.CleanMigrationLog(file)
			utils.Success("Migration Rolled Back", file)
			fmt.Println("")
		}
	} else {
		utils.Success("Rollback", "Nothing to rollback")
	}
}
