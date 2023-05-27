package model

import (
	"context"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"social-im/common/functions"
	"strconv"
	"time"
)

var _ SysDictionaryDetailsModel = (*customSysDictionaryDetailsModel)(nil)

const DictionaryDetailList = 24 * 60 * 60 * time.Second

var(
	cacheGvaSysDictionaryDetailsTypeListPrefix = "cache:gva:sysDictionaryDetails:list:"
)

type (
	// SysDictionaryDetailsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysDictionaryDetailsModel.
	SysDictionaryDetailsModel interface {
		sysDictionaryDetailsModel
		customSysDictionaryDetailsLogicModel
	}

	customSysDictionaryDetailsModel struct {
		*defaultSysDictionaryDetailsModel
	}

	customSysDictionaryDetailsLogicModel interface {
		DictionaryDetailList(ctx context.Context, dictionaryId int64) ([]*SysDictionaryDetails, error)
	}
)

// NewSysDictionaryDetailsModel returns a model for the database table.
func NewSysDictionaryDetailsModel(conn *gorm.DB, c cache.CacheConf) SysDictionaryDetailsModel {
	return &customSysDictionaryDetailsModel{
		defaultSysDictionaryDetailsModel: newSysDictionaryDetailsModel(conn, c),
	}
}

func (m *defaultSysDictionaryDetailsModel) customCacheKeys(data *SysDictionaryDetails) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *defaultSysDictionaryDetailsModel) DictionaryDetailList(ctx context.Context, dictionaryId int64) ([]*SysDictionaryDetails, error) {
	cacheGvaSysBackImgConfListKey := getDictionaryDetailListCacheKey(dictionaryId)
	var resp []*SysDictionaryDetails
	err := m.QueryWithExpireCtx(ctx, &resp, cacheGvaSysBackImgConfListKey, DictionaryDetailList, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&SysDictionaryDetails{}).Where("`sys_dictionary_id` = ?", dictionaryId).Find(&resp).Error
	})
	switch err {
	case nil:
		return resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func getDictionaryDetailListCacheKey(dictionaryId int64) string {
	return functions.BuildRedisKey(cacheGvaSysDictionaryDetailsTypeListPrefix, strconv.FormatInt(dictionaryId, 10))
}
