package verifycode

//验证模块不和特定的存储服务关联
//单独创建store interface

type Store interface {
	//保存验证码
	Set(id string, value string) bool

	//获取验证码
	Get(id string, clear bool) string

	//校验验证码
	Verify(id, answer string, clear bool) bool
}
