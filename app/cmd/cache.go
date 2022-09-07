package cmd

import (
	"fmt"
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
var CmdCacheForget = &cobra.Command{
	Use:   "forget",
	Short: "Delete redis key, example: cache forget cache-key",
	Run:   runCacheForget,
}

// forget 命令的选项
var cacheKey string

func init() {
	// 注册 cache 命令的子命令
	CmdCache.AddCommand(CmdCacheClear, CmdCacheForget)

	// 设置 cache forget 命令的选项
	// 设置 cache forget 命令的选项
	CmdCacheForget.Flags().StringVarP(&cacheKey, "key", "k", "", "KEY of the cache")
	err := CmdCacheForget.MarkFlagRequired("key")
	if err != nil {
		return
	}
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

func runCacheForget(cmd *cobra.Command, args []string) {
	if helpers.Empty(cache.Get(cacheKey)) {
		console.Error(fmt.Sprintf("Cache key [%s] 不存在.", cacheKey))
	}
	cache.Forget(cacheKey)
	console.Success(fmt.Sprintf("Cache key [%s] deleted.", cacheKey))
}
