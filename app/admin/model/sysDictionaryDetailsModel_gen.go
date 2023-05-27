// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"social-im/common/xorm/errs"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (
	cacheGvaSysDictionaryDetailsIdPrefix = "cache:gva:sysDictionaryDetails:id:"
)

type (
	sysDictionaryDetailsModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *SysDictionaryDetails) error

		FindOne(ctx context.Context, id int64) (*SysDictionaryDetails, error)
		Update(ctx context.Context, tx *gorm.DB, data *SysDictionaryDetails) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultSysDictionaryDetailsModel struct {
		gormc.CachedConn
		table string
	}

	SysDictionaryDetails struct {
		Id              int64          `gorm:"column:id"`
		CreatedAt       sql.NullTime   `gorm:"column:created_at"`
		UpdatedAt       sql.NullTime   `gorm:"column:updated_at"`
		DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;index"`
		Label           sql.NullString `gorm:"column:label"`             // 展示值
		Value           sql.NullInt64  `gorm:"column:value"`             // 字典值
		Status          sql.NullInt64  `gorm:"column:status"`            // 启用状态
		Sort            sql.NullInt64  `gorm:"column:sort"`              // 排序标记
		SysDictionaryId sql.NullInt64  `gorm:"column:sys_dictionary_id"` // 关联标记
	}
)

func (SysDictionaryDetails) TableName() string {
	return "`sys_dictionary_details`"
}

func newSysDictionaryDetailsModel(conn *gorm.DB, c cache.CacheConf) *defaultSysDictionaryDetailsModel {
	return &defaultSysDictionaryDetailsModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`sys_dictionary_details`",
	}
}

func (m *defaultSysDictionaryDetailsModel) Insert(ctx context.Context, tx *gorm.DB, data *SysDictionaryDetails) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultSysDictionaryDetailsModel) FindOne(ctx context.Context, id int64) (*SysDictionaryDetails, error) {
	gvaSysDictionaryDetailsIdKey := fmt.Sprintf("%s%v", cacheGvaSysDictionaryDetailsIdPrefix, id)
	var resp SysDictionaryDetails
	err := m.QueryCtx(ctx, &resp, gvaSysDictionaryDetailsIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&SysDictionaryDetails{}).Where("`id` = ?", id).First(&resp).Error
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

func (m *defaultSysDictionaryDetailsModel) Update(ctx context.Context, tx *gorm.DB, data *SysDictionaryDetails) error {
	old, err := m.FindOne(ctx, data.Id)
	if err != nil && err != errs.ErrNotFound {
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(data).Error
	}, m.getCacheKeys(old)...)
	return err
}

func (m *defaultSysDictionaryDetailsModel) getCacheKeys(data *SysDictionaryDetails) []string {
	if data == nil {
		return []string{}
	}
	gvaSysDictionaryDetailsIdKey := fmt.Sprintf("%s%v", cacheGvaSysDictionaryDetailsIdPrefix, data.Id)
	cacheKeys := []string{
		gvaSysDictionaryDetailsIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultSysDictionaryDetailsModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		if err == errs.ErrNotFound {
			return nil
		}
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&SysDictionaryDetails{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultSysDictionaryDetailsModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultSysDictionaryDetailsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGvaSysDictionaryDetailsIdPrefix, primary)
}

func (m *defaultSysDictionaryDetailsModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&SysDictionaryDetails{}).Where("`id` = ?", primary).Take(v).Error
}

func (m *defaultSysDictionaryDetailsModel) tableName() string {
	return m.table
}
