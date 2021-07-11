package auth

import (
	"github.com/gomodule/redigo/redis"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"time"
	"wall-bbs/config"
	"wall-bbs/db"
)

var s *service

type service struct {
	AuthDB *gorm.DB
	Redis *redis.Pool
	Cache *cache.Cache
}

func init() {
	s = new(service)
	s.AuthDB = db.Get().DB
	s.Cache = cache.New(time.Minute*10, time.Hour*24)
	s.Redis = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   60,
		IdleTimeout: 10 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", config.Get().Conf.Redis.Addr+config.Get().Conf.Redis.Port,
				redis.DialPassword(config.Get().Conf.Redis.Pass),
				redis.DialDatabase(config.Get().Conf.Redis.DBID),
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