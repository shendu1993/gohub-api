package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin 实例
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	// 注册一个路由
	r.GET("/", func(c *gin.Context) {

		// 以 JSON 格式响应:
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})
	//处理404请求
	r.NoRoute(func(c *gin.Context) {
		//获取表头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")

		if strings.Contains(acceptString, "text/html") {
			//如果是html
			c.String(http.StatusNotFound, "shendu-Not Found 404")
		} else {
			//默认返回json格式
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
	// 运行服务
	r.Run(":8090")
}
