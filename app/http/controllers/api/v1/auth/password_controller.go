package auth

import (
	v1 "gohubv2/app/http/controllers/api/v1"
	"gohubv2/app/models/user"
	"gohubv2/app/requests"
	"gohubv2/pkg/response"

	"github.com/gin-gonic/gin"
)

type PasswordController struct {
	v1.BaseController
}

func (pc *PasswordController) ResetPassword(c *gin.Context) {
	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}

// ResetByEmail 使用 Email 和验证码重置密码
func (pc *PasswordController) ResetByEmail(c *gin.Context) {
	request := requests.ResetByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
		return
	}
	userModel := user.GetByPhone(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}
