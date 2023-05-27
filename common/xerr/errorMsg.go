package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"
	message[USER_ALREADY_REGISTER_ERROR] = "用户已经注册过"

	message[DB_INFO_ERRO] = "信息已存在"
	message[IDENTITY_INFO_ERRO] = "身份证信息有误"
	message[ROOM_USER_ALREADY_CREATE_ERROR] = "该用户已经创建过派对房还未退出"
    message[ROOM_ALREADY_CREATE_ERROR] = "该房间已经创建"
	message[USER_OVER_CREATE_PARTY_NUM_ERROR] = "用户今日已经超过创建派对房最大次数"
	message[DB_DATA_EMPTY] = "数据表数据为空"
	message[PERIOD_LIMIT_ERROR] = "频率限制,请稍后再试"

}

func GetCommomErrorMsg() string {
	return message[SERVER_COMMON_ERROR]
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return ""
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
