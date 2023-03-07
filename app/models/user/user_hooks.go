package user

import (
	"gohubv2/pkg/hash"

	"gorm.io/gorm"
)

//before save 模型钩子,usermodel创建和更新前进行调用

func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}
	return
}
