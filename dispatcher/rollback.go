package dispatcher

import (
	"migoro/adapters"
	"migoro/query"
	"migoro/utils"
	"strings"
)

func Rollback() {
	adapter := adapters.Init()

	md := query.GetMigrations(adapter.GetLatestMigrationsQuery())
	var dmi []string

	{
		for _, m := range md {
			f := m.MigrationFile
			if !utils.Exists(f) {
				utils.Error("Rollback", "Migration file is missing: "+f)
				continue
			}

			m := utils.GetFileContent(f)
			c := strings.TrimSpace(utils.GetStringInBetween(m, "/* DOWN-START */", "/* DOWN-END */"))

			if len(c) == 0 {
				utils.Warning("Rollback", "Script not defined in: "+f)
				continue
			}

			dmi = append(dmi, f)
		}
	}

	if len(dmi) != 0 {
		for _, f := range dmi {
			m := utils.GetFileContent(f)
			c := strings.TrimSpace(utils.GetStringInBetween(m, "/* DOWN-START */", "/* DOWN-END */"))

			if len(c) == 0 {
				utils.Warning("Rollback", "Script not defined in: "+f)
				continue
			}

			query.ApplyMigration(c)
			query.CleanMigrationLog(adapter.RollbackMigrationLogQuery(), f)
			utils.Success("Migration Rolled Back", f)
		}
	} else {
		utils.Success("Rollback", "Nothing to rollback")
	}
}
