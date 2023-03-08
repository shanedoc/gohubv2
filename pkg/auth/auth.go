package auth

import (
	"errors"
	"gohubv2/app/models/user"
	"gohubv2/pkg/logger"

	"github.com/gin-gonic/gin"
)

//auth授权包

//尝试登录
func Attempt(email string, password string) (user.User, error) {
	userModel := user.GetByMulti(email)
	if userModel.ID == 0 {
		return user.User{}, errors.New("用户不存在")
	}
	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("用户密码有误")
	}
	return userModel, nil
}

//登录指定用户
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("用户不存在")
	}
	return userModel, nil
}

//获取当前用户信息
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	return userModel
}

//CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
