package svc

import (
	"social-im/app/room/cmd/api/internal/config"
	"social-im/app/room/cmd/rpc/roomrpc"
	"social-im/app/user/cmd/rpc/userrpc"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	RoomRpc roomrpc.RoomRpc
	UserRpc userrpc.UserRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		RoomRpc: roomrpc.NewRoomRpc(zrpc.MustNewClient(c.RoomRpcConf)),
		UserRpc: userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
