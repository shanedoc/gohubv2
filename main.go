package main

import (
	"fmt"
	"gohubv2/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	//实例化gin实例
	r := gin.New()
	//注册路由
	bootstrap.SetupRoute(r)

	err := r.Run(":8001")
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}

}
