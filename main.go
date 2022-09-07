package main

import (
	"fmt"
	"gohub-api/app/cmd"
	"gohub-api/app/cmd/make"
	"gohub-api/bootstrap"
	btsConfig "gohub-api/config"
	"gohub-api/pkg/config"
	"gohub-api/pkg/console"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"),
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)
			// 初始化 Logger
			bootstrap.SetupLogger()
			// 初始化 DB
			bootstrap.SetupDB()
			//初始化 redis
			bootstrap.SetupRedis()
			// 初始化缓存
			bootstrap.SetupCache()
		},
	}
	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make.CmdMake,
		cmd.CmdMigrate,
		cmd.CmdDBSeed,
	)
	//配置默认运行Web服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)
	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)
	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
