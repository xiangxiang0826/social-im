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
	"time"
)

var _ AppRoomMicModel = (*customAppRoomMicModel)(nil)

var (
	cacheRoomPartyListPrex = "cache:room:partyList:list:"
	listCacheExpireTime    = 5 * 60 * time.Second
)

type (
	// AppRoomMicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppRoomMicModel.
	AppRoomMicModel interface {
		appRoomMicModel
		customAppRoomMicLogicModel
	}

	customAppRoomMicModel struct {
		*defaultAppRoomMicModel
	}

	customAppRoomMicLogicModel interface {
		FindUserPartyRoom(ctx context.Context, uid int64) (*AppRoomMic, error)
		PartyListByLastId(ctx context.Context, onlineNums, lastId, size int64) (int64, []*AppRoomMic, error)
		UpsertUserOnlineNums(ctx context.Context, tx *gorm.DB, data *AppRoomMic) error
	}
)

// NewAppRoomMicModel returns a model for the database table.
func NewAppRoomMicModel(conn *gorm.DB, c cache.CacheConf) AppRoomMicModel {
	return &customAppRoomMicModel{
		defaultAppRoomMicModel: newAppRoomMicModel(conn, c),
	}
}

func (m *defaultAppRoomMicModel) customCacheKeys(data *AppRoomMic) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *defaultAppRoomMicModel) FindUserPartyRoom(ctx context.Context, uid int64) (*AppRoomMic, error) {
	var resp AppRoomMic
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&AppRoomMic{}).Where("`create_user` = ?", uid).First(&resp).Error
	})
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultAppRoomMicModel) PartyListByLastId(ctx context.Context, onlineNums, lastId, size int64) (int64, []*AppRoomMic, error) {
	cacheGvaSysBackImgConfListKey := getPartyListCacheKey(onlineNums, lastId, size)
	var resp []*AppRoomMic
	var count int64
	var onlineNumsWhere xorm.GormWhere
	model := AppRoomMic{}
	err := m.QueryWithExpireCtx(ctx, &resp, cacheGvaSysBackImgConfListKey, constant.PARTY_LIST_CACHE_DURING_TIME, func(conn *gorm.DB, v interface{}) error {
		statusWhere := xorm.Where("status = ?", 0)
		if lastId > 0 {
			onlineNumsWhere = xorm.Where("online_nums <= ?", onlineNums)
		}
		return xorm.ListWithLastId(ctx, conn, &resp, model, lastId, size, &count, "online_nums DESC", statusWhere, onlineNumsWhere)
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

func getPartyListCacheKey(onlineNums, lastId, size int64) string {
	return functions.BuildRedisKey(cacheRoomPartyListPrex, strconv.FormatInt(onlineNums, 10), strconv.FormatInt(lastId, 10), strconv.FormatInt(size, 10))
}

func (m *defaultAppRoomMicModel) UpsertUserOnlineNums(ctx context.Context, tx *gorm.DB, data *AppRoomMic) error {
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		errs := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"online_nums"}),
		}).Create(&data).Error
		return errs
	}, m.getCacheKeys(data)...)
	return err
}
