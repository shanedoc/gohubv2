package middlewares

import (
	"fmt"
	"gohubv2/app/models/user"
	"gohubv2/pkg/config"
	"gohubv2/pkg/jwt"
	"gohubv2/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//从表头中获取token
		claims, err := jwt.NewJWT().ParseToken(ctx)
		if err != nil {
			response.Unauthorized(ctx, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}
		//解析token
		//校验token
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(ctx, "找不到对应用户，用户可能已删除")
			return
		}
		//放置token:将token信息放入gin.Context中,auth包将从这里获取用户信息
		ctx.Set("current_user_id", userModel.GetStringID())
		ctx.Set("current_user_name", userModel.Name)
		ctx.Set("current_user", userModel)
		ctx.Next()
	}
}
