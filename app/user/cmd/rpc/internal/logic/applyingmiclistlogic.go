package logic

import (
	"context"
	"social-im/app/user/cmd/rpc/internal/repository"

	"github.com/zeromicro/go-zero/core/logx"
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"
)

type ApplyingMicListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewApplyingMicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyingMicListLogic {
	return &ApplyingMicListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *ApplyingMicListLogic) ApplyingMicList(in *pb.ApplyingMicListReq) (*pb.ApplyingMicListResp, error) {
	// todo: add your logic here and delete this line
	users, err := l.rep.RoomManagerOnMicer.FindUsersApplyingMic(l.ctx, in.Room)
	if err != nil {
		return nil, err
	}

	if users != nil {
		return &pb.ApplyingMicListResp{Users: *users}, nil
	}
	return nil, nil
}
