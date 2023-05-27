package logic

import (
	"context"
	"social-im/app/room/cmd/rpc/internal/repository"
	timeUtils "social-im/common/utils/time"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"

	"social-im/app/room/cmd/rpc/internal/svc"
	"social-im/app/room/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoomLimitGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewRoomLimitGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoomLimitGetLogic {
	return &RoomLimitGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *RoomLimitGetLogic) RoomLimitGet(in *pb.RoomLimitReq) (*pb.RoomLimitResp, error) {
	// todo: add your logic here and delete this line
	currentUid := in.Uid // 当前用户id
	currentDayNum, err := timeUtils.DateDayNum(timeUtils.Now())
	if err != nil {
		l.Logger.Error(err.Error())
		return nil, xerr.NewErrMsg(err.Error())
	}
	if in.DayNum > 0 { // 如果传入dayNum则使用传入的
		currentDayNum = in.DayNum
	}
	dayRemaiNum, err := l.rep.ValidateGetUserCreatePartyDayNum(l.ctx, currentUid, currentDayNum)
	if err != nil && err != errs.ErrNotFound {
		l.Logger.Error(err.Error())
		return nil, xerr.NewErrMsg(err.Error())
	}
	return &pb.RoomLimitResp{DayRemaiNum: dayRemaiNum}, nil
}
