package auth

import (
	"github.com/gomodule/redigo/redis"
	"github.com/patrickmn/go-cache"
	"time"
	"wall-bbs/service/config"
	"wall-bbs/service/db"
)

var s *service

type service struct {
	Redis *redis.Pool
	Cache *cache.Cache
}

func Init() {
	s = new(service)
	if err:=db.MySQL().AutoMigrate(&authToken{});err != nil {
		panic(err)
	}
	s.Cache = cache.New(time.Minute*10, time.Hour*24)
	s.Redis = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   60,
		IdleTimeout: 10 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", config.GetRedisConfig().Addr+config.GetRedisConfig().Port,
				redis.DialPassword(config.GetRedisConfig().Pass),
				redis.DialDatabase(config.GetRedisConfig().DBID),
				redis.DialConnectTimeout(3*time.Second),
				redis.DialReadTimeout(3*time.Second),
				redis.DialWriteTimeout(3*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}