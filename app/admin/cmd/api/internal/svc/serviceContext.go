package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"social-im/app/admin/cmd/api/internal/config"
	"social-im/app/admin/cmd/rpc/adminrpc"
)

type ServiceContext struct {
	Config config.Config
	AdminRpc adminrpc.AdminRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		AdminRpc: adminrpc.NewAdminRpc(zrpc.MustNewClient(c.AdminRpcConf)),
	}
}
