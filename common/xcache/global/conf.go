package global

import "github.com/zeromicro/go-zero/core/stores/redis"

type RedisConfig struct {
	redis.RedisConf
	DB int
}
