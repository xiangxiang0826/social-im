package logic

import (
	"context"
	"social-im/app/user/cmd/rpc/internal/repository"

	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUnFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewUserUnFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUnFollowLogic {
	return &UserUnFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *UserUnFollowLogic) UserUnFollow(in *pb.RoomUsers) (*pb.CommonRespNew, error) {
	// todo: add your logic here and delete this line
	err := l.rep.UserFollowerModel.UpdateUserunFollow(l.ctx, l.rep.Mysql, in.Room, in.Uid, in.User)

	return &pb.CommonRespNew{}, err

}
