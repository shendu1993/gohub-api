// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "gohub-api/app/http/controllers/api/v1"
	"gohub-api/app/locale"
	"gohub-api/app/models/user"
	"gohub-api/app/requests"
	"gohub-api/pkg/response"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 获取请求参数，并做表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}
	//  检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist":          user.IsPhoneExist(request.Phone),
		"test-exist":     locale.Translate(c, "test"),
		"test-not-exist": locale.Translate(c, "test_not_exist"),
	})
}

// IsEmailExist 检测邮箱是否已注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}
	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}

// SignupUsingPhone 使用手机和验证码进行注册
func (sc *SignupController) SignupUsingPhone(c *gin.Context) {
	// 1.验证表单
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}
	// 2.验证成功，创建数据
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	_user.Create()
	response.Success(c)
}
func (sc *SignupController) SignupUsingEmail(c *gin.Context) {
	//参数校验
	request := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
		return
	}
	//创建数据
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Email,
		Password: request.Password,
	}
	_user.Create()
	response.Success(c)
}
