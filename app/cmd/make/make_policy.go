package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakePolicy = &cobra.Command{
	Use:   "policy",
	Short: " create policy file example : make policy topic",
	Run:   runMakePolicy,
	Args:  cobra.ExactArgs(1),
}

func runMakePolicy(cmd *cobra.Command, args []string) {
	//格式化模型名称
	model := makeModelFromString(args[0])
	//拼接文件路径
	dirPath := "app/policies"
	filePath := fmt.Sprintf(dirPath+"/%s_policy.go", model.PackageName)
	//基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "policy", model)
}
