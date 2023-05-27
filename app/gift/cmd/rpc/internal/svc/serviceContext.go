package svc

import (
	"social-im/app/gift/cmd/rpc/internal/config"

	"github.com/zeromicro/go-queue/kq"
)

type ServiceContext struct {
	Config           config.Config
	KqGiftSendClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		KqGiftSendClient: kq.NewPusher(c.KqGiftSendConf.Brokers, c.KqGiftSendConf.Topic),
	}
}
