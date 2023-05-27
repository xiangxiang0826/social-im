package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ AppUserBaseModel = (*customAppUserBaseModel)(nil)

type (
	// AppUserBaseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppUserBaseModel.
	AppUserBaseModel interface {
		appUserBaseModel
		customAppUserBaseLogicModel
	}

	customAppUserBaseModel struct {
		*defaultAppUserBaseModel
	}

	customAppUserBaseLogicModel interface {
	}
)

// NewAppUserBaseModel returns a model for the database table.
func NewAppUserBaseModel(conn *gorm.DB, c cache.CacheConf) AppUserBaseModel {
	return &customAppUserBaseModel{
		defaultAppUserBaseModel: newAppUserBaseModel(conn, c),
	}
}

func (m *defaultAppUserBaseModel) customCacheKeys(data *AppUserBase) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
