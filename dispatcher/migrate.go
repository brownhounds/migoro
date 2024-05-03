package dispatcher

import (
	"migoro/adapters"
	"migoro/query"
	"migoro/utils"
	"strings"
)

func Migrate() {
	adapter := adapters.Init()

	l := 0
	f := utils.IOReadDir(utils.Env("MIGRATION_DIR"))
	h := utils.MakeRandom()

	for _, file := range f {
		if !utils.InSliceOfStructs(query.GetMigrations(adapter.GetMigrationsQuery()), "MigrationFile", file) {

			m := utils.GetFileContent(file)
			c := strings.TrimSpace(utils.GetStringInBetween(m, "/* UP-START */", "/* UP-END */"))

			if len(c) == 0 {
				utils.Warning("Migration file is empty", file)
				continue
			}

			query.ApplyMigration(c)
			query.WriteMigrationLog(adapter.WriteMigrationLogQuery(), file, h)
			utils.Success("Migration Applied", file)

			l++
		}
	}

	if l == 0 {
		utils.Success("Migrate", "Nothing to migrate")
	}
}
