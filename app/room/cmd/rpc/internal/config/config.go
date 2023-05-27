package config

import (
	"social-im/common/conf"
	redisConf "social-im/common/xcache/global"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	TokenRateLimiter conf.RateLimitConfig
	RedisConf        redisConf.RedisConfig
	Mysql            gormc.Mysql
	Cache            cache.CacheConf
	AgoraConf        conf.AgoraConf
	AliOssConf       conf.AliOssConf
	AdminRpcConf     zrpc.RpcClientConf
	UserRpcConf      zrpc.RpcClientConf
}
