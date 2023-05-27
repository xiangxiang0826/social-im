package types

import (
	"fmt"
	"social-im/common/xerr"
)

var WSDataError = xerr.NewErrCodeMsg(3001, "ws data error")
var (
	ErrTokenExpired     = fmt.Errorf("token 过期")
	ErrTokenInvalid     = fmt.Errorf("token 无效")
	ErrTokenMalformed   = fmt.Errorf("token 格式错误")
	ErrTokenNotValidYet = fmt.Errorf("token 还未生效")
	ErrTokenUnknown     = fmt.Errorf("未知错误")
	ErrTokenKicked      = fmt.Errorf("被踢出")
	ErrUserRegistered   = fmt.Errorf("玩家已经注册过了")
	ErrUserNotfound     = fmt.Errorf("玩家还未注册")
	ErrPwdNotMatch      = fmt.Errorf("玩家密码不正确")
	ErrSysBusy          = fmt.Errorf("系统繁忙，请稍后重试~")
	ErrCodeNotMatch     = fmt.Errorf("验证码不正确~")
	ErrPartyNotfound    = fmt.Errorf("派对不存在~")
	ErrPartyCoolTime    = fmt.Errorf("还在冷却期内~")
	ErrPartyLeaved      = fmt.Errorf("已经离开派对了~")
	ErrBalanceShort     = fmt.Errorf("账户余额不足~")
	ErrToList           = fmt.Errorf("礼物接受者错误~")
)

// error code
const (
	ErrCodeOK     = iota // 成功
	ErrCodeFailed        // 失败
	ErrCodeLimit         // 限流
	ErrSysError

	ErrCodeProtoUnmarshal = 400 // proto解析错误
	ErrCodeParams         = 401 // 参数错误
	ErrCodeSysBusy        = -1  // 参数错误
)
