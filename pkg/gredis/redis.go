package gredis

import (
	"encoding/json"
	"gin-blog/pkg/setting"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisConn *redis.Pool

// 创建Redis 链接池
func Setup() error {
	RedisConn = &redis.Pool{
		// 在给定的时间内,允许分配的最大链接数(当为零时,没有限制)
		MaxActive:   setting.RedisSetting.MaxActive,
		// 最大空闲连接数
		MaxIdle:     setting.RedisSetting.MaxIdle,
		// 在给定的时间内将会保持空闲状态,若到达时间限制则关闭连接(当为零时,没有限制)
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		// 提供创建和配置应用程序链连接的一个函数
		Dial: func() (conn redis.Conn, e error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			return c, err
		},
		// 可选的应用程序检查健康功能
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}

func Set(key string, data interface{}, time int) error {
	// 在连接池获取一个活跃连接
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)  // 向redis 服务器发送命令并返回收到的答复
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))  // 将命令返回转为布尔值
	if err != nil {
		return false
	}
	return exists
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))  // 将命令返回转为bytes
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("EYS", "*"+key+"*"))  // 将命令返回转为[]string
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}
	return nil
}
