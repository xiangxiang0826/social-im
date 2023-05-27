package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"
)

type RemoveUserRoomStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewRemoveUserRoomStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserRoomStatusLogic {
	return &RemoveUserRoomStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *RemoveUserRoomStatusLogic) RemoveUserRoomStatus(in *pb.RemoveUserRoomStatusReq) (*pb.RemoveUserRoomStatusResq, error) {
	// todo: add your logic here and delete this line
	err := l.rep.RoomManagerOnMicer.DeleteUserRoomStatus(l.ctx, l.rep.Mysql, in.User, int64(in.Room))
	if err != nil { //排除没用户记录
		return nil, err
	}
	return &pb.RemoveUserRoomStatusResq{}, nil
}
