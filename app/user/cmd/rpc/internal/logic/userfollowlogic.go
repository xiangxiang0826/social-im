package logic

import (
	"context"
	"social-im/app/user/cmd/rpc/internal/repository"

	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFollowLogic {
	return &UserFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *UserFollowLogic) UserFollow(in *pb.RoomUsers) (*pb.CommonRespNew, error) {
	// todo: add your logic here and delete this line
	err := l.rep.UserFollowerModel.UpdateUserFollow(l.ctx, l.rep.Mysql, in.Room, in.Uid, in.User)

	return &pb.CommonRespNew{}, err
}
