package timeUtils

import (
	"strconv"
	"time"
)

const (
	TimeFormat   = "2006-01-02 15:04:05"
	TimeMSFormat = "2006-01-02 15:04:05.000"
	DateFormat   = "2006-01-02"
	TimeTZFormat = `2006-01-02T15:04:05Z`
	DateFormatDayNum = "20060102"
)

func StrDatetime(t time.Time) string {
	return t.Format(TimeFormat)
}

func StrDatetimeMs(t time.Time) string {
	return t.Format(TimeMSFormat)
}

func StrDate(t time.Time) string {
	return t.Format(DateFormat)
}

func ParseTime(str string) time.Time {
	t, _ := time.Parse(TimeFormat, str)
	return t
}

func DateDayNum(t time.Time) (int64, error) {
	strDateNum := t.Format(DateFormatDayNum)
	num, err := strconv.Atoi(strDateNum)
	if err != nil {
		return 0, err
	}
	return int64(num), nil
}
