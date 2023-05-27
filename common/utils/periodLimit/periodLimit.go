package periodLimit

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
)

// IsPeriodLimit 是否限流
func IsPeriodLimit(limitInstance *limit.PeriodLimit, key string) (bool, error) {
	code, err := limitInstance.Take(key)
	if err != nil {
		return false, err
	}
	switch code {
	case limit.OverQuota:
		return true, fmt.Errorf("OverQuota key: %s", key)
	case limit.Allowed:
		return false, nil
	case limit.HitQuota:
		return true, fmt.Errorf("HitQuota key: %s", key)
	default:
		logx.Errorf("DefaultQuota key: %v", key) // unknown response, we just let the sms go
		return false, fmt.Errorf("HitQuota key: %s", key)
	}
}