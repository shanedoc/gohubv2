package redis

import (
	"context"
	"gohubv2/pkg/logger"
	"sync"
	"time"

	redis "github.com/go-redis/redis/v8"
)

//对第三方的 go-redis/redis 库进行封装
//解耦业务代码和第三方包关系

type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

//确保确保全局的 Redis 对象只实例一次
var once sync.Once

//Redis 全局 Redis，使用 db 1
var Reids *RedisClient

//connectRedis 连接redis数据库,设置全局的Redis对象
func ConnectRedis(address, username, password string, db int) {
	once.Do(func() {
		Reids = NewClient(address, username, password, db)
	})
}

//NewClient 创建一个redis连接
func NewClient(address, username, password string, db int) *RedisClient {
	//初始化RedisClient实例
	rds := &RedisClient{}
	//使用默认的context
	rds.Context = context.Background()
	//使用redis中NewClient初始化连接
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})
	err := rds.Ping()
	logger.LogIf(err)
	return rds
}

//测试链接状态
func (rds *RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

//set
func (rds *RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

//get
func (rds *RedisClient) Get(key string) string {
	ret, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		logger.ErrorString("Redis", "Get", err.Error())
		return ""
	}
	return ret
}

//has
func (rds *RedisClient) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Has", err.Error())
		}
		return false
	} else {
		return true
	}
}

//del:支持多个key传参
func (rds *RedisClient) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}
	return true
}

// FlushDB 清空当前 redis db 里的所有数据
func (rds RedisClient) FlushDB() bool {
	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
		logger.ErrorString("Redis", "FlushDB", err.Error())
		return false
	}
	return true
}

// Increment 当参数只有 1 个时，为 key，其值增加 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要增加的值 int64 类型。
func (rds RedisClient) Increment(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Increment", "参数过多")
		return false
	}
	return true
}

// Decrement 当参数只有 1 个时，为 key，其值减去 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要减去的值 int64 类型。
func (rds RedisClient) Decrement(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.DecrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Decrement", "参数过多")
		return false
	}
	return true
}
