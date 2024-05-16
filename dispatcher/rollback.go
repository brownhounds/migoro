package dispatcher

import (
	"fmt"
	"migoro/adapters"
	"migoro/error_context"
	"migoro/query"
	"migoro/utils"
	"strings"
)

func Rollback() {
	adapter := adapters.Init()

	err, md := adapter.GetLatestMigrationsFromLog()
	if err != nil {
		error_context.Context.SetError()
		return
	}
	var dmi []string

	for _, m := range *md {
		f := m.MigrationFile + "_" + utils.DOWN + ".sql"
		if !utils.Exists(f) {
			utils.Error("Rollback", "Migration file is missing: "+f)
			continue
		}

		dmi = append(dmi, f)
	}

	if len(dmi) != 0 {
		for _, file := range dmi {
			err, contents := utils.GetMigrationFileContent(file)
			if err != nil {
				error_context.Context.SetError()
				return
			}

			migrationContents := strings.TrimSpace(contents)

			if migrationContents == "" {
				utils.Warning("Rollback", fmt.Sprintf("Script not defined in: %s", file))
				continue
			}

			utils.Info("Rolling Back Migration", "...")
			utils.Notice("Migration Content", migrationContents)

			err, con := adapter.Connection()
			if err != nil {
				error_context.Context.SetError()
				return
			}

			query.ApplyMigration(con, migrationContents)
			fileNoSuffix, _ := strings.CutSuffix(file, "_"+utils.DOWN+".sql")
			adapter.CleanMigrationLog(fileNoSuffix)
			utils.Success("Migration Rolled Back", file)
		}
	} else {
		utils.Success("Rollback", "Nothing to rollback")
	}
}
