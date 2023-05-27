package svc

import (
	"social-im/app/admin/cmd/rpc/adminrpc"
	"social-im/app/gift/cmd/api/internal/config"
	"social-im/app/gift/cmd/rpc/giftrpc"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	AdminRpc adminrpc.AdminRpc
	GiftRpc  giftrpc.GiftRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		AdminRpc: adminrpc.NewAdminRpc(zrpc.MustNewClient(c.AdminRpcConf)),
		GiftRpc:  giftrpc.NewGiftRpc(zrpc.MustNewClient(c.GiftRpcConf)),
	}
}
