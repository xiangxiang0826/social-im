package repository

import (
	"social-im/app/admin/cmd/rpc/adminrpc"
	"social-im/app/gift/cmd/rpc/internal/svc"
	"social-im/app/gift/model"
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
	BalanceModel  model.AppUserBalanceModel
	GiftFlowModel model.AppGiftFlowModel
	GiftBagModel  model.AppGiftBagModel
	UserRpc       userrpc.UserRpc
	AdminRpc      adminrpc.AdminRpc
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
	rep.BalanceModel = model.NewAppUserBalanceModel(rep.Mysql, svcCtx.Config.Cache)
	rep.GiftFlowModel = model.NewAppGiftFlowModel(rep.Mysql, svcCtx.Config.Cache)
	rep.GiftBagModel = model.NewAppGiftBagModel(rep.Mysql, svcCtx.Config.Cache)
	rep.UserRpc = userrpc.NewUserRpc(zrpc.MustNewClient(svcCtx.Config.UserRpcConf))
	rep.AdminRpc = adminrpc.NewAdminRpc(zrpc.MustNewClient(svcCtx.Config.AdminRpcConf))
	rep.RateLimiter = limit.NewPeriodLimit(
		constant.PARTY_USER_ONLINE_NUM_REPORT_PERIOD,
		constant.PARTY_USER_ONLINE_NUM_REPORT_QUOTA,
		xcache.NewRedis(svcCtx.Config.Redis.Host, svcCtx.Config.Redis.Pass, svcCtx.Config.Redis.Type, false),
		"periodlimit:giftTokenRpc:",
		limit.Align(),
	)
	return rep
}
