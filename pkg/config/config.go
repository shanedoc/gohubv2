package config

import (
	"gohubv2/pkg/helpers"
	"os"

	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
)

//viper实例
var viper *viperlib.Viper

//todo::configFunc动态加载配置信息
type ConfigFunc func() map[string]interface{}

//ConfigFuncs先加载此数组,loadingConfig再动态生成配置信息
var ConfigFuncs map[string]ConfigFunc

func init() {
	//初始化viper库
	viper = viperlib.New()
	//配置类型 支持 "json", "toml", "yaml", "yml", "properties",
	//             "props", "prop", "env", "dotenv"
	viper.SetConfigType("env")
	//环境变量的配置文件查找路径 相对于 main.go
	viper.AddConfigPath(".")
	//设置环境变量前缀 用以区分go的系统环境
	viper.SetEnvPrefix("appenv")
	//读取环境变量 支持flags
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

//初始化配置信息
func InitConfig(env string) {
	//加载环境变量
	loadEnv(env)

	//注册配置信息
	loadConfig()
}

func loadEnv(envSuffix string) {
	//默认加载.env文件 如果有传参 --env=name 的话 加载 .env.name文件
	envPath := ".env"
	if len(envSuffix) > 0 {
		filepath := ".env" + envSuffix
		//返回描述符对应的文件信息
		if _, err := os.Stat(filepath); err == nil {
			envPath = filepath
		}
	}
	//加载env
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	//监控env文件 变更时随时加载
	viper.WatchConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

//读取环境变量,支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

//新增配置项
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

//获取配置信息
func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue...)
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	// config 或者环境变量不存在的情况
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

//转换类型

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
