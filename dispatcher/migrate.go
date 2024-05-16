package dispatcher

import (
	"migoro/adapters"
	"migoro/error_context"
	"migoro/query"
	"migoro/utils"
	"strings"
)

func Migrate() {
	adapter := adapters.Init()

	l := 0
	err, f := utils.IOReadDir(utils.Env("MIGRATION_DIR")) // TODO: String ???
	if err != nil {
		error_context.Context.SetError()
		return
	}

	err, hash := utils.MakeRandom()
	if err != nil {
		error_context.Context.SetError()
		return
	}

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

		err, contents := utils.GetMigrationFileContent(file)
		if err != nil {
			error_context.Context.SetError()
			return
		}

		migrationContents := strings.TrimSpace(contents)

		if strings.TrimSpace(migrationContents) == "" {
			utils.Warning("Migration file is empty", file)
			continue
		}

		utils.Info("Applying Migration", "...")
		utils.Notice("Migration Content", migrationContents)

		err, con := adapter.Connection()
		if err != nil {
			error_context.Context.SetError()
			return
		}

		query.ApplyMigration(con, migrationContents)
		adapter.WriteMigrationLog(fileNoSuffix, hash)

		utils.Success("Migration Applied", file)

		l++
	}

	if l == 0 {
		utils.Success("Migrate", "Nothing to migrate")
	}
}
