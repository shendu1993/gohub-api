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
