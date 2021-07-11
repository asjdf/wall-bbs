package auth

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"time"
)

func SetCache(key string, val interface{}, expire int) {
	s.Cache.Set(key, val, time.Duration(expire*10e9))
	go setRedisCache(key, expire, val)
}

func GetCache(c *gin.Context, school, key string, expire int) (interface{}, error) {
	data, found := s.Cache.Get(school + "-" + key)
	if !found {
		d, err := getRedisCache(key)
		if err != nil || d == nil {
			return nil, errors.New("cache not found")
		}
		data = d
		SetCache(key, d, expire)
	}
	return data, nil
}

//检验token是否缓存
func GetCacheT(key string, v interface{}, expire int) error {
	var v2 interface{}
	v2, found := s.Cache.Get(key)
	if !found {
		err := getRedisCacheT(key, v)
		if err != nil || v == nil {
			return errors.New("cache not found")
		}
		SetCache(key, v, expire)
		return nil
	} else {
		bb, err := json.Marshal(v2)
		if err != nil {
			return err
		}
		err = json.Unmarshal(bb, v)
		return err
	}
}

func DeleteCache(key string, memF bool, redisF bool) {
	if key == "*" || key == "" {
		return
	}
	if memF {
		s.Cache.Delete(key)
	}
	if redisF {
		conn := s.Redis.Get()
		defer func(conn redis.Conn) {
			_ = conn.Close()
		}(conn)
		_, _ = conn.Do("DEL", key)
	}
}

func setRedisCache(key string, expire int, data interface{}) {
	if key == "" || data == nil {
		return
	}
	b, err := json.Marshal(data)
	if err != nil {
		return
	}
	conn := s.Redis.Get()
	defer func(conn redis.Conn) {
		_ = conn.Close()
	}(conn)
	_, _ = conn.Do("SETEX", key, expire, b)
}

func getRedisCache(key string) (i interface{}, err error) {
	conn := s.Redis.Get()
	defer func(conn redis.Conn) {
		_ = conn.Close()
	}(conn)
	b, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return
	}
	i = new(interface{})
	err = json.Unmarshal(b, i)
	if err != nil {
		return nil, errors.New("nof found")
	}
	return
}

func getRedisCacheT(key string, data interface{}) (err error) {
	conn := s.Redis.Get()
	defer func(conn redis.Conn) {
		_ = conn.Close()
	}(conn)
	b, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return
	}
	err = json.Unmarshal(b, data)
	return
}

