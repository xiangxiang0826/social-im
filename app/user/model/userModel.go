package model

import (
	"context"
	"fmt"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		customUserLogicModel
	}

	customUserModel struct {
		*defaultUserModel
	}

	customUserLogicModel interface {
		FindOneData(ctx context.Context, id int64) (*User, error)
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn *gorm.DB, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *defaultUserModel) customCacheKeys(data *User) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

// 自己逻辑业务定义的方法
func (m *defaultUserModel) FindOneData(ctx context.Context, id int64) (*User, error) {
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
