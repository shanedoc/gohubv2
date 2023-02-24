package bootstrap

import (
	"gohubv2/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//程序初始化

func SetupRoute(router *gin.Engine) {
	//全局中间件
	registerGlobalMiddleWare(router)
	//注册api路由
	routes.RegisterAPIRoutes(router)
	//配置404路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

//404请求处理
func setup404Handler(router *gin.Engine) {
	//判断请求头中是否包含accept
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			//html请求
			c.String(http.StatusNotFound, "页面未找到")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

}
