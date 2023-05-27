package gormLogger

import (
	"context"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
	"time"
)

type dbLog struct {
	LogLevel        logger.LogLevel
	SlowQueryDuring time.Duration
}

func New(cfg gormc.GormLogConfigI) *dbLog {
	return &dbLog{SlowQueryDuring: cfg.GetSlowThreshold(), LogLevel: cfg.GetGormLogMode()}
}

func (l *dbLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *dbLog) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	logx.WithContext(ctx).Debugf(msg, data)
}
func (l *dbLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	logx.WithContext(ctx).Infof(msg, data)
}

func (l *dbLog) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	logx.WithContext(ctx).Errorf(msg, data)
}

func (l *dbLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//这块的逻辑可以自己根据业务情况修改
	elapsed := time.Since(begin)
	sql, rows := fc()
	if elapsed >= l.SlowQueryDuring { //慢日志
		logx.WithContext(ctx).WithDuration(elapsed).Slowf("slow sql: %v  row： %v  err: %v", sql, rows, err)
	} else {
		if err != nil {
			logx.WithContext(ctx).Errorf(sql, "Trace sql: %v  row： %v  err: %v", sql, rows, err)
		} else {
			logx.WithContext(ctx).Infof(sql, "Trace sql: %v  row： %v", sql, rows)
		}
	}
}
