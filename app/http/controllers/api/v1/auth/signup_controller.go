package auth

import (
	v1 "gohubv2/app/http/controllers/api/v1"
	"gohubv2/app/models/user"
	"gohubv2/app/requests"
	"gohubv2/pkg/response"

	"github.com/gin-gonic/gin"
)

//用户身份认证

type SignUpController struct {
	v1.BaseController
}

//校验手机存在
func (sc *SignUpController) IsPhoneExists(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}
	//解析json
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExists(request.Phone),
	})

}

//邮箱检测
func (sc *SignUpController) IsEmailExist(c *gin.Context) {
	//初始化请求对象
	request := requests.SignupEmailExistRequest{}
	//json
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}
	//json back
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExists(request.Email),
	})
}

//使用手机验证码注册
func (sc *SignUpController) SignUpUsingPhone(c *gin.Context) {
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignUpUsingPhone); !ok {
		return
	}
	//创建user
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	_user.Create()
	if _user.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": _user,
		})
	} else {
		response.Abort500(c, "创建用户失败")
	}
}

//邮箱验证码注册
func (sc *SignUpController) SignUpUsingEmail(c *gin.Context) {
	request := requests.SignUpUsingEmailRequest{}
	if ok := requests.Validate(c, &request, requests.SignUpUsingEmail); !ok {
		return
	}
	_user := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	_user.Create()
	if _user.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": _user,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}
