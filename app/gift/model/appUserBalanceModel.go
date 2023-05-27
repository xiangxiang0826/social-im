package model

import (
	"context"
	"fmt"
	"social-im/common/xorm"
	"time"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ AppUserBalanceModel = (*customAppUserBalanceModel)(nil)

type (
	// AppUserBalanceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserBalanceModel.
	AppUserBalanceModel interface {
		appUserBalanceModel
		customAppUserBalanceLogicModel
		FindListByUid(ctx context.Context, uid int64) (int64, []*AppUserBalance, error)
	}

	customAppUserBalanceModel struct {
		*defaultAppUserBalanceModel
	}

	customAppUserBalanceLogicModel interface {
	}
)

// NewAppUserBalanceModel returns a model for the database table.
func NewAppUserBalanceModel(conn *gorm.DB, c cache.CacheConf) AppUserBalanceModel {
	return &customAppUserBalanceModel{
		defaultAppUserBalanceModel: newAppUserBalanceModel(conn, c),
	}
}

func (m *defaultAppUserBalanceModel) customCacheKeys(data *AppUserBalance) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *defaultAppUserBalanceModel) FindListByUid(ctx context.Context, uid int64) (int64, []*AppUserBalance, error) {
	gvaXiangxiangAppUserBalanceUidKey := fmt.Sprintf("%s%v", cacheGvaXiangxiangAppUserBalanceUidPrefix, uid)
	var resp []*AppUserBalance
	var count int64
	lastId := int64(0)
	model := AppUserBalance{}
	size := int64(100)
	err := m.QueryWithExpireCtx(ctx, &resp, gvaXiangxiangAppUserBalanceUidKey, 60*time.Second, func(conn *gorm.DB, v interface{}) error {
		listWhere := xorm.Where("`uid`=uid")
		return xorm.ListWithLastId(ctx, conn, &resp, model, lastId, size, &count, "updated_at DESC", listWhere)
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
