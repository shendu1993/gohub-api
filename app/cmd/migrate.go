package cmd

import (
	"gohub-api/pkg/database/migrations"
	"gohub-api/pkg/migrate"

	"github.com/spf13/cobra"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: " Run database migration",
	//所有 migrate的下面的子命令都会执行一下代码
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unmigrated  migrations",
	Run:   runUp,
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
	)
}

func migrator() *migrate.Migrator {
	//注册 database/migrations 下面的迁移文件
	migrations.Initialize()
	// 初始化 migrator
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}
