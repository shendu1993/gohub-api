package auth

import (
	v1 "gohub-api/app/http/controllers/api/v1"
	"gohub-api/app/requests"
	"gohub-api/pkg/auth"
	"gohub-api/pkg/jwt"
	"gohub-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseAPIController
}

//通过手机号登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {
	//1.表单校验
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}
	//2.登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在或密码错误")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}

}
