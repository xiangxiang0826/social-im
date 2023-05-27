package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AcceptMicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAcceptMicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AcceptMicLogic {
	return &AcceptMicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AcceptMicLogic) AcceptMic(req *types.AcceptMicReq) (resp *types.AcceptMicResp, err error) {
	// todo: add your logic here and delete this line
	var mresp types.AcceptMicResp
	rep, err := l.svcCtx.UserRpc.AcceptMic(l.ctx, &pb.AcceptMicReq{Room: req.Room, User: req.User})

	if rep != nil {
		mresp.Status.Status = rep.Status.Status
		mresp.Status.Code = int(rep.Status.Code)
		return &mresp, nil

	}

	return nil, err

}
