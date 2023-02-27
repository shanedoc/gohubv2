package main

import (
	"flag"
	"fmt"
	"gohubv2/bootstrap"
	btsConfig "gohubv2/config"
	"gohubv2/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	//加载config目录下的配置文件
	btsConfig.Initialize()
}

func main() {

	//配置初始化、依赖命令行 --env参数
	var env string
	//用指定的名称、默认值、使用信息注册一个string类型flag，并将flag的值保存到p指向的变量。
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)
	//初始化logger
	bootstrap.SetupLogger()
	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	//gin.SetMode(gin.ReleaseMode)

	//初始化db
	bootstrap.SetupDB()

	//初始化redis
	bootstrap.SetUpRedis()

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
