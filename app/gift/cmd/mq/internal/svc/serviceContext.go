package svc

import (
	"social-im/app/gift/cmd/mq/internal/config"
	"social-im/app/gift/cmd/rpc/giftrpc"
	"social-im/app/user/cmd/rpc/userrpc"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	GiftRpc giftrpc.GiftRpc
	UserRpc userrpc.UserRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		GiftRpc: giftrpc.NewGiftRpc(zrpc.MustNewClient(c.GiftRpcConf)),
		UserRpc: userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
