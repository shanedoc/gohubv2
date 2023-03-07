package user

import (
	"gohubv2/app/models"
	"gohubv2/pkg/database"
)

//user model

type User struct {
	//基类模型
	models.BaseModel
	Name string `json:"name,omitempty"`
	//json解释器忽略字段
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`
	//时间戳设置
	models.CommonTimestampsField
}

//创建用户
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}
