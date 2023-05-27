package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFollowsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFollowsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFollowsLogic {
	return &UserFollowsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFollowsLogic) UserFollows(req *types.UserFollowListReq) (resp *types.UserFollowListResq, err error) {
	// todo: add your logic here and delete this line
	rep, err := l.svcCtx.UserRpc.UserFollows(l.ctx, &pb.RoomUser{Room: req.Room, User: req.Uid})

	return &types.UserFollowListResq{Users: rep.Users}, err
	return
}
