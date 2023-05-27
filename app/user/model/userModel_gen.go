// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"social-im/common/rediskey"
	"time"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (	
	cacheSocialImUserIdPrefix     = rediskey.CacheSocialImUserIdPrefix	
	cacheSocialImUserMobilePrefix = rediskey.CacheSocialImUserMobilePrefix
)

type (
	userModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *User) error

		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByMobile(ctx context.Context, mobile string) (*User, error)
		Update(ctx context.Context, tx *gorm.DB, data *User) error
		FindOneByUserIdentity(ctx context.Context, realName, identity string) (*User, error)
		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultUserModel struct {
		gormc.CachedConn
		table string
	}

	User struct {
		Id           int64  `gorm:"column:id"`            // 主键id
		Mobile       string `gorm:"column:mobile"`        // 用户名
		Password     string `gorm:"column:password"`      // 密码
		NickName     string `gorm:"column:nick_name"`     // 昵称
		Avatar       string `gorm:"column:avatar"`        // 头像
		Birthday     time.Time `gorm:"column:birthday"`      // 生日
		RegisterTime time.Time  `gorm:"column:register_time"` // 注册时间
		Sex          int64  `gorm:"column:sex"`       // 性别
		RealName     sql.NullString `gorm:"column:real_name"`
		Identity     sql.NullString `gorm:"column:identity"`
		Authtime     sql.NullInt64  `gorm:"column:authtime"` // 实名认证成功的时间
	}
)

func (User) TableName() string {
	return "app_user"
}

func newUserModel(conn *gorm.DB, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`app_user`",
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, tx *gorm.DB, data *User) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	socialImUserIdKey := fmt.Sprintf("%s%v", cacheSocialImUserIdPrefix, id)
	var resp User
	err := m.QueryCtx(ctx, &resp, socialImUserIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&User{}).Where("`id` = ?", id).First(&resp).Error
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

func (m *defaultUserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	socialImUserMobileKey := fmt.Sprintf("%s%v", cacheSocialImUserMobilePrefix, mobile)
	var resp User
	
	err := m.QueryRowIndexCtx(ctx, &resp, socialImUserMobileKey, m.formatPrimary, func(conn *gorm.DB, v interface{}) (interface{}, error) {
		if err := conn.Debug().Model(&User{}).Where("`mobile` = ?", mobile).Take(&resp).Error; err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return &resp, ErrNotFound
	// case context.DeadlineExceeded:
	// 	return nil, ErrNotFound
	default:
		return nil, err
	}
	// return &resp, nil
}

func (m *defaultUserModel) FindOneByUserIdentity(ctx context.Context, realName, identity string) (*User, error) {
	var resp User

	err := m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
		if err := conn.Debug().Model(&User{}).Where("`real_name` = ? And `identity`= ? ", realName, identity).Take(&resp).Error; err != nil {
			return err
		}
		return nil
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return &resp, ErrNotFound
	// case context.DeadlineExceeded:
	// 	return nil, ErrNotFound
	default:
		return nil, err
	}
	// return &resp, nil
}

func (m *defaultUserModel) Update(ctx context.Context, tx *gorm.DB, data *User) error {
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

func (m *defaultUserModel) getCacheKeys(data *User) []string {
	if data == nil {
		return []string{}
	}
	socialImUserIdKey := fmt.Sprintf("%s%v", cacheSocialImUserIdPrefix, data.Id)
	socialImUserMobileKey := fmt.Sprintf("%s%v", cacheSocialImUserMobilePrefix, data.Mobile)
	cacheKeys := []string{
		socialImUserIdKey, socialImUserMobileKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultUserModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
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
		return db.Delete(&User{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultUserModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSocialImUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&User{}).Where("`id` = ?", primary).Take(v).Error
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
