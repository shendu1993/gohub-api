// Package bootstrap 启动程序功能
package bootstrap

import (
	"fmt"
	"gohub-api/pkg/cache"
	"gohub-api/pkg/config"
)

//SetupCache 缓存
func SetupCache() {
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
	)
	cache.InitWithCacheStore(rds)
}
