package repository

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"social-im/common/constant"
	"social-im/common/rediskey"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"
	"strconv"
)

// PartyList 派对列表
func (rep *Rep) PartyList(ctx context.Context, log logx.Logger, onlineNums, lastId, pageSize int64) (total int64, list []*PartyInfo, err error) {
	total, listRes, err := rep.RoomModel.PartyListByLastId(ctx, onlineNums, lastId, pageSize)
	if err != nil && err != errs.ErrNotFound {
		err = xerr.NewErrMsg(err.Error())
	}
	if len(listRes) == 0 { // 没数据直接返回
		return
	}
	for _, eachRoomInfo := range listRes {
		var partyInfo PartyInfo
		partyInfo.Id = eachRoomInfo.Id
		partyInfo.Name = eachRoomInfo.Name
		partyInfo.BackgroundSmallUrl = eachRoomInfo.BackgroundSmallUrl
		partyInfo.Mark = eachRoomInfo.Mark
		partyInfo.PartyType = eachRoomInfo.PartyType
		partyInfo.CreatedAt = eachRoomInfo.CreatedAt.Unix()
		partyInfo.IsHot = 0 //产品预留字段默认为不热门
		currentRoomKey := rediskey.CacheSocialImRoomAvatarList + strconv.FormatInt(eachRoomInfo.Id, 10)
		currentOnlineAvatarNums := int(eachRoomInfo.OnlineNums) //在线人数
		userAvatars, err := rep.PartyRoomUserAvatars(ctx, currentRoomKey)
		if err != nil {
			log.WithContext(ctx).Errorf("get PartyRoomUserAvatars err: %s", err.Error())
		}
		currentDisplayNums := 0
		userAvatarsLength := len(userAvatars)
		if userAvatarsLength < constant.PARTY_LIST_USER_MAX_AVATAR_NUM { // 如果当前加入用户数小于显示的头像数那么取头像数量
			if userAvatarsLength <= currentOnlineAvatarNums {
				currentOnlineAvatarNums = len(userAvatars)
				currentDisplayNums = len(userAvatars)
			} else {
				currentDisplayNums = currentOnlineAvatarNums
			}
		} else { // 用户头像数大于等于显示时，看上报在线数
			currentDisplayNums = constant.PARTY_LIST_USER_MAX_AVATAR_NUM
			if userAvatarsLength >= currentOnlineAvatarNums {
				currentDisplayNums = currentOnlineAvatarNums
			}
		}
		partyInfo.UserList = userAvatars
		if currentDisplayNums <= 0 && len(userAvatars) > 0 {
			currentDisplayNums = 1
			partyInfo.UserList = userAvatars[0:currentDisplayNums]
		}
		partyInfo.OnlineNums = int64(currentOnlineAvatarNums)
	    list = append(list, &partyInfo)
	}
	return
}

// ValidateGetUserCreatePartyDayNum 验证获取该用户今日创建派对的次数
func (rep *Rep) PartyRoomUserAvatars(ctx context.Context, currentRoomKey string) (userAvatars []*UserInfo, err error) {
	redisDatas, err := rep.Redis.LRange(ctx, currentRoomKey, -4, -1).Result()
	if err != nil {
		return
	}
	if len(redisDatas) > 0 {
		for _, userAvatar := range redisDatas {
			userAvatars = append(userAvatars, &UserInfo{Avatar: userAvatar})
		}
	}
	return
}
