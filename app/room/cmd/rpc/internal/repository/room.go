package repository

import (
	"context"
	"social-im/app/admin/cmd/rpc/adminrpc"
	"social-im/app/room/cmd/rpc/internal/types"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"
	"strconv"
)

// ValidateGetUserCreatePartyDayNum 验证获取该用户今日创建派对的次数
func (rep *Rep) ValidateGetUserCreatePartyDayNum(ctx context.Context, currentUid, currentDayNum int64) (dayRemaiNum int64, err error) {
	// 验证用户今日已经创建过派对房次数
	limitRes, err := rep.LimitModel.FindOneByUidCurDayNumLimitType(ctx, currentUid, currentDayNum, types.ROOM_LIMIT_TYPE_PARTY)
	if err != nil && err != errs.ErrNotFound {
		err = xerr.NewErrMsg(err.Error())
		return
	}
	configKey := types.USER_DAY_PARTY_CREATE_NUM_CONFIG_KEY
	// 获取派对房每日创建最大配置值
	projectConfigRes, err := rep.AdminRpc.ProjectConfigDetail(ctx, &adminrpc.ProjectConfigDetailReq{
		ConfigType: int64(types.USER_DAY_PARTY_CREATE_NUM_CONFIG_TYPE_VALUE),
		ConfigKey:  configKey,
	})
	if err != nil {
		err = xerr.NewErrWithFormatMsg(xerr.SERVER_COMMON_ERROR, "执行调用 rpc ProjectConfigDetail 失败 , ConfigKey : %s , err : %v", configKey, err)
		return
	}
	// 获取系统派对每日限额最大配置值 调用rpc
	configNum, err := strconv.ParseInt(projectConfigRes.ProjectConfigInfo.ConfigValue, 10, 64)
	if err != nil {
		err = xerr.NewErrMsg(err.Error())
		return
	}
	dayRemaiNum = configNum
	if limitRes != nil { // 存在记录则取出对比验证
		dayRemaiNum = configNum - limitRes.CurNums
	}
	return
}