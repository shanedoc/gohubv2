package validator

import (
	"errors"
	"fmt"
	"gohubv2/pkg/database"
	"strings"

	"github.com/thedevsaddam/govalidator"
)

//自定义验证规则和验证器

func init() {
	govalidator.AddCustomRule("not_exists", func(field, rule, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")
		tableName := rng[0]
		dbFiled := rng[1]

		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[2]
		}
		//用户请求数据
		requestValue := value.(string)
		//拼接sql
		query := database.DB.Table(tableName).Where(dbFiled+" = ?", requestValue)
		// 如果传参第三个参数，加上 SQL Where 过滤
		if len(exceptID) > 0 {
			query.Where("id != ?", exceptID)
		}
		//查询数据库
		var count int64
		query.Count(&count)
		//验证不通过，数据库能找到对应的数据
		if count != 0 {
			// 如果有自定义错误消息的话
			if message != "" {
				return errors.New(message)
			}
			// 默认的错误消息
			return fmt.Errorf("%v 已被占用", requestValue)
		}
		// 验证通过
		return nil
	})

}
