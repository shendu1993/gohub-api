package cmd

import (
	"gohub-api/pkg/cache"
	"gohub-api/pkg/console"
	"gohub-api/pkg/helpers"

	"github.com/spf13/cobra"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
}
var CmdCacheClear = &cobra.Command{
	Use:   "clear",
	Short: "Clear cache",
	Run:   runCacheClear,
	Args:  cobra.MaximumNArgs(1), //最大允许1个参数
}

func init() {
	// 注册 cache 命令的子命令
	CmdCache.AddCommand(CmdCacheClear)
}

func runCacheClear(cmd *cobra.Command, args []string) {
	if helpers.Empty(args) {
		cache.Flush()
	} else {
		console.Success("cache key:" + args[0])
		cache.Forget(args[0])
	}
	console.Success("Cache cleared.")
}
