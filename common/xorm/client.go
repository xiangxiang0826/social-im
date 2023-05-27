package xorm

import (
	"errors"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"social-im/common/xorm/gormLogger"
	"time"
)

func GetClient(
	cfg gormc.Mysql,
) *gorm.DB {
	db, err := ConnectMysql(cfg)
	if err != nil {
		log.Printf("mysql connect error: %+v", err)
		panic(err)
	}
	return db
}

func GetDefaultClient(
	cfg gormc.Mysql,
) *gorm.DB {
	db, err := gormc.ConnectMysql(cfg)
	if err != nil {
		log.Printf("mysql connect error: %+v", err)
		panic(err)
	}
	return db
}

func newDefaultGormLogger(cfg gormc.GormLogConfigI) logger.Interface {
	return gormLogger.New(cfg)
}

func ConnectMysql(m gormc.Mysql) (*gorm.DB, error) {
	if m.Dbname == "" {
		return nil, errors.New("database name is empty")
	}
	mysqlCfg := mysql.Config{
		DSN: m.Dsn(),
	}
	newLogger := newDefaultGormLogger(&m)
	db, err := gorm.Open(mysql.New(mysqlCfg), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	} else {
		sqldb, _ := db.DB()
		sqldb.SetMaxIdleConns(m.MaxIdleConns)
		sqldb.SetMaxOpenConns(m.MaxOpenConns)
		sqldb.SetConnMaxLifetime(60 * time.Second) //连接不活动时的最大生存时间(秒)
		return db, nil
	}
}
