package auth

import (
	v1 "gohub-api/app/http/controllers/api/v1"
	"gohub-api/app/models/user"
	"gohub-api/app/requests"
	"gohub-api/pkg/response"

	"github.com/gin-gonic/gin"
)

//
type PasswordController struct {
	v1.BaseAPIController
}

//通过手机号重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	request := requests.ResetByPhoneRequest{}
	//1.表单校验
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}
	//2.更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Phone = request.Password
		userModel.Save()
		response.Success(c)
	}

}

//通过邮件重置密码
func (pc *PasswordController) ResetByEmail(c *gin.Context) {
	request := requests.ResetByEmailRequest{}
	//1.表单校验
	if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
		return
	}
	//2.更新密码
	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Phone = request.Password
		userModel.Save()
		response.Success(c)
	}

}
