package model

import (
	"context"
	"fmt"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"social-im/common/xorm/errs"
)

var _ SysDictionariesModel = (*customSysDictionariesModel)(nil)
var (
	cacheGvaAppRoomMicTypePrefix = "cache:gva:sysDictionaries:type:"
)

type (
	// SysDictionariesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysDictionariesModel.
	SysDictionariesModel interface {
		sysDictionariesModel
		customSysDictionariesLogicModel
	}

	customSysDictionariesModel struct {
		*defaultSysDictionariesModel
	}

	customSysDictionariesLogicModel interface {
		FindOneByType(ctx context.Context, dictionaryType string) (*SysDictionaries, error)
	}
)

// NewSysDictionariesModel returns a model for the database table.
func NewSysDictionariesModel(conn *gorm.DB, c cache.CacheConf) SysDictionariesModel {
	return &customSysDictionariesModel{
		defaultSysDictionariesModel: newSysDictionariesModel(conn, c),
	}
}

func (m *defaultSysDictionariesModel) customCacheKeys(data *SysDictionaries) []string {
	if data == nil {
		return []string{}
	}
	gvaSysDictionariesTypeKey := fmt.Sprintf("%s%v", cacheGvaAppRoomMicTypePrefix, data.Type)
	cacheKeys := []string{
		gvaSysDictionariesTypeKey,
	}
	return cacheKeys
}

func (m *defaultSysDictionariesModel) FindOneByType(ctx context.Context, dictionaryType string) (*SysDictionaries, error) {
	gvaSysDictionariesIdKey := fmt.Sprintf("%s%v", cacheGvaAppRoomMicTypePrefix, dictionaryType)
	var resp SysDictionaries
	err := m.QueryCtx(ctx, &resp, gvaSysDictionariesIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&SysDictionaries{}).Where("`type` = ?", dictionaryType).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, errs.ErrNotFound
	default:
		return nil, err
	}
}
