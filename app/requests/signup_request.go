package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"emial"`
}

func SignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}
	return validate(data, rules, messages)
}

func ValidateSignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	//验证规则
	rule := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}
	//定义错误提示
	msg := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}
	return validate(data, rule, msg)
}

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

//校验phone
func SignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	rule := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	msg := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	return validate(data, rule, msg)
}
