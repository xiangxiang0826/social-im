package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ RoomManagerOnmicerModel = (*customRoomManagerOnmicerModel)(nil)

type (
	// RoomManagerOnmicerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoomManagerOnmicerModel.
	RoomManagerOnmicerModel interface {
		roomManagerOnmicerModel
		customRoomManagerOnmicerLogicModel
	}

	customRoomManagerOnmicerModel struct {
		*defaultRoomManagerOnmicerModel
	}

	customRoomManagerOnmicerLogicModel interface {
	}
)

// NewRoomManagerOnmicerModel returns a model for the database table.
func NewRoomManagerOnmicerModel(conn *gorm.DB, c cache.CacheConf) RoomManagerOnmicerModel {
	return &customRoomManagerOnmicerModel{
		defaultRoomManagerOnmicerModel: newRoomManagerOnmicerModel(conn, c),
	}
}

func (m *defaultRoomManagerOnmicerModel) customCacheKeys(data *RoomManagerOnmicer) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
