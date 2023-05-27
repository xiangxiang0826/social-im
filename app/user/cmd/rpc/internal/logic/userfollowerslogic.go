package logic

import (
	"context"
	"fmt"
	"social-im/app/user/cmd/rpc/internal/repository"

	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFollowersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewUserFollowersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFollowersLogic {
	return &UserFollowersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *UserFollowersLogic) UserFollowers(in *pb.RoomUser) (*pb.Users, error) {
	// todo: add your logic here and delete this line
	fmt.Println(" get user followers list----11111")
	user, err := l.rep.UserFollowerModel.FindUserFollowers(l.ctx, l.rep.Mysql, in.Room, in.User)
	fmt.Println(" get user followers list----222222")
	if user != nil {
		fmt.Println(" get user followers list----33333:", *user, "----and len:", len(*user))
		return &pb.Users{Users: *user}, nil
	}
	fmt.Println(" get user followers list----44444444")
	return nil, err

}
