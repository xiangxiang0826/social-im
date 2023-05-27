package model

import (
	"context"
	"social-im/common/rediskey"
	"social-im/common/xorm"
	"strconv"
	"time"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ GiftModel = (*customGiftModel)(nil)

type (
	// GiftModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGiftModel.
	GiftModel interface {
		giftModel
		customGiftLogicModel
		GiftList(ctx context.Context, index, size int64) (int64, []*Gift, error)
	}

	customGiftModel struct {
		*defaultGiftModel
	}

	customGiftLogicModel interface {
	}
)

// NewGiftModel returns a model for the database table.
func NewGiftModel(conn *gorm.DB, c cache.CacheConf) GiftModel {
	return &customGiftModel{
		defaultGiftModel: newGiftModel(conn, c),
	}
}

func (m *defaultGiftModel) customCacheKeys(data *Gift) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *defaultGiftModel) GiftList(ctx context.Context, lastId, size int64) (int64, []*Gift, error) {
	cacheGiftListKey := rediskey.CacheSocialImGiftListKey + strconv.FormatInt(lastId, 10) + "_" + strconv.FormatInt(size, 10)
	var resp []*Gift
	var count int64

	model := Gift{}
	err := m.QueryWithExpireCtx(ctx, &resp, cacheGiftListKey, 60*time.Second, func(conn *gorm.DB, v interface{}) error {
		deleteWhere := xorm.Where("`deleted_at` IS NULL")
		return xorm.ListWithLastId(ctx, conn, &resp, model, lastId, size, &count, "updated_at DESC", deleteWhere)
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
