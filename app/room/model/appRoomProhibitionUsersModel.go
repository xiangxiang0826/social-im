package model

import (
	"context"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"social-im/common/constant"
	"social-im/common/functions"
	"social-im/common/xorm"
	"strconv"
)

var _ AppRoomProhibitionUsersModel = (*customAppRoomProhibitionUsersModel)(nil)

var (
	cacheRoomProhibitionUsersPrex = "cache:gva:sysBackImgConf:list:"
	prohibitionUserListCacheExpireTime = constant.LIST_CACHE_DURING_TIME
)

type (
	// AppRoomProhibitionUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppRoomProhibitionUsersModel.
	AppRoomProhibitionUsersModel interface {
		appRoomProhibitionUsersModel
		customAppRoomProhibitionUsersLogicModel
	}

	customAppRoomProhibitionUsersModel struct {
		*defaultAppRoomProhibitionUsersModel
	}

	customAppRoomProhibitionUsersLogicModel interface {
		RoomProhibitionUsersList(ctx context.Context, roomId, roomType, lastId, size int64) (int64, []*AppRoomProhibitionUsers, error)
		UpsertRoomProhibitionUsersStatus(ctx context.Context, tx *gorm.DB, data *AppRoomProhibitionUsers) error
	}
)

// NewAppRoomProhibitionUsersModel returns a model for the database table.
func NewAppRoomProhibitionUsersModel(conn *gorm.DB, c cache.CacheConf) AppRoomProhibitionUsersModel {
	return &customAppRoomProhibitionUsersModel{
		defaultAppRoomProhibitionUsersModel: newAppRoomProhibitionUsersModel(conn, c),
	}
}

func (m *defaultAppRoomProhibitionUsersModel) customCacheKeys(data *AppRoomProhibitionUsers) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *defaultAppRoomProhibitionUsersModel) RoomProhibitionUsersList(ctx context.Context, roomId, roomType, lastId, size int64) (int64, []*AppRoomProhibitionUsers, error) {
	cacheRoomProhibitionUsersListKey := getRoomProhibitionUsersCacheKey(roomId, roomType, lastId, size)
	var resp []*AppRoomProhibitionUsers
	var count int64
	model := AppRoomProhibitionUsers{}
	err := m.QueryWithExpireCtx(ctx, &resp, cacheRoomProhibitionUsersListKey, constant.LIST_CACHE_DURING_TIME, func(conn *gorm.DB, v interface{}) error {
		where := xorm.Where("status = ?", 1)
		roomIdWhere := xorm.Where("room_id = ?", roomId)
		roomTypeWhere := xorm.Where("room_type = ?", roomType)
		return xorm.ListWithLastId(ctx, conn, &resp, model, lastId, size, &count, "", where, roomIdWhere, roomTypeWhere)
	})
	switch err {
	case nil:
		return count, resp, nil
	case gormc.ErrNotFound:
		return 0, nil, ErrNotFound
	default:
		return 0, nil, err
	}
}

func (m *defaultAppRoomProhibitionUsersModel) UpsertRoomProhibitionUsersStatus(ctx context.Context, tx *gorm.DB, data *AppRoomProhibitionUsers) error {
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		errs := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"operator_user", "status", "update_time"}),
		}).Create(&data).Error
		return errs
	}, m.getCacheKeys(data)...)
	return err
}

func getRoomProhibitionUsersCacheKey(roomId, roomType, lastId, size int64) string {
	return functions.BuildRedisKey(cacheRoomProhibitionUsersPrex, strconv.FormatInt(roomId, 10), strconv.FormatInt(roomType, 10), strconv.FormatInt(lastId, 10), strconv.FormatInt(size, 10))
}
