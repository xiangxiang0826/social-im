package repository

import (
	"social-im/app/admin/cmd/rpc/adminrpc"
	"social-im/app/room/cmd/rpc/internal/svc"
	"social-im/app/room/model"
	"social-im/app/user/cmd/rpc/userrpc"
	"social-im/common/constant"
	"social-im/common/xcache"
	"social-im/common/xcache/global"
	"social-im/common/xorm"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type Rep struct {
	svcCtx        *svc.ServiceContext
	Redis         redis.UniversalClient
	Mysql         *gorm.DB //业务库
	RateLimiter   *limit.PeriodLimit
	RoomModel     model.AppRoomMicModel
	LimitModel    model.AppLimitModel
	RoomUserModel model.AppRoomUserModel
	RoomProhibitionUsersModel model.AppRoomProhibitionUsersModel
	AdminRpc      adminrpc.AdminRpc
	UserRpc       userrpc.UserRpc
}

var rep *Rep

func NewRep(svcCtx *svc.ServiceContext) *Rep {
	if rep != nil {
		return rep
	}
	rep = &Rep{
		svcCtx: svcCtx,
		Redis:  xcache.GetClient(svcCtx.Config.RedisConf.RedisConf, global.DB(svcCtx.Config.RedisConf.DB)),
		Mysql:  xorm.GetClient(svcCtx.Config.Mysql),
	}
	rep.RoomModel = model.NewAppRoomMicModel(rep.Mysql, svcCtx.Config.Cache)
	rep.LimitModel = model.NewAppLimitModel(rep.Mysql, svcCtx.Config.Cache)
	rep.RoomUserModel = model.NewAppRoomUserModel(rep.Mysql, svcCtx.Config.Cache)
	rep.RoomProhibitionUsersModel = model.NewAppRoomProhibitionUsersModel(rep.Mysql, svcCtx.Config.Cache)
	rep.AdminRpc = adminrpc.NewAdminRpc(zrpc.MustNewClient(svcCtx.Config.AdminRpcConf))
	rep.UserRpc = userrpc.NewUserRpc(zrpc.MustNewClient(svcCtx.Config.UserRpcConf))
	rep.RateLimiter = limit.NewPeriodLimit(
		constant.PARTY_USER_ONLINE_NUM_REPORT_PERIOD,
		constant.PARTY_USER_ONLINE_NUM_REPORT_QUOTA,
		xcache.NewRedis(svcCtx.Config.Redis.Host, svcCtx.Config.Redis.Pass, svcCtx.Config.Redis.Type, false),
		"periodlimit:roomTokenRpc:",
		limit.Align(),
	)
	return rep
}


