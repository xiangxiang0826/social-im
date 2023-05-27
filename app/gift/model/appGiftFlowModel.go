package model

import (
	"context"
	"fmt"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ AppGiftFlowModel = (*customAppGiftFlowModel)(nil)

type (
	// AppGiftFlowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppGiftFlowModel.
	AppGiftFlowModel interface {
		appGiftFlowModel
		customAppGiftFlowLogicModel
		BatchInsert(ctx context.Context, tx *gorm.DB, dataList []*AppGiftFlow) error
		FindOneByMessage(ctx context.Context, from, to, giftId, giftCount int64) (*AppGiftFlow, error)
		FindOneByAttrStatus(ctx context.Context, from, to, giftId, giftCount int64) (*AppGiftFlow, error)
	}

	customAppGiftFlowModel struct {
		*defaultAppGiftFlowModel
	}

	customAppGiftFlowLogicModel interface {
	}
)

// NewAppGiftFlowModel returns a model for the database table.
func NewAppGiftFlowModel(conn *gorm.DB, c cache.CacheConf) AppGiftFlowModel {
	return &customAppGiftFlowModel{
		defaultAppGiftFlowModel: newAppGiftFlowModel(conn, c),
	}
}

func (m *defaultAppGiftFlowModel) customCacheKeys(data *AppGiftFlow) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *defaultAppGiftFlowModel) BatchInsert(ctx context.Context, tx *gorm.DB, dataList []*AppGiftFlow) error {
	return tx.CreateInBatches(dataList, len(dataList)).Error
}

func (m *defaultAppGiftFlowModel) FindOneByMessage(ctx context.Context, from, to, giftId, giftCount int64) (*AppGiftFlow, error) {
	gvaXiangxiangAppGiftFlowByMessageKey := fmt.Sprintf("%s_%v_%v_%v_%v", cacheGvaXiangxiangAppGiftFlowByMessagePrefix, from, to, giftId, giftCount)
	var resp AppGiftFlow
	err := m.QueryRowIndexCtx(ctx, &resp, gvaXiangxiangAppGiftFlowByMessageKey, m.formatPrimary, func(conn *gorm.DB, v interface{}) (interface{}, error) {
		if err := conn.Model(&AppGiftFlow{}).Where("`from_uid` = ? and `to_uid` = ? and `gift_id` = ? and `gift_count` = ? and `status`= 0", from, to, giftId, giftCount).Take(&resp).Error; err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAppGiftFlowModel) FindOneByAttrStatus(ctx context.Context, from, to, giftId, giftCount int64) (*AppGiftFlow, error) {
	gvaXiangxiangAppGiftFlowByAttrStatusKey := fmt.Sprintf("%s_%v_%v_%v_%v", cacheGvaXiangxiangAppGiftFlowByAttrStatusPrefix, from, to, giftId, giftCount)
	var resp AppGiftFlow
	err := m.QueryRowIndexCtx(ctx, &resp, gvaXiangxiangAppGiftFlowByAttrStatusKey, m.formatPrimary, func(conn *gorm.DB, v interface{}) (interface{}, error) {
		if err := conn.Model(&AppGiftFlow{}).Where("`from_uid` = ? and `to_uid` = ? and `gift_id` = ? and `gift_count` = ? and `attr_status`= 0", from, to, giftId, giftCount).Take(&resp).Error; err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
