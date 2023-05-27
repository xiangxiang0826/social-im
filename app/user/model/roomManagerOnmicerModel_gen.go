// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"social-im/common/rediskey"
	"time"
)

var (
	cacheSocialImRoomManagerOnmicerIdPrefix                   = "cache:socialIm:id:"
	cacheSocialImRoomManagerOnmicerRoomIdTypeUidUidTypePrefix = "cache:socialIm:roomId:type:uid:uidType:"
	cacheSocialImRoomManagerOnmicerRoomIdApplyingMicfix       = "cache:socialIm:roomId:applyingMic:"
	cacheSocialImRoomManagerOnmicerRoomIdOnMicfix             = "cache:socialIm:roomId:onMic:"
	cacheSocialImRoomManagerOnmicerRoomIdUserMicStatusfix     = "cache:socialIm:roomId:micStatus:"
	cacheSocialImRoomManagerOnmicerRoomIdInvitedMicfix        = "cache:socialIm:roomId:invitedMic:"
	cacheSocialImRoomInvitedAdminsfix                         = rediskey.CacheSocialImRoomAdminPrefix
	cacheSocialImRoomAdminStatusfix                           = rediskey.CacheSocialImRoomUserAdminStatusPrefix
)

type (
	roomManagerOnmicerModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *RoomManagerOnmicer) error

		FindOne(ctx context.Context, id int64) (*RoomManagerOnmicer, error)
		FindOneByRoomIdTypeUidUidType(ctx context.Context, roomId int64, tp int64, uid int64, uidType int64) (*RoomManagerOnmicer, error)
		Update(ctx context.Context, tx *gorm.DB, data *RoomManagerOnmicer) error
		UpdateMicStatus(ctx context.Context, tx *gorm.DB, uid int64, room int64, micStatus int64) error

		FindUsersApplyingMic(ctx context.Context, room int64) (*[]int64, error)
		FindUsersOnMic(ctx context.Context, room int64) (*[]UserTime, error)
		GetUserMicStatus(ctx context.Context, room int64, user int64) (*RoomManagerOnmicer, error)
		FindUsersInvitedMic(ctx context.Context, room int64) (*[]int64, error)

		FindUsersIsAdmin(ctx context.Context, room int64) (*[]int64, error)
		GetUserAdminStatus(ctx context.Context, room int64, user int64) (*RoomManagerOnmicer, error)
		DeleteUserAdminStatus(ctx context.Context, tx *gorm.DB, uid int64, room int64) error
		UpdateAdminStatus(ctx context.Context, tx *gorm.DB, uid int64, room int64, adminStatus int64) error

		DeleteUserRoomStatus(ctx context.Context, tx *gorm.DB, uid int64, room int64) error
		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		DeleteUserMicStatus(ctx context.Context, tx *gorm.DB, uid int64, room int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultRoomManagerOnmicerModel struct {
		gormc.CachedConn
		table string
	}

	RoomManagerOnmicer struct {
		Id       int64 `gorm:"column:id"`
		RoomId   int64 `gorm:"column:room_id"` // 房间ID
		Type     int64 `gorm:"column:type"`    // 房间类型
		Uid      int64 `gorm:"column:uid"`
		UidType  int64 `gorm:"column:uid_type"`  // 0:表示 uid 是房间管理员UID , 1: 表示uid 是房间中在麦位上的uid
		JoinTime int64 `gorm:"column:join_time"` // 成为管理员/加入麦位的时间戳
	}

	UserTime struct {
		Uid      int64 `gorm:"column:uid"`
		JoinTime int64 `gorm:"column:join_time"`
	}
)

func (RoomManagerOnmicer) TableName() string {
	return "`room_manager_onmicer`"
}

func newRoomManagerOnmicerModel(conn *gorm.DB, c cache.CacheConf) *defaultRoomManagerOnmicerModel {
	return &defaultRoomManagerOnmicerModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`room_manager_onmicer`",
	}
}

func (m *defaultRoomManagerOnmicerModel) Insert(ctx context.Context, tx *gorm.DB, data *RoomManagerOnmicer) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultRoomManagerOnmicerModel) FindOne(ctx context.Context, id int64) (*RoomManagerOnmicer, error) {
	socialImRoomManagerOnmicerIdKey := fmt.Sprintf("%s%v", cacheSocialImRoomManagerOnmicerIdPrefix, id)
	var resp RoomManagerOnmicer
	err := m.QueryCtx(ctx, &resp, socialImRoomManagerOnmicerIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&RoomManagerOnmicer{}).Where("`id` = ?", id).First(&resp).Error
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

func (m *defaultRoomManagerOnmicerModel) FindOneByRoomIdTypeUidUidType(ctx context.Context, roomId int64, tp int64, uid int64, uidType int64) (*RoomManagerOnmicer, error) {
	socialImRoomManagerOnmicerRoomIdTypeUidUidTypeKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheSocialImRoomManagerOnmicerRoomIdTypeUidUidTypePrefix, roomId, tp, uid, uidType)
	var resp RoomManagerOnmicer
	err := m.QueryRowIndexCtx(ctx, &resp, socialImRoomManagerOnmicerRoomIdTypeUidUidTypeKey, m.formatPrimary, func(conn *gorm.DB, v interface{}) (interface{}, error) {
		if err := conn.Model(&RoomManagerOnmicer{}).Where("`room_id` = ? and `type` = ? and `uid` = ? and `uid_type` = ?", roomId, tp, uid, uidType).Take(&resp).Error; err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRoomManagerOnmicerModel) FindUsersInvitedMic(ctx context.Context, room int64) (*[]int64, error) {

	socialImRoomManagerOnmicerUserInvitedMicKey := fmt.Sprintf("%s%v", cacheSocialImRoomManagerOnmicerRoomIdInvitedMicfix, room)
	resp := []int64{}

	err := m.QueryCtx(ctx, &resp, socialImRoomManagerOnmicerUserInvitedMicKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&RoomManagerOnmicer{}).Select("uid").Where("`uid_type` = ? and `room_id` = ?", 3, room).Find(&resp).Error
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

func (m *defaultRoomManagerOnmicerModel) FindUsersApplyingMic(ctx context.Context, room int64) (*[]int64, error) {

	socialImRoomManagerOnmicerUsersApplyingMicKey := fmt.Sprintf("%s%v", cacheSocialImRoomManagerOnmicerRoomIdApplyingMicfix, room)
	resp := []int64{}

	err := m.QueryCtx(ctx, &resp, socialImRoomManagerOnmicerUsersApplyingMicKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&RoomManagerOnmicer{}).Select("uid").Where("`uid_type` = ? and `room_id` = ?", 2, room).Find(&resp).Error
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

func (m *defaultRoomManagerOnmicerModel) FindUsersOnMic(ctx context.Context, room int64) (*[]UserTime, error) {

	cacheSocialImRoomManagerOnmicerRoomIdOnMicKey := fmt.Sprintf("%s%v", cacheSocialImRoomManagerOnmicerRoomIdOnMicfix, room)

	resp := []UserTime{}

	err := m.QueryCtx(ctx, &resp, cacheSocialImRoomManagerOnmicerRoomIdOnMicKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&RoomManagerOnmicer{}).Select("uid", "join_time").Where("`uid_type` = ? and `room_id` = ? ", 1, room).Find(&resp).Error
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

func (m *defaultRoomManagerOnmicerModel) GetUserMicStatus(ctx context.Context, room int64, user int64) (*RoomManagerOnmicer, error) {

	cacheSocialImRoomManagerOnmicerRoomIdUserMicStatuskey := fmt.Sprintf("%s%v:%v", cacheSocialImRoomManagerOnmicerRoomIdUserMicStatusfix, room, user)

	var resp RoomManagerOnmicer
	err := m.QueryCtx(ctx, &resp, cacheSocialImRoomManagerOnmicerRoomIdUserMicStatuskey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&RoomManagerOnmicer{}).Where("`room_id` = ? and `uid` = ?  and `uid_type` >=  ?", room, user, 0).Take(&resp).Error

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
func (m *defaultRoomManagerOnmicerModel) FindUsersIsAdmin(ctx context.Context, room int64) (*[]int64, error) {

	socialImRoomManagerInvitedAdminKey := fmt.Sprintf("%s%v", cacheSocialImRoomInvitedAdminsfix, room)
	resp := []int64{}

	err := m.QueryCtx(ctx, &resp, socialImRoomManagerInvitedAdminKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&RoomManagerOnmicer{}).Select("uid").Where("`uid_type` = ? and `room_id` = ?", -1, room).Find(&resp).Error
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

func (m *defaultRoomManagerOnmicerModel) GetUserAdminStatus(ctx context.Context, room int64, user int64) (*RoomManagerOnmicer, error) {

	cacheSocialImRoomManagerOnmicerRoomIdUserAdminStatuskey := fmt.Sprintf("%s%v:%v", cacheSocialImRoomAdminStatusfix, room, user)

	var resp RoomManagerOnmicer
	err := m.QueryCtx(ctx, &resp, cacheSocialImRoomManagerOnmicerRoomIdUserAdminStatuskey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&RoomManagerOnmicer{}).Where("`room_id` = ? and `uid` = ?  and `uid_type` <  ?", room, user, 0).Take(&resp).Error

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

func (m *defaultRoomManagerOnmicerModel) DeleteUserAdminStatus(ctx context.Context, tx *gorm.DB, uid int64, room int64) error {

	err := m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
		if err := conn.Debug().Model(&RoomManagerOnmicer{}).Exec("delete from room_manager_onmicer where  `uid_type` < ? And `room_id` = ? and `uid` = ? ", 0, room, uid).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

func (m *defaultRoomManagerOnmicerModel) DeleteUserRoomStatus(ctx context.Context, tx *gorm.DB, uid int64, room int64) error {

	var delkeys []string
	socialImRoomManagerOnmicerUsersApplyingMicKey := fmt.Sprintf("%s%v", cacheSocialImRoomManagerOnmicerRoomIdApplyingMicfix, room)
	delkeys = append(delkeys, socialImRoomManagerOnmicerUsersApplyingMicKey)

	cacheSocialImRoomManagerOnmicerRoomIdOnMicKey := fmt.Sprintf("%s%v", cacheSocialImRoomManagerOnmicerRoomIdOnMicfix, room)
	delkeys = append(delkeys, cacheSocialImRoomManagerOnmicerRoomIdOnMicKey)
	cacheSocialImRoomManagerOnmicerRoomIdUserMicStatuskey := fmt.Sprintf("%s%v:%v", cacheSocialImRoomManagerOnmicerRoomIdUserMicStatusfix, room, uid)
	delkeys = append(delkeys, cacheSocialImRoomManagerOnmicerRoomIdUserMicStatuskey)
	socialImRoomManagerOnmicerUserInvitedMicKey := fmt.Sprintf("%s%v", cacheSocialImRoomManagerOnmicerRoomIdInvitedMicfix, room)
	delkeys = append(delkeys, socialImRoomManagerOnmicerUserInvitedMicKey)
	socialImRoomManagerInvitedAdminKey := fmt.Sprintf("%s%v", cacheSocialImRoomInvitedAdminsfix, room)
	delkeys = append(delkeys, socialImRoomManagerInvitedAdminKey)
	cacheSocialImRoomManagerOnmicerRoomIdUserAdminStatuskey := fmt.Sprintf("%s%v:%v", cacheSocialImRoomAdminStatusfix, room, uid)
	delkeys = append(delkeys, cacheSocialImRoomManagerOnmicerRoomIdUserAdminStatuskey)

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {

		return conn.Debug().Model(&RoomManagerOnmicer{}).Exec("delete from room_manager_onmicer where `room_id` = ? and `uid` = ? ", room, uid).Error
	}, delkeys...)

	return err
}

func (m *defaultRoomManagerOnmicerModel) UpdateAdminStatus(ctx context.Context, tx *gorm.DB, uid int64, room int64, adminStatus int64) error {

	err := m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
		if err := conn.Debug().Model(&RoomManagerOnmicer{}).Exec("update room_manager_onmicer SET `uid_type` = ? ,`join_time` = ? WHERE `uid` = ? AND `room_id` = ? AND `uid_type` < ?", adminStatus, time.Now().Unix(), uid, room, 0).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (m *defaultRoomManagerOnmicerModel) Update(ctx context.Context, tx *gorm.DB, data *RoomManagerOnmicer) error {
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

func (m *defaultRoomManagerOnmicerModel) getCacheKeys(data *RoomManagerOnmicer) []string {
	if data == nil {
		return []string{}
	}
	socialImRoomManagerOnmicerIdKey := fmt.Sprintf("%s%v", cacheSocialImRoomManagerOnmicerIdPrefix, data.Id)
	socialImRoomManagerOnmicerRoomIdTypeUidUidTypeKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheSocialImRoomManagerOnmicerRoomIdTypeUidUidTypePrefix, data.RoomId, data.Type, data.Uid, data.UidType)
	cacheKeys := []string{
		socialImRoomManagerOnmicerIdKey, socialImRoomManagerOnmicerRoomIdTypeUidUidTypeKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultRoomManagerOnmicerModel) DeleteUserMicStatus(ctx context.Context, tx *gorm.DB, uid int64, room int64) error {

	err := m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
		if err := conn.Debug().Model(&RoomManagerOnmicer{}).Exec("delete from room_manager_onmicer where  `uid_type` >= ? And `room_id` = ? and `uid` = ? ", 0, room, uid).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

func (m *defaultRoomManagerOnmicerModel) UpdateMicStatus(ctx context.Context, tx *gorm.DB, uid int64, room int64, micStatus int64) error {

	err := m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
		if err := conn.Debug().Model(&RoomManagerOnmicer{}).Exec("update room_manager_onmicer SET `uid_type` = ? ,`join_time` = ? WHERE `uid` = ? AND `room_id` = ? AND `uid_type` >= ?", micStatus, time.Now().Unix(), uid, room, 0).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (m *defaultRoomManagerOnmicerModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
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
		return db.Delete(&RoomManagerOnmicer{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultRoomManagerOnmicerModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultRoomManagerOnmicerModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSocialImRoomManagerOnmicerIdPrefix, primary)
}

func (m *defaultRoomManagerOnmicerModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&RoomManagerOnmicer{}).Where("`id` = ?", primary).Take(v).Error
}

func (m *defaultRoomManagerOnmicerModel) tableName() string {
	return m.table
}
