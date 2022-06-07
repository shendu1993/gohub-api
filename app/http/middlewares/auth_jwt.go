package middlewares

import (
	"errors"
	"fmt"
	"gohub-api/app/models/user"
	"gohub-api/pkg/config"
	"gohub-api/pkg/jwt"
	"gohub-api/pkg/logger"
	"gohub-api/pkg/response"

	"github.com/gin-gonic/gin"
)

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	// db is now a *DB value
	return userModel
}
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取JWT token
		claims, err := jwt.NewJWT().ParserToken(c)
		// JWT 解析失败，有错误发生
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}
		//JWT 解析成功，设置用户信息
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
			return
		}
		//讲用户信息存入gin.context ,后续auth包将从这里拿到当前用户数据
		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)

		c.Next()
	}
}
