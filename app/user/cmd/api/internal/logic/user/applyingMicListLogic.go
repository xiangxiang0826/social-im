package user

import (
	"context"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"social-im/app/user/cmd/rpc/pb"
)

type ApplyingMicListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyingMicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyingMicListLogic {
	return &ApplyingMicListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyingMicListLogic) ApplyingMicList(req *types.ApplyingMicListReq) (resp *types.ApplyingMicListResp, err error) {
	// todo: add your logic here and delete this line
	var mresp types.ApplyingMicListResp
	users, err := l.svcCtx.UserRpc.ApplyingMicList(l.ctx, &pb.ApplyingMicListReq{Room: req.Room})

	if users != nil {
		mresp.Users = users.Users
		return &mresp, nil
	}

	return nil, err
}
