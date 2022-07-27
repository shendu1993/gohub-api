package limiter

import (
	"gohub-api/pkg/config"
	"gohub-api/pkg/logger"
	"gohub-api/pkg/redis"
	"strings"

	"github.com/gin-gonic/gin"
	limiterlib "github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

// GetKeyIP 获取 Limitor 的 Key，IP
func GetKeyIP(c *gin.Context) string {
	return c.ClientIP()
}

// GetKeyRouteWithIP Limitor 的 Key，路由+IP，针对单个路由做限流
func GetKeyRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.FullPath() + c.ClientIP())
}

// routeToKeyString 辅助方法，将 URL 中的 / 格式为 -
func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName
}

func CheckRate(c *gin.Context, key string, formatted string) (limiterlib.Context, error) {
	//实例化依赖的limiter包的 limiter.Rate
	var context limiterlib.Context
	rate, err := limiterlib.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogIf(err)
		return context, err
	}
	//初始化存储，使用我们程序里公用的redis.Redis对象
	// 为 limiter 设置前缀，保持redis里面数据的整洁
	store, err := sredis.NewStoreWithOptions(redis.Redis.Client, limiterlib.StoreOptions{
		Prefix: config.GetString("app.name") + ":limiter",
	})
	if err != nil {
		logger.LogIf(err)
		return context, err
	}
	//使用上面的初始化的 limiter.Rate 对象和存储对象
	limiterObj := limiterlib.New(store, rate)
	//获取限流的结果
	if c.GetBool("limiter-once") {
		// peek() 取结果，不增加访问次数
		return limiterObj.Peek(c, key)
	} else {
		//确保多个路由组里调用 LimitIP 进行限流时，只增加一次访问数量
		c.Set("limiter-once", true)
		//Get()取结果且增加访问次数
		return limiterObj.Get(c, key)
	}

}
