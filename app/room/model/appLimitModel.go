package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ AppLimitModel = (*customAppLimitModel)(nil)

type (
	// AppLimitModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppLimitModel.
	AppLimitModel interface {
		appLimitModel
		customAppLimitLogicModel
	}

	customAppLimitModel struct {
		*defaultAppLimitModel
	}

	customAppLimitLogicModel interface {
		UpsertUserAppLimitWithNumExpr(ctx context.Context, tx *gorm.DB, data *AppLimit) error
		UpsertUserAppLimit(ctx context.Context, tx *gorm.DB, data *AppLimit, assignColumns []string) error
	}
)

// NewAppLimitModel returns a model for the database table.
func NewAppLimitModel(conn *gorm.DB, c cache.CacheConf) AppLimitModel {
	return &customAppLimitModel{
		defaultAppLimitModel: newAppLimitModel(conn, c),
	}
}

func (m *defaultAppLimitModel) customCacheKeys(data *AppLimit) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *defaultAppLimitModel) UpsertUserAppLimitWithNumExpr(ctx context.Context, tx *gorm.DB, data *AppLimit) error {
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		errs := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"cur_nums": gorm.Expr("cur_nums + ?", 1)}),
		}).Create(&data).Error
		return errs
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultAppLimitModel) UpsertUserAppLimit(ctx context.Context, tx *gorm.DB, data *AppLimit, assignColumns []string) error {
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		errs := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns(assignColumns),
		}).Create(&data).Error
		return errs
	}, m.getCacheKeys(data)...)
	return err
}
