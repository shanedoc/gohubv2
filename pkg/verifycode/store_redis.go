package verifycode

import (
	"gohubv2/pkg/app"
	"gohubv2/pkg/config"
	"gohubv2/pkg/redis"
	"time"
)

//Redistore实现verifycode.Store_interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

//set
func (s *RedisStore) Set(key, value string) bool {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}
	return s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime)
}

//get 实现 verifycode.Store interface 的 Get 方法
func (s *RedisStore) Get(key string, clear bool) (value string) {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

//verify
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
