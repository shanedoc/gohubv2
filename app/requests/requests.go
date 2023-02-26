package requests

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

//处理请求和表单验证

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

//解析request请求:支持json、query、url、表单
func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	//解析requsest请求
	if err := c.ShouldBind(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return false
	}
	//表单验证
	errs := handler(obj, c)

	//返回结果
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求验证不通过，具体请查看 errors",
			"errors":  errs,
		})
		return false
	}
	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	//配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}
	//验证
	return govalidator.New(opts).ValidateStruct()
}
