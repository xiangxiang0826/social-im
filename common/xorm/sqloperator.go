package xorm

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"social-im/common/xorm/errs"
)

func Transaction(tx *gorm.DB, fs ...func(tx *gorm.DB) error) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var err error
		for _, f := range fs {
			if f != nil {
				err = f(tx)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func DetailByWhere(tx *gorm.DB, model interface{}, wheres ...GormWhere) error {
	tableName := model.(schema.Tabler).TableName()
	tx = tx.Table(tableName)
	for _, where := range wheres {
		tx = tx.Where(where.Where, where.args...)
	}
	err := tx.First(model).Error
	if err != nil {
		// 表不存在
		if errs.TableNotFound(err) {
			// 创建表
			_ = tx.Table(tableName).AutoMigrate(model)
		}
		return err
	}
	return nil
}

// Count 获取数量
func Count(tx *gorm.DB, model interface{}, where string, args ...interface{}) (int64, error) {
	var total int64
	err := tx.Model(model).Where(where, args...).Count(&total).Error
	if errs.TableNotFound(err) {
		_ = tx.AutoMigrate(model)
		err = tx.Model(model).Where(where, args...).Count(&total).Error
	}
	return total, err
}

func Insert(tx *gorm.DB, model interface{}) error {
	tableName := model.(schema.Tabler).TableName()
	err := tx.Table(tableName).Create(model).Error
	if err != nil {
		// 表不存在
		if errs.TableNotFound(err) {
			// 创建表
			err = tx.Table(tableName).AutoMigrate(model)
			if err != nil {
				return err
			} else {
				// 创建记录
				return tx.Table(tableName).Create(model).Error
			}
		} else {
			return err
		}
	}
	return nil
}

func ListWithPaging(
	tx *gorm.DB,
	models interface{},
	model interface{},
	no int, size int,
	where string, args ...interface{}) (int64, error) {
	tableName := model.(schema.Tabler).TableName()
	var count int64
	db := tx.Table(tableName).Where(where, args...)
	db.Count(&count)
	return count, Paging(db, no, size).Find(models).Error
}

func Paging(tx *gorm.DB, no int, size int) *gorm.DB {
	return tx.Offset((no - 1) * size).Limit(size)
}

func LastIDLimit(tx *gorm.DB, lastId int64, limit int64) *gorm.DB {
	if lastId > 0 {
		tx = tx.Where(" id < ? ", lastId)
	}
	tx = tx.Order("id DESC")
	if limit > 0 {
		tx = tx.Limit(int(limit))
	}
	return tx
}

func LastIDLimitByOrder(tx *gorm.DB, lastId int64, limit int64, order string) *gorm.DB {
	if lastId > 0 {
		tx = tx.Where(" id > (?)", lastId)
	}
	tx = tx.Order(order)
	if limit > 0 {
		tx = tx.Limit(int(limit))
	}
	return tx
}

func ListWithLastIdRes(tx *gorm.DB, lastId int64, limit int64, order string) *gorm.DB {
	if len(order) > 0 { //非主键字段倒叙排序
		return LastIDLimitByOrder(tx, lastId, limit, order)
	}
	return LastIDLimit(tx, lastId, limit)
}

func ListWithLastId(
	ctx context.Context,
	tx *gorm.DB,
	models interface{},
	model interface{},
	lastId, size int64, count *int64,
	order string, wheres ...GormWhere) error {
	tableName := model.(schema.Tabler).TableName()
	db := tx.WithContext(ctx).Table(tableName)
	if len(wheres) > 0 {
		for _, where := range wheres {
			db = db.Where(where.Where, where.args...)
		}
	}
	db.Count(count)
	return ListWithLastIdRes(db, lastId, size, order).Find(models).Error
}
