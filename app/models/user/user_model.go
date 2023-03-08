package user

import (
	"gohubv2/app/models"
	"gohubv2/pkg/database"
	"gohubv2/pkg/hash"
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

//比较验证码
func (userModel *User) ComparePassword(password string) bool {
	return hash.BcryptCheck(password, userModel.Password)
}

//Get 通过 ID 获取用户
func Get(idstr string) (userModel User) {
	database.DB.Where("id", idstr).First(&userModel)
	return
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
