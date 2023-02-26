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
