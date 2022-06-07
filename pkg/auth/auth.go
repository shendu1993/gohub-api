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

//LoginByPhone 通过手机号登录
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}
	return userModel, nil
}
