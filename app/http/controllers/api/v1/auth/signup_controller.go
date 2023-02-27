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
