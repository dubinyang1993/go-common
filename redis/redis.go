package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"time"
)

var RedisClient *redis.Pool

func Init() {
	initRedis()

	r := RedisClient.Get()
	defer r.Close()

	fmt.Println("Redis conn successful.")
	log.Info("Redis conn successful.")
}

func initRedis() {
	addr := viper.GetString("redis.addr")
	password := viper.GetString("redis.password")
	maxidle := viper.GetInt("redis.maxidle")
	maxactive := viper.GetInt("redis.maxactive")
	timeout := viper.GetDuration("redis.timeout")

	// RedisPool
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     maxidle,
		MaxActive:   maxactive,
		IdleTimeout: timeout * time.Second,
		Dial: func() (redis.Conn, error) {
			options := redis.DialPassword(password) // 有密码时，把options传入下面Dial方法的最后一个参数上
			c, err := redis.Dial("tcp", addr, options)
			if err != nil {
				log.Error("Redis conn failed:", err)
				panic("Redis conn failed:" + err.Error())
			}
			return c, nil
		},
	}
}
