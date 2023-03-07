package requests

import (
	"gohubv2/app/requests/validator"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

//表单验证
func LoginByPhone(data interface{}, c *gin.Context) map[string][]string {
	rule := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
	}
	msg := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}
	errs := validate(data, rule, msg)
	//手机验证
	_data := data.(*LoginByPhoneRequest)
	errs = validator.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}
