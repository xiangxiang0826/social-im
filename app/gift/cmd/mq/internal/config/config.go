package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf

	Redis redis.RedisConf

	// kq : pub sub
	KqGiftSendConf kq.KqConf

	// rpc
	GiftRpcConf zrpc.RpcClientConf
	UserRpcConf zrpc.RpcClientConf
}
