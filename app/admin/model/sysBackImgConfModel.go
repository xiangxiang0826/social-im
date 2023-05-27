package model

import (
	"context"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"social-im/common/constant"
	"social-im/common/functions"
	"social-im/common/xorm"
	"strconv"
)

var _ SysBackImgConfModel = (*customSysBackImgConfModel)(nil)

var (
	cacheGvaSysBackImgConfListPrex = "cache:gva:sysBackImgConf:list:"
	listCacheExpireTime            = constant.LIST_CACHE_DURING_TIME
)

type (
	// SysBackImgConfModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysBackImgConfModel.
	SysBackImgConfModel interface {
		sysBackImgConfModel
		customSysBackImgConfLogicModel
	}

	customSysBackImgConfModel struct {
		*defaultSysBackImgConfModel
	}

	customSysBackImgConfLogicModel interface {
		BackGroundImageList(ctx context.Context, lastId, size, backgroundImageType int64) (int64, []*SysBackImgConf, error)
	}
)

// NewSysBackImgConfModel returns a model for the database table.
func NewSysBackImgConfModel(conn *gorm.DB, c cache.CacheConf) SysBackImgConfModel {
	return &customSysBackImgConfModel{
		defaultSysBackImgConfModel: newSysBackImgConfModel(conn, c),
	}
}

func (m *defaultSysBackImgConfModel) customCacheKeys(data *SysBackImgConf) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *defaultSysBackImgConfModel) BackGroundImageList(ctx context.Context, lastId, size, backgroundImageType int64) (int64, []*SysBackImgConf, error) {
	cacheGvaSysBackImgConfListKey := getBackGroundImageListCacheKey(lastId, size, backgroundImageType)
	var resp []*SysBackImgConf
	var count int64
	model := SysBackImgConf{}
	err := m.QueryWithExpireCtx(ctx, &resp, cacheGvaSysBackImgConfListKey, constant.LIST_CACHE_DURING_TIME, func(conn *gorm.DB, v interface{}) error {
		typeWhere := xorm.Where("type = ?", backgroundImageType)
		return xorm.ListWithLastId(ctx, conn, &resp, model, lastId, size, &count, "", typeWhere)
	})
	switch err {
	case nil:
		return count, resp, nil
	case gormc.ErrNotFound:
		return 0, nil, ErrNotFound
	default:
		return 0, nil, err
	}

}

func getBackGroundImageListCacheKey(lastId, size, backgroundImageType int64) string {
	return functions.BuildRedisKey(cacheGvaSysBackImgConfListPrex, strconv.FormatInt(lastId, 10), strconv.FormatInt(size, 10), strconv.FormatInt(backgroundImageType, 10))
}
