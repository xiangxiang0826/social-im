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

var _ SysAreaModel = (*customSysAreaModel)(nil)

var (
	cacheGvaSysAreaConfListPrex = "cache:gva:SysAreaConf:list:"
	listCacheAreaExpireTime     = constant.LIST_AREA_CACHE_DURING_TIME
)

type (
	// SysAreaModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysAreaModel.
	SysAreaModel interface {
		sysAreaModel
		customSysAreaLogicModel
	}

	customSysAreaModel struct {
		*defaultSysAreaModel
	}

	customSysAreaLogicModel interface {
		AreaList(ctx context.Context, pid, level, lastId, size int64) (int64, []*SysArea, error)
	}
)

// NewSysAreaModel returns a model for the database table.
func NewSysAreaModel(conn *gorm.DB, c cache.CacheConf) SysAreaModel {
	return &customSysAreaModel{
		defaultSysAreaModel: newSysAreaModel(conn, c),
	}
}

func (m *defaultSysAreaModel) customCacheKeys(data *SysArea) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *defaultSysAreaModel) AreaList(ctx context.Context, pid, level, lastId, size int64) (int64, []*SysArea, error) {
	areaListCacheKey := getAreaListCacheKey(pid, level, lastId, size)
	var resp []*SysArea
	var count int64
	model := SysArea{}
	err := m.QueryWithExpireCtx(ctx, &resp, areaListCacheKey, listCacheAreaExpireTime, func(conn *gorm.DB, v interface{}) error {
		pidWhere := xorm.Where("pid = ?", pid)
		levelWhere := xorm.Where("level = ?", level)
		return xorm.ListWithLastId(ctx, conn, &resp, model, lastId, size, &count, "", pidWhere, levelWhere)
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

func getAreaListCacheKey(pid, level, lastId, size int64) string {
	return functions.BuildRedisKey(cacheGvaSysAreaConfListPrex, strconv.FormatInt(pid, 10), strconv.FormatInt(level, 10), strconv.FormatInt(lastId, 10), strconv.FormatInt(size, 10))
}
