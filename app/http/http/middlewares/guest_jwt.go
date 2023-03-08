package middlewares

import (
	"gohubv2/pkg/jwt"
	"gohubv2/pkg/response"

	"github.com/gin-gonic/gin"
)

//强制使用游客身份进行访问

func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) > 0 {
			_, err := jwt.NewJWT().ParseToken(c)
			if err == nil {
				response.Unauthorized(c, "请使用游客身份访问")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
