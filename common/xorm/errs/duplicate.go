package errs

import (
	"errors"
	m "github.com/go-sql-driver/mysql"
	"github.com/pingcap/parser/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"strings"
)

var ErrNotFound = gorm.ErrRecordNotFound

func DuplicateError(err error) bool {
	if err != nil {
		return strings.HasPrefix(err.Error(), "Error 1062")
	}
	return false
}

func TableNotFound(err error) bool {
	if err == nil {
		return false
	}
	if errMysql, ok := err.(*mysql.SQLError); ok {
		if errMysql.Code == mysql.ErrNoSuchTable {
			return true
		} else {
			logx.Error("errMysql.Code:", errMysql.Code)
		}
	}
	if errMysql, ok := err.(*m.MySQLError); ok {
		if errMysql.Number == mysql.ErrNoSuchTable {
			return true
		} else {
			logx.Error("errMysql.Code:", errMysql.Number)
		}
	}
	return false
}
func RecordNotFound(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, ErrNotFound) {
		return true
	}
	return false
}
