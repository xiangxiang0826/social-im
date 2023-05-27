package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"social-im/app/room/cmd/rpc/roomrpc"
	"social-im/app/user/cmd/rpc/internal/config"
	"social-im/app/user/cmd/rpc/userrpc"
)

type ServiceContext struct {
	Config  config.Config
	Redis   *redis.Redis
	RoomRpc roomrpc.RoomRpc
	UserRpc userrpc.UserRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis: redis.MustNewRedis(c.RedisConf, func(r *redis.Redis) {
			r.Type = c.RedisConf.Type
			r.Pass = c.RedisConf.Pass
		}),
		RoomRpc: roomrpc.NewRoomRpc(zrpc.MustNewClient(c.RoomRpcConf)),
		UserRpc: userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
