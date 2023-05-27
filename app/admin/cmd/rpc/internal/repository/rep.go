package repository

import (
	"social-im/app/admin/cmd/rpc/internal/svc"
	"social-im/app/admin/model"
	"social-im/common/xcache"
	"social-im/common/xcache/global"
	"social-im/common/xorm"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/limit"
	"gorm.io/gorm"
)

type Rep struct {
	svcCtx                   *svc.ServiceContext
	Redis                    redis.UniversalClient
	Mysql                    *gorm.DB //业务库
	RateLimiter              *limit.PeriodLimit
	BackgroundImageModel     model.SysBackImgConfModel
	DictionariesModel        model.SysDictionariesModel
	DictionariesDetailsModel model.SysDictionaryDetailsModel
	GiftMoel                 model.GiftModel
	ProjectConfigModel       model.SysProjectConfigModel
	AreaModel                model.SysAreaModel
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
	rep.BackgroundImageModel = model.NewSysBackImgConfModel(rep.Mysql, svcCtx.Config.Cache)
	rep.DictionariesModel = model.NewSysDictionariesModel(rep.Mysql, svcCtx.Config.Cache)
	rep.DictionariesDetailsModel = model.NewSysDictionaryDetailsModel(rep.Mysql, svcCtx.Config.Cache)
	rep.ProjectConfigModel = model.NewSysProjectConfigModel(rep.Mysql, svcCtx.Config.Cache)
	rep.GiftMoel = model.NewGiftModel(rep.Mysql, svcCtx.Config.Cache)
	rep.AreaModel = model.NewSysAreaModel(rep.Mysql, svcCtx.Config.Cache)
	rep.RateLimiter = limit.NewPeriodLimit(
		svcCtx.Config.TokenRateLimiter.Seconds,
		svcCtx.Config.TokenRateLimiter.Quota,
		xcache.NewRedis(svcCtx.Config.RedisConf.Host, svcCtx.Config.RedisConf.Pass, svcCtx.Config.RedisConf.Type, false),
		"periodlimit:adminTokenRpc:",
		limit.Align(),
	)
	return rep
}
