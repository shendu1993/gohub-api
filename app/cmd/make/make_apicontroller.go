package make

import (
	"fmt"
	"gohub-api/pkg/console"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Create api controller ,exmple:apicontroller v1/user ",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1), //只允许且比必须传值1个参数
}

func runMakeAPIController(cmd *cobra.Command, args []string) {
	//处理参数，要求附带API版本（V1 或者 V2）
	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controller name format: v1/user")
	}
	//apiVersion 用来拼接目标路径
	//name 用来生成cmd.Model实例
	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name, apiVersion)

	//os.MkdirAll 会确保父子目录都会创建，第二个参数为目录权限，使用 0777
	dir := fmt.Sprintf("app/http/controllers/api/%s/", apiVersion)
	os.MkdirAll(dir, os.ModePerm)
	//组建目标目录
	filePath := fmt.Sprintf(dir+"%s_controller.go", model.TableName)
	//基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "apicontroller", model)
}
