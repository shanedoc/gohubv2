package auth

import (
	"fmt"
	v1 "gohubv2/app/http/controllers/api/v1"
	"gohubv2/app/models/user"
	"gohubv2/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

//用户身份认证

type SignUpController struct {
	v1.BaseController
}

//校验手机存在
func (sc *SignUpController) IsPhoneExists(c *gin.Context) {

	//请求对象
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}
	request := PhoneExistRequest{}

	//解析json
	if err := c.ShouldBindJSON(&request); err != nil {
		//解析失败 返回422
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		//打印
		fmt.Println(err.Error())
		//返回
		return
	}

	fmt.Println(request.Phone)
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExists(request.Phone),
	})

}

//邮箱检测
func (sc *SignUpController) IsEmailExist(c *gin.Context) {
	//初始化请求对象
	request := requests.SignupEmailExistRequest{}
	//json
	if err := c.ShouldBindJSON(&request); err != nil {
		//解析失败
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	//表单验证
	errs := requests.ValidateSignupEmailExist(&request, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}
	//json back
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExists(request.Email),
	})

}
