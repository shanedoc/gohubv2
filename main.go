package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	//实例化gin实例
	r := gin.New()
	//中间件
	r.Use(gin.Logger(), gin.Recovery())
	//注册路由
	r.GET("/", func(c *gin.Context) {
		//json格式相应
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	//处理404
	r.NoRoute(func(c *gin.Context) {
		//获取标头信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面返回404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}

	})
	r.Run(":8000")

}
