package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ SysProjectConfigModel = (*customSysProjectConfigModel)(nil)

type (
	// SysProjectConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysProjectConfigModel.
	SysProjectConfigModel interface {
		sysProjectConfigModel
		customSysProjectConfigLogicModel
	}

	customSysProjectConfigModel struct {
		*defaultSysProjectConfigModel
	}

	customSysProjectConfigLogicModel interface {
	}
)

// NewSysProjectConfigModel returns a model for the database table.
func NewSysProjectConfigModel(conn *gorm.DB, c cache.CacheConf) SysProjectConfigModel {
	return &customSysProjectConfigModel{
		defaultSysProjectConfigModel: newSysProjectConfigModel(conn, c),
	}
}

func (m *defaultSysProjectConfigModel) customCacheKeys(data *SysProjectConfig) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
