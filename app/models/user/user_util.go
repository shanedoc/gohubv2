package user

import "gohubv2/pkg/database"

//模型相关的数据库操作

func IsEmailExists(email string) bool {
	var count int64
	database.DB.Model(&User{}).Where("email=?", email).Count(&count)
	return count > 0
}

func IsPhoneExists(phone string) bool {
	var count int64
	database.DB.Model(&User{}).Where("phone=?", phone).Count(&count)
	return count > 0
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (userModel User) {
	database.DB.
		Where("phone = ?", loginID).
		Or("email = ?", loginID).
		Or("name = ?", loginID).
		First(&userModel)
	return
}

//通过手机号获取用户信息
func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

// GetByEmail 通过 Email 来获取用户
func GetByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)
	return
}
