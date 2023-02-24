package config

import (
	viperlib "github.com/spf13/viper"

)

//配置信息

//viper实例
var viper *viperlib.Viper
//configFunc动态加载配置信息
type ConfigFunc func() map[string]interface{}
//ConfigFuncs先加载此数组,loadingConfig再动态生成配置信息
var ConfigFuncs map[string]ConfigFunc


