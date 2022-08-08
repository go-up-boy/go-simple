package captcha

import (
	"errors"
	"fmt"
	"go-simple/app"
	"go-simple/pkg/config"
	"go-simple/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	keyPrefix string
}

func (s *RedisStore) Set(id string, value string) error {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))

	// 方便本地开发调试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}
	if ok := s.RedisClient.Set(s.keyPrefix + id, value, ExpireTime); !ok {
		return errors.New("无法存储图片验证码答案")
	}

	return nil
}

func (s RedisStore) Get(key string, clear bool) string {
	key = s.keyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	fmt.Println(v)
	return v == answer
}