package config

import "gohubv2/pkg/config"

//Package config 站点配置信息

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			//应用名称
			"name": config.Env("APP_NAME", "Gohub"),
			//当前环境 区分多环境
			"env": config.Env("APP_ENV", "production"),
			//是否debug模式
			"debug": config.Env("APP_DEBUG", false),
			//应用服务端口
			"port": config.Env("APP_PORT", "3000"),
			//加密
			"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),
			//用以生成链接
			"url": config.Env("APP_URL", "http://localhost:8001"),
			//设置时区
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
