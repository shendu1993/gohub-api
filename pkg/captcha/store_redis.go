package captcha

import (
	"errors"
	"gohub-api/pkg/app"
	"gohub-api/pkg/config"
	"gohub-api/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, value string) error {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))
	//方便本地开发测试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}
	if ok := s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime); !ok {
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}
