package repository

import (
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/model"
	"social-im/common/xorm"

	"github.com/zeromicro/go-zero/core/limit"
	zeroredis "github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type Rep struct {
	svcCtx *svc.ServiceContext
	// Redis       redis.UniversalClient
	Mysql              *gorm.DB
	RateLimiter        *limit.PeriodLimit
	UserModel          model.UserModel
	RoomManagerOnMicer model.RoomManagerOnmicerModel
	UserBaseModel      model.AppUserBaseModel
	UserFollowerModel  model.UserFollowerModel
}

var rep *Rep

func NewRep(svcCtx *svc.ServiceContext) *Rep {
	if rep != nil {
		return rep
	}
	rep = &Rep{
		svcCtx: svcCtx,
		// Redis:  xcache.GetClient(svcCtx.Config.Redis.RedisConf, global.DB(svcCtx.Config.Redis.DB)),
		Mysql: xorm.GetClient(svcCtx.Config.Mysql),
	}
	// err := rep.Mysql.AutoMigrate(&model.User{})
	// if err != nil {
	// 	panic(err)
	// }
	rep.UserModel = model.NewUserModel(rep.Mysql, svcCtx.Config.Cache)
	rep.RoomManagerOnMicer = model.NewRoomManagerOnmicerModel(rep.Mysql, svcCtx.Config.Cache)
	rep.UserBaseModel = model.NewAppUserBaseModel(rep.Mysql, svcCtx.Config.Cache)
	rep.UserFollowerModel = model.NewUserFollowerModel(rep.Mysql, svcCtx.Config.Cache)
	rep.RateLimiter = limit.NewPeriodLimit(
		svcCtx.Config.TokenRateLimiter.Seconds,
		svcCtx.Config.TokenRateLimiter.Quota,
		newRedis(svcCtx.Config.Redis.Host, svcCtx.Config.Redis.Pass, svcCtx.Config.Redis.Type, false),
		"periodlimit:tokenrpc:",
		limit.Align(),
	)
	return rep
}

func newRedis(addr string, password string, typ string, tls bool) *zeroredis.Redis {
	ops := make([]zeroredis.Option, 0)
	if password != "" {
		ops = append(ops, zeroredis.WithPass(password))
	}
	if typ == "cluster" {
		ops = append(ops, zeroredis.Cluster())
	}
	if tls {
		ops = append(ops, zeroredis.WithTLS())
	}
	return zeroredis.New(addr, ops...)
}
