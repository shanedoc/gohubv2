package auth

import (
	v1 "gohubv2/app/http/controllers/api/v1"
	"gohubv2/pkg/captcha"
	"gohubv2/pkg/logger"
	"gohubv2/pkg/response"

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
