// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (
	cacheGvaSysBackImgConfIdPrefix = "cache:gva:sysBackImgConf:id:"
)

type (
	sysBackImgConfModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *SysBackImgConf) error

		FindOne(ctx context.Context, id int64) (*SysBackImgConf, error)
		Update(ctx context.Context, tx *gorm.DB, data *SysBackImgConf) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultSysBackImgConfModel struct {
		gormc.CachedConn
		table string
	}

	SysBackImgConf struct {
		Id        int64          `gorm:"column:id"` // ID
		CreatedAt sql.NullTime   `gorm:"column:created_at"`
		UpdatedAt sql.NullTime   `gorm:"column:updated_at"`
		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
		CreatedBy sql.NullInt64  `gorm:"column:created_by"` // 创建者
		UpdatedBy sql.NullInt64  `gorm:"column:updated_by"` // 更新者
		DeletedBy sql.NullInt64  `gorm:"column:deleted_by"` // 删除者
		Name      string         `gorm:"column:name"`       // 名称
		Url       string         `gorm:"column:url"`        // 文件地址
		SmallUrl  string         `gorm:"column:small_url"`  // 小图地址
		Key       string         `gorm:"column:key"`        // 编号
		Tag       string         `gorm:"column:tag"`        // 标签
		Type      int64          `gorm:"column:type"`
	}
)

func (SysBackImgConf) TableName() string {
	return "`sys_back_img_conf`"
}

func newSysBackImgConfModel(conn *gorm.DB, c cache.CacheConf) *defaultSysBackImgConfModel {
	return &defaultSysBackImgConfModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`sys_back_img_conf`",
	}
}

func (m *defaultSysBackImgConfModel) Insert(ctx context.Context, tx *gorm.DB, data *SysBackImgConf) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultSysBackImgConfModel) FindOne(ctx context.Context, id int64) (*SysBackImgConf, error) {
	gvaSysBackImgConfIdKey := fmt.Sprintf("%s%v", cacheGvaSysBackImgConfIdPrefix, id)
	var resp SysBackImgConf
	err := m.QueryCtx(ctx, &resp, gvaSysBackImgConfIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&SysBackImgConf{}).Where("`id` = ?", id).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysBackImgConfModel) Update(ctx context.Context, tx *gorm.DB, data *SysBackImgConf) error {
	old, err := m.FindOne(ctx, data.Id)
	if err != nil && err != ErrNotFound {
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

func (m *defaultSysBackImgConfModel) getCacheKeys(data *SysBackImgConf) []string {
	if data == nil {
		return []string{}
	}
	gvaSysBackImgConfIdKey := fmt.Sprintf("%s%v", cacheGvaSysBackImgConfIdPrefix, data.Id)
	cacheKeys := []string{
		gvaSysBackImgConfIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultSysBackImgConfModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		if err == ErrNotFound {
			return nil
		}
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&SysBackImgConf{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultSysBackImgConfModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultSysBackImgConfModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGvaSysBackImgConfIdPrefix, primary)
}

func (m *defaultSysBackImgConfModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&SysBackImgConf{}).Where("`id` = ?", primary).Take(v).Error
}

func (m *defaultSysBackImgConfModel) tableName() string {
	return m.table
}
