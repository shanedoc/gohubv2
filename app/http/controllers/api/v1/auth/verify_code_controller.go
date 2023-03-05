package auth

import (
	v1 "gohubv2/app/http/controllers/api/v1"
	"gohubv2/app/requests"
	"gohubv2/pkg/captcha"
	"gohubv2/pkg/logger"
	"gohubv2/pkg/response"
	"gohubv2/pkg/verifycode"

	"github.com/gin-gonic/gin"
)

//VerifyCodeController 用户控制器

type VerfiyController struct {
	v1.BaseController
}

//显示图片
func (vc *VerfiyController) ShowCaptcha(c *gin.Context) {
	//生成验证码
	id, b64s, err := captcha.NewCapture().GenerateCaptcha()
	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	logger.LogIf(err)
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

// SendUsingPhone 发送手机验证码
func (vc *VerfiyController) SendUsingPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifCodePhone); !ok {
		return
	}

	// 2. 发送 SMS
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信失败~")
	} else {
		response.Success(c)
	}
}
