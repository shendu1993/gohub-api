package make

import (
	"fmt"
	"gohub-api/pkg/app"
	"gohub-api/pkg/console"
	"gohub-api/pkg/file"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeMigration = &cobra.Command{
	Use:   "migration",
	Short: "Create a migration file, example: make migration add_users_table",
	Run:   runMakeMigration,
	Args:  cobra.ExactArgs(1), //值允许一个参数
}

func runMakeMigration(cmd *cobra.Command, args []string) {
	//日期格式化
	timeStr := app.TimenowInTimezone().Format("2006_01_02_150405")
	model := makeModelFromString(args[0])
	fileName := timeStr + "_" + model.PackageName
	//判断文件夹是否存在，不存在就创建一个
	dirPath := "database/migrations"
	if !file.Exists(dirPath) {
		os.MkdirAll(dirPath, os.ModePerm)
	}
	filePath := fmt.Sprintf(dirPath+"/%s.go", fileName)
	createFileFromStub(filePath, "migration", model, map[string]string{"{{FileName}}": fileName})

	console.Success("Migration file created，after modify it, use `migrate up` to migrate database.")
}
