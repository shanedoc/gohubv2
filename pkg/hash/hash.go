package hash

import (
	"gohubv2/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

//hash操作

//加密操作
func BcryptHash(password string) string {
	// GenerateFromPassword 的第二个参数是 cost 值。建议大于 12，数值越大耗费时间越长
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogIf(err)

	return string(bytes)
}

//对比明文密码和数据库hash值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//判断字符串是否是hash过字符串
func BcryptIsHashed(str string) bool {
	//bcrypt后字符串长度60
	return len(str) == 60
}
