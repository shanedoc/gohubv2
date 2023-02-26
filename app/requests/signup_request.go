package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"emial"`
}

func ValidateSignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	//验证规则
	rule := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}
	//定义错误提示
	msg := govalidator.MapData{
		"email": []string{
			"required.Email 为必填项",
			"min.Email 长度需大于 4",
			"max.Email 长度需小于 30",
			"email.Email 格式不正确，请提供有效的邮箱地址",
		},
	}
	//配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rule,
		TagIdentifier: "valid",
		Messages:      msg,
	}
	//验证
	return govalidator.New(opts).ValidateStruct()
}
