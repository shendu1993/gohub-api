package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CmdMakeModel = &cobra.Command{
	Use:   "model",
	Short: "Create model file, example: make model user ",
	Run:   runMakeModel,
	Args:  cobra.ExactArgs(1),
}

func runMakeModel(cmd *cobra.Command, args []string) {
	//格式化模型名称，返回一个model对象
	model := makeModelFromString(args[0])
	//确保模型的目录存在，例如 `app/models/user`'
	dir := fmt.Sprintf("app/models/%s/", model.PackageName)
	//os.MkdirAll 会确保父子目录都会创建，第二个参数为目录权限，使用 0777
	os.MkdirAll(dir, os.ModePerm)

	//替换变量
	//地址+名字 前缀
	modelNamePre := dir + model.PackageName
	createFileFromStub(modelNamePre+"_model.go", "model/model", model)
	createFileFromStub(modelNamePre+"_util.go", "model/model_util", model)
	createFileFromStub(modelNamePre+"_hooks.go", "model/model_hooks", model)
}
