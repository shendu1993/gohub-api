package cmd

import (
	"gohub-api/database/migrations"
	"gohub-api/pkg/migrate"

	"github.com/spf13/cobra"
)

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateRollback,
		CmdMigrateRefresh,
		CmdMigrateReset,
		CmdMigrateFresh,
	)
}

func migrator() *migrate.Migrator {
	//注册 database/migrations 下面的迁移文件
	migrations.Initialize()
	// 初始化 migrator
	return migrate.NewMigrator()
}

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

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}

var CmdMigrateRollback = &cobra.Command{
	Use:   "down",
	Short: "Reverse the up command",
	Run:   runDown,
}

func runDown(cmd *cobra.Command, args []string) {
	migrator().Rollback()
}

var CmdMigrateReset = &cobra.Command{
	Use:   "reset",
	Short: "Rollback all database migrations",
	Run:   runReset,
}

func runReset(cmd *cobra.Command, args []string) {
	migrator().Reset()
}

var CmdMigrateRefresh = &cobra.Command{
	Use:   "refresh",
	Short: "Reset and re-run all migrations",
	Run:   runRefresh,
}

func runRefresh(cmd *cobra.Command, args []string) {
	migrator().Refresh()
}

var CmdMigrateFresh = &cobra.Command{
	Use:   "fresh",
	Short: "Drop all tables and re-run all migrations",
	Run:   runRFresh,
}

func runRFresh(cmd *cobra.Command, args []string) {
	migrator().Fresh()
}
