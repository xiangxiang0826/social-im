package model

import (
	"context"
	"social-im/common/rediskey"
	"strconv"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ AppRoomUserModel = (*customAppRoomUserModel)(nil)

type (
	// AppRoomUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppRoomUserModel.
	AppRoomUserModel interface {
		appRoomUserModel
		customAppRoomUserLogicModel
	}

	customAppRoomUserModel struct {
		*defaultAppRoomUserModel
	}

	customAppRoomUserLogicModel interface {
		FindOneByRoomId(ctx context.Context, roomId, userId int64) (*AppRoomUser, error)
	}
)

// NewAppRoomUserModel returns a model for the database table.
func NewAppRoomUserModel(conn *gorm.DB, c cache.CacheConf) AppRoomUserModel {
	return &customAppRoomUserModel{
		defaultAppRoomUserModel: newAppRoomUserModel(conn, c),
	}
}

func (m *defaultAppRoomUserModel) customCacheKeys(data *AppRoomUser) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *defaultAppRoomUserModel) FindOneByRoomId(ctx context.Context, roomId, userId int64) (*AppRoomUser, error) {
	socialImRoomUserKey := rediskey.CacheSocialImRoomUser + strconv.FormatInt(roomId, 10) + "-" + strconv.FormatInt(userId, 10)
	var resp AppRoomUser
	err := m.QueryCtx(ctx, &resp, socialImRoomUserKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Debug().Model(&AppRoomUser{}).Where("`party_id` = ? and `user_id` = ?", roomId, userId).Order("id DESC").Last(&resp).Error
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
