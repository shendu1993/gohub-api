package auth

import (
	"errors"
	"gohub-api/app/models/user"
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
