package routes

import (
	"gohub-api/app/http/controllers/api/v1/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		// 注册一个路由
		v1.GET("/ok", func(c *gin.Context) {
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"Hello":   "World!",
				"name":    "shendu",
				"version": "1",
			})
		})
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			//判断手机是否注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			//判断手机是否注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			//用手机注册
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
			//发送验证码
			vcc := new(auth.VerifyCodeController)
			//图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)

		}

	}

	v2 := r.Group("v2")
	{
		v2.GET("/ok", func(c *gin.Context) {
			content := c.GetHeader("Accept-Language")
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"Hello":   "World!",
				"name":    "shendu",
				"version": "2",
				"content": content,
			})
		})

	}

}
