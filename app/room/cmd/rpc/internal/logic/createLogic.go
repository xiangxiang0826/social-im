package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"social-im/app/room/cmd/rpc/internal/repository"
	"social-im/app/room/cmd/rpc/internal/svc"
	"social-im/app/room/cmd/rpc/internal/types"
	"social-im/app/room/cmd/rpc/pb"
	"social-im/app/room/model"
	"social-im/app/user/cmd/rpc/userrpc"
	"social-im/common/rediskey"
	randUtils "social-im/common/utils/rand"
	timeUtils "social-im/common/utils/time"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

// genericPartyRoomMark 生成房间唯一标识
func genericPartyRoomMark() string {
	pre := strings.ToUpper(randUtils.RandString(types.RAND_CHAR_NUM))
	suffix := strconv.Itoa(randUtils.RandInt(types.RAND_NUM_MIN, types.RAND_NUM_MAX))
	return pre + suffix
}

// validateCreatePartyRoom 验证创建派对房间
func (l *CreateLogic) validateCreatePartyRoom(currentUid, currentDayNum int64) (dayRemaiNum int64, err error) {
	if currentUid <= 0 {
		err = xerr.NewErrMsg("用户ID不存在")
		return
	}
	roomRes, err := l.rep.RoomModel.FindUserPartyRoom(l.ctx, currentUid) //判断该用户当前是否已经有在开启的派对房间
	if err != nil && err != errs.ErrNotFound {
		err = xerr.NewErrMsg(err.Error())
		return
	}
	if err == nil {
		err = xerr.NewErrWithFormatMsg(xerr.USER_ALREADY_REGISTER_ERROR, "uid:%d ", roomRes.CreateUser)
	}
	dayRemaiNum, err = l.rep.ValidateGetUserCreatePartyDayNum(l.ctx, currentUid, currentDayNum)
	if err != nil && err != errs.ErrNotFound {
		err = xerr.NewErrMsg(err.Error())
		return
	}
	if dayRemaiNum <= 0 { // 如果超过每天限制的创建房间数报错
		err = xerr.NewErrWithFormatMsg(xerr.USER_OVER_CREATE_PARTY_NUM_ERROR, "uid:%d ", roomRes.CreateUser)
	}
	return
}

func (l *CreateLogic) Create(in *pb.PartyCreateReq) (*pb.PartyCreateResp, error) {
	currentUid := in.Uid // 当前用户id
	currentDayNum, err := timeUtils.DateDayNum(timeUtils.Now())
	if err != nil {
		return nil, xerr.NewErrMsg(err.Error())
	}
	dayRemaiNum, err := l.validateCreatePartyRoom(currentUid, currentDayNum)
	if err != nil {
		return nil, err
	}
	mark := genericPartyRoomMark() // 生成房间唯一标识
	_, err = l.rep.RoomModel.FindOneByMark(l.ctx, mark)
	if err == nil { //存在标识重新生成一次
		mark = genericPartyRoomMark()
	}
	room := &model.AppRoomMic{}
	room.CreateUser = currentUid
	room.Mark = mark
	room.Name = in.Name
	room.Status = 0
	room.BackgroundUrl = in.BackgroundUrl
	room.BackgroundSmallUrl = in.BackgroundSmallUrl
	room.PartyType = in.PartyType
	roomLimit := &model.AppLimit{}
	roomLimit.LimitType = types.ROOM_LIMIT_TYPE_PARTY
	roomLimit.Uid = currentUid
	roomLimit.CurDayNum = currentDayNum
	roomLimit.CurNums = 1
	roomUser := &model.AppRoomUser{} // 群主加入
	roomUser.UserId = currentUid
	err = l.rep.RoomModel.Transaction(l.ctx, func(db *gorm.DB) error {
		err := l.rep.RoomModel.Insert(l.ctx, db, room)
		if err != nil {
			return err
		}
		err = l.rep.LimitModel.UpsertUserAppLimitWithNumExpr(l.ctx, db, roomLimit)
		if err != nil {
			return err
		}
		roomUser.PartyId = room.Id
		err = l.rep.RoomUserModel.Insert(l.ctx, db, roomUser)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "")
	}
	dayRemaiNum--
	_ = l.SetRoomOwerAvatar(l.rep.Redis, rediskey.CacheSocialImRoomAvatarList+strconv.FormatInt(room.Id, 10), strconv.FormatInt(currentUid, 10))
	rtcToken, _ := l.rep.UserRpc.GetRtcToken(l.ctx, &userrpc.GetRtcTokenReq{
		Uid:         strconv.FormatInt(in.Uid, 10),
		ChannelName: room.Mark,
	})
	return &pb.PartyCreateResp{
		Id:                 room.Id,
		Mark:               mark,
		DayRemaiNum:        dayRemaiNum,
		Name:               room.Name,
		PartyType:          room.PartyType,
		BackgroundUrl:      room.BackgroundUrl,
		BackgroundSmallUrl: room.BackgroundSmallUrl,
		CreatedAt:          timeUtils.Now().Unix(),
		RtcToken:           rtcToken.RtcToken,
	}, nil
}

func (l *CreateLogic) SetRoomOwerAvatar(redisConn redis.UniversalClient, redisListKey, userId string) error {
	userString := redisConn.Get(l.ctx, rediskey.CacheSocialImUserIdPrefix+userId).Val()
	if len(userString) == 0 {
		fmt.Println("error SetRoomOwerAvatar len")
		return nil
	}
	var userAvatar User
	err := json.Unmarshal([]byte(userString), &userAvatar)
	if err != nil {
		fmt.Println("error saveRoomAvatar Unmarshal")
		return err
	}
	err = redisConn.RPush(l.ctx, redisListKey, userAvatar.Avatar).Err()
	if err != nil {
		fmt.Println("error saveRoomAvatar RPush")
		return nil
	}
	return nil
}

func (l *CreateLogic) SaveRoomOwerAvatar(redisConn redis.UniversalClient, userId string) error {
	userAvatar := User{
		Avatar: "http://weixin.com",
	}
	//将monster系列化
	data, err := json.Marshal(&userAvatar)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	err = redisConn.Set(l.ctx, rediskey.CacheSocialImUserIdPrefix+userId, data, time.Second*1000).Err()
    if err != nil {
		fmt.Println("error saveRoomAvatar Unmarshal")
		return err
	}
    return nil
}
