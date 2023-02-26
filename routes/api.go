package routes

import (
	"gohubv2/app/http/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

//注册路由

func RegisterAPIRoutes(r *gin.Engine) {
	//test v1路由
	v1 := r.Group("/v1")
	{
		//json返回
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignUpController)
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExists)
		}
	}
}
