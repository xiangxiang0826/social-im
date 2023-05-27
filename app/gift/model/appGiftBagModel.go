package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ AppGiftBagModel = (*customAppGiftBagModel)(nil)

type (
	// AppGiftBagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppGiftBagModel.
	AppGiftBagModel interface {
		appGiftBagModel
		customAppGiftBagLogicModel
	}

	customAppGiftBagModel struct {
		*defaultAppGiftBagModel
	}

	customAppGiftBagLogicModel interface {
	}
)

// NewAppGiftBagModel returns a model for the database table.
func NewAppGiftBagModel(conn *gorm.DB, c cache.CacheConf) AppGiftBagModel {
	return &customAppGiftBagModel{
		defaultAppGiftBagModel: newAppGiftBagModel(conn, c),
	}
}

func (m *defaultAppGiftBagModel) customCacheKeys(data *AppGiftBag) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
