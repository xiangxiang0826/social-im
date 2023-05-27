package logic

import (
	"context"
	"social-im/app/user/cmd/rpc/internal/repository"

	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFollowsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewUserFollowsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFollowsLogic {
	return &UserFollowsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *UserFollowsLogic) UserFollows(in *pb.RoomUser) (*pb.Users, error) {
	// todo: add your logic here and delete this line
	user, err := l.rep.UserFollowerModel.FindUserFollows(l.ctx, l.rep.Mysql, in.Room, in.User)

	return &pb.Users{Users: *user}, err

}
