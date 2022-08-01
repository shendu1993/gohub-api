package routes

import (
	controllers "gohub-api/app/http/controllers/api/v1"
	"gohub-api/app/http/controllers/api/v1/auth"
	"gohub-api/app/http/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册 API 相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	v1.Use(middlewares.LimitIP("200-H"))
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
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			// 登录
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)

			// 注册用户
			suc := new(auth.SignupController)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)
			// 图片验证码
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
		}

		//用户相关
		uc := new(controllers.UsersController)
		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
		usersGroup := v1.Group("/users")
		{
			usersGroup.GET("", uc.Index)
		}
		//分类相关
		cgc := new(controllers.CategoriesController)
		cgcGroup := v1.Group("/categories")
		{
			//分类详情
			cgcGroup.GET("/:id", middlewares.AuthJWT(), cgc.Show)
			//分类列表
			cgcGroup.GET("", middlewares.AuthJWT(), cgc.Index)
			//创建分类
			cgcGroup.POST("", middlewares.AuthJWT(), cgc.Store)
			//更新分类
			cgcGroup.PUT("/:id", middlewares.AuthJWT(), cgc.Update)
			//删除分类
			cgcGroup.DELETE("/:id", middlewares.AuthJWT(), cgc.Delete)
		}
		//话题相关
		tc := new(controllers.TopicsController)
		tcGroup := v1.Group("/topics")
		{
			//详情
			tcGroup.GET("/:id", middlewares.AuthJWT(), tc.Show)
			//创建
			tcGroup.POST("", middlewares.AuthJWT(), tc.Store)
			//更新
			tcGroup.PUT("/:id", middlewares.AuthJWT(), tc.Update)
			//删除
			tcGroup.DELETE("/:id", middlewares.AuthJWT(), tc.Delete)
			//列表
			tcGroup.GET("", tc.Index)
		}

		//links相关
		lkc := new(controllers.LinksController)
		lkcGroup := v1.Group("/links")
		{
			lkcGroup.GET("", lkc.Index)
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
