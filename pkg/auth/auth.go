package auth

import (
	"errors"
	"gohub-api/app/models/user"
	"gohub-api/pkg/logger"

	"github.com/gin-gonic/gin"
)

//Attempt 多账号登录
func Attempt(account string, password string) (user.User, error) {
	userModel := user.GetByMulti(account)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return userModel, nil
}

//LoginByPhone 通过手机号登录
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}
	return userModel, nil
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

//CurrentUID 从gin.context 中获取当前登录用户ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
