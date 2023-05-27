package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFollowLogic {
	return &UserFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFollowLogic) UserFollow(req *types.UserFollowReq) (resp *types.UserFollowResp, err error) {
	// todo: add your logic here and delete this line
	rep, err := l.svcCtx.UserRpc.UserFollow(l.ctx, &pb.RoomUsers{Room: req.Room, Uid: req.Uid, User: req.User})

	return &types.UserFollowResp{types.CommonResp{Iret: rep.Iret, Smsg: rep.Smsg}}, err
}
