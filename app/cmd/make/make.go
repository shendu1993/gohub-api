package make

import (
	"embed"
	"fmt"
	"gohub-api/pkg/console"
	"gohub-api/pkg/file"
	"gohub-api/pkg/str"

	"strings"

	"github.com/spf13/cobra"
)

// Package make 命令行的 make 命令

type Model struct {
	TableName          string
	StructName         string
	StructNamePlural   string
	VariableName       string
	VariableNamePlural string
	PackageName        string
	VersionName        string
}

//stubsFs 方便我们后面打包这些 .stub为后缀的文件

//go:embed stubs
var stubsFS embed.FS

// CmdMake 说明cobra命令

var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	//注册 make的子命令
	CmdMake.AddCommand(
		CmdMakeCMD,
		CmdMakeModel,
		CmdMakeAPIController,
		CmdMakeRequest,
		CmdMakeMigration,
		CmdMakeFactory,
	)
}

//makeModelFromString 格式化用户输入的内容
func makeModelFromString(name string, version ...string) Model {
	model := Model{}
	model.StructName = str.Singular(str.Camel(name))
	model.StructNamePlural = str.Plural(model.StructName)
	model.TableName = str.Snake(model.StructNamePlural)
	model.VariableName = str.LowerCamel(model.StructName)
	model.PackageName = str.Snake(model.StructName)
	model.VariableNamePlural = str.LowerCamel(model.StructNamePlural)
	if len(version) > 0 {
		model.VersionName = version[0]
	} else {
		model.VersionName = ""
	}
	return model
}

//createFileFromStub 读取stub 文件并进行变量替换
func createFileFromStub(filePath string, stubName string, model Model, variables ...interface{}) {
	//实现最后一个参数可选
	replaces := make(map[string]string)
	if len(variables) > 0 {
		replaces = variables[0].(map[string]string)
	}
	//目标文件已经存在
	if file.Exists(filePath) {
		console.Exit(filePath + "already exists!")
	}
	//读取 stub模板文件
	modelData, err := stubsFS.ReadFile("stubs/" + stubName + ".stub")
	if err != nil {
		console.Exit(err.Error())
	}
	modelStub := string(modelData)

	//添加默认的替换变量
	replaces["{{VariableName}}"] = model.VariableName
	replaces["{{VariableNamePlural}}"] = model.VariableNamePlural
	replaces["{{StructName}}"] = model.StructName
	replaces["{{StructNamePlural}}"] = model.StructNamePlural
	replaces["{{PackageName}}"] = model.PackageName
	replaces["{{TableName}}"] = model.TableName
	replaces["{{VersionName}}"] = model.VersionName

	//对模板内容做变量替换
	for search, replace := range replaces {
		modelStub = strings.ReplaceAll(modelStub, search, replace)
	}
	//存储到目标文件中
	err = file.Put([]byte(modelStub), filePath)
	{
		if err != nil {
			console.Exit(err.Error())
		}
	}
	//提示成功
	console.Success(fmt.Sprint("[%] created ", filePath))
}
