package logic

import (
	"context"
	"fmt"
	"time"

	"social-im/app/room/cmd/rpc/internal/repository"
	"social-im/app/room/cmd/rpc/internal/svc"
	"social-im/app/room/cmd/rpc/pb"
	errTypes "social-im/common/types"
	"social-im/common/xorm/errs"

	"github.com/zeromicro/go-zero/core/logx"
)

type LeaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewLeaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LeaveLogic {
	return &LeaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *LeaveLogic) Leave(in *pb.PartyLeaveReq) (*pb.PartyLeaveResp, error) {
	// todo: add your logic here and delete this line
	//1 判断roomid是否存在
	//2 取roomid的最后一次加入记录，判断退出时间是否为空看是否能正常退出房间
	//3 写入退出时间

	//1 判断roomid是否存在
	_, err := l.rep.RoomModel.FindOne(l.ctx, in.RoomId) //判断该用户当前是否已经有在开启的派对房间
	if err != nil && err != errs.ErrNotFound {
		fmt.Println("RoomModel.FindOne err")
		return &pb.PartyLeaveResp{
			Iret: errTypes.ErrCodeSysBusy,
			Smsg: errTypes.ErrSysBusy.Error(),
		}, nil
	}

	if err == errs.ErrNotFound {
		return &pb.PartyLeaveResp{
			Iret: 1,
			Smsg: errTypes.ErrPartyNotfound.Error(),
		}, nil
	}

	//2 取roomid的最后一次加入记录，判断退出时间是否ok
	oldRoomUser, err := l.rep.RoomUserModel.FindOneByRoomId(l.ctx, in.RoomId, in.Uid)
	fmt.Printf("oldroomuser is %v \n", oldRoomUser)
	if err != nil && err != errs.ErrNotFound {
		fmt.Println("RoomModel.FindOneByRoomId err %v \n", err)
		return &pb.PartyLeaveResp{
			Iret: errTypes.ErrCodeSysBusy,
			Smsg: errTypes.ErrSysBusy.Error(),
		}, nil
	}

	if oldRoomUser != nil && (oldRoomUser.LeaveAt.Unix() > 0 || oldRoomUser.CoolAt.Unix() > 0) {
		return &pb.PartyLeaveResp{
			Iret: 1,
			Smsg: errTypes.ErrPartyLeaved.Error(),
		}, nil
	}

	//更新退出时间
	oldRoomUser.LeaveAt = time.Now()
	// oldRoomUser.CoolAt = time.Now().Add(time.Duration(coolMinute) * time.Minute)
	err = l.rep.RoomUserModel.Update(l.ctx, l.rep.Mysql, oldRoomUser)
	if err != nil {
		return &pb.PartyLeaveResp{
			Iret: errTypes.ErrCodeSysBusy,
			Smsg: errTypes.ErrSysBusy.Error(),
		}, nil
	}

	return &pb.PartyLeaveResp{}, nil
}
