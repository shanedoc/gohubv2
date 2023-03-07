package auth

import (
	v1 "gohubv2/app/http/controllers/api/v1"
	"gohubv2/app/requests"
	"gohubv2/pkg/auth"
	"gohubv2/pkg/jwt"
	"gohubv2/pkg/response"

	"github.com/gin-gonic/gin"
)

//login

type LoginController struct {
	v1.BaseController
}

//手机号登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}
	//尝试连接
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在")
	} else {
		//ok
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.CreatedJSON(c, gin.H{
			"token": token,
		})
	}
}
