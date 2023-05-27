package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"social-im/app/room/cmd/rpc/internal/repository"
	"social-im/app/room/cmd/rpc/internal/svc"
	"social-im/app/room/cmd/rpc/pb"
	"social-im/app/user/cmd/rpc/userrpc"
	"social-im/common/agora"
	"social-im/common/rediskey"
	errTypes "social-im/common/types"
	"social-im/common/xorm/errs"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type TerminateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewTerminateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TerminateLogic {
	return &TerminateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *TerminateLogic) Terminate(in *pb.PartyTerminateReq) (*pb.PartyTerminateResp, error) {
	// todo: add your logic here and delete this line
	//1 判断roomid是否存在
	roomData, err := l.rep.RoomModel.FindOne(l.ctx, in.RoomId) //判断该用户当前是否已经有在开启的派对房间
	if err != nil && err != errs.ErrNotFound {
		fmt.Println("RoomModel.FindOne err")
		return &pb.PartyTerminateResp{
			Iret: errTypes.ErrCodeSysBusy,
			Smsg: errTypes.ErrSysBusy.Error(),
		}, nil
	}

	if err == errs.ErrNotFound {
		return &pb.PartyTerminateResp{
			Iret: 1,
			Smsg: errTypes.ErrPartyNotfound.Error(),
		}, nil
	}

	//2 判断房主是否是当前用户
	if roomData.CreateUser != in.Uid {
		return &pb.PartyTerminateResp{
			Iret: errTypes.ErrCodeSysBusy,
			Smsg: errTypes.ErrSysBusy.Error(),
		}, nil
	}

	//3 取roomid的最后一次加入记录，判断退出时间是否ok
	oldRoomUser, err := l.rep.RoomUserModel.FindOneByRoomId(l.ctx, in.RoomId, in.Uid)
	fmt.Printf("oldroomuser is %v \n", oldRoomUser)
	if err != nil && err != errs.ErrNotFound {
		fmt.Println("RoomModel.FindOneByRoomId err %v \n", err)
		return &pb.PartyTerminateResp{
			Iret: errTypes.ErrCodeSysBusy,
			Smsg: errTypes.ErrSysBusy.Error(),
		}, nil
	}

	if oldRoomUser != nil && (oldRoomUser.LeaveAt.Unix() > 0 || oldRoomUser.CoolAt.Unix() > 0) {
		return &pb.PartyTerminateResp{
			Iret: 1,
			Smsg: errTypes.ErrPartyLeaved.Error(),
		}, nil
	}

	//4 更新room数据
	roomData.Status = 1
	roomData.TerminateAt = time.Now()
	roomUv, _ := l.rep.Redis.Get(l.ctx, rediskey.CacheSocialImRoomUV+strconv.FormatInt(in.RoomId, 10)).Result()
	roomData.UV, _ = strconv.ParseInt(roomUv, 10, 64)

	//5 更新roomuser
	oldRoomUser.LeaveAt = time.Now()
	err = l.rep.RoomModel.Transaction(l.ctx, func(db *gorm.DB) error {
		err := l.rep.RoomModel.Update(l.ctx, db, roomData)
		if err != nil {
			return err
		}

		err = l.rep.RoomUserModel.Update(l.ctx, db, oldRoomUser)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return &pb.PartyTerminateResp{
			Iret: errTypes.ErrCodeSysBusy,
			Smsg: errTypes.ErrSysBusy.Error(),
		}, nil
	}

	//todo 清除filter
	l.rep.Redis.Del(l.ctx, rediskey.CacheSocialImRoomBloom+strconv.FormatInt(in.RoomId, 10), rediskey.CacheSocialImRoomUV+strconv.FormatInt(in.RoomId, 10)).Err()

	terminateMsg := &pb.PartyTerminateResp{
		Name:             roomData.Name,
		PartyType:        roomData.PartyType,
		CreateUser:       roomData.CreateUser,
		CreatedAt:        roomData.CreatedAt.Unix(),
		TerminateAt:      roomData.TerminateAt.Unix(),
		RoomUV:           roomData.UV,
		RoomNewFollowers: 123,
	}
	msgData, _ := json.Marshal(terminateMsg)
	//发送频道消息
	l.rep.UserRpc.SendRtmChannel(l.ctx, &userrpc.SendRtmChannelReq{
		From:        agora.ADMINUSER,
		ChannelName: roomData.Mark,
		MessageType: agora.TERMINATEROOM,
		MessageBody: string(msgData),
	})

	return terminateMsg, nil
}
