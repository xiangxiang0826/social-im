package config

import (
	"social-im/common/conf"
	"sync"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

var Micmutex sync.Mutex

type Config struct {
	zrpc.RpcServerConf
	TokenRateLimiter conf.RateLimitConfig
	RedisConf        redis.RedisConf
	Mysql            gormc.Mysql
	Cache            cache.CacheConf
	JwtAuth          conf.JwtAuth
	AgoraConf        conf.AgoraConf
	AliOssConf       conf.AliOssConf
	RoomRpcConf      zrpc.RpcClientConf
	UserRpcConf      zrpc.RpcClientConf
}
