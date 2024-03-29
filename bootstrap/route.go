package bootstrap

import (
	"gohub-api/app/http/middlewares"
	"gohub-api/pkg/locale"
	"gohub-api/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {
	//注册全局中间件
	registerGlobalMiddleWare(router)
	//注册Api路由
	routes.RegisterAPIRoutes(router)
	//  配置 404 路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {

	//处理404请求
	router.NoRoute(func(c *gin.Context) {
		//获取表头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")

		if strings.Contains(acceptString, "text/html") {
			//如果是html
			c.String(http.StatusNotFound, locale.Translate(c, "404_default_message"))
		} else {
			//默认返回json格式
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": locale.Translate(c, "404_default_message"),
			})
		}
	})
}
