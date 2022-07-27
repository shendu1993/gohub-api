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
	Args:  cobra.ExactArgs(2), //第一个参数是模型/表名字 第二个参数是行为名字
}

func runMakeMigration(cmd *cobra.Command, args []string) {
	//日期格式化
	timeStr := app.TimenowInTimezone().Format("2006_01_02_150405")
	//模型名称
	model := makeModelFromString(args[0])
	//行为名称
	actName := args[1]
	fileName := timeStr + "_" + actName
	//判断文件夹是否存在，不存在就创建一个
	dirPath := "database/migrations"
	if !file.Exists(dirPath) {
		os.MkdirAll(dirPath, os.ModePerm)
	}
	filePath := fmt.Sprintf(dirPath+"/%s.go", fileName)
	createFileFromStub(filePath, "migration", model, map[string]string{"{{FileName}}": fileName})

	console.Success("Migration file created，after modify it, use `migrate up` to migrate database.")
}
