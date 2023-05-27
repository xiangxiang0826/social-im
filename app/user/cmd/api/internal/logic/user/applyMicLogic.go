package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyMicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyMicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyMicLogic {
	return &ApplyMicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyMicLogic) ApplyMic(req *types.ApplyMicReq) (resp *types.ApplyMicResp, err error) {
	// todo: add your logic here and delete this line
	var mresp types.ApplyMicResp
	rep, err := l.svcCtx.UserRpc.ApplyMic(l.ctx, &pb.ApplyMicReq{Room: req.Room, User: req.User})

	if rep != nil {
		mresp.Status.Status = rep.Status.Status
		mresp.Status.Code = int(rep.Status.Code)
		return &mresp, nil

	}
	return nil, err
}
