package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//注册路由

func RegisterAPIRoutes(r *gin.Engine) {
	//test v1路由
	v1 := r.Group("/v1")
	{
		//json返回
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})

	}
}
