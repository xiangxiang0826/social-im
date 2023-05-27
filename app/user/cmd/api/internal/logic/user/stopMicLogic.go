package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StopMicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStopMicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StopMicLogic {
	return &StopMicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StopMicLogic) StopMic(req *types.StopMicReq) (resp *types.StopMicResp, err error) {
	// todo: add your logic here and delete this line
	var mresp types.StopMicResp
	rep, err := l.svcCtx.UserRpc.StopMic(l.ctx, &pb.StopMicReq{Room: req.Room, Uid: req.Uid, User: req.User})

	if rep != nil {
		mresp.Status.Status = rep.Status.Status
		mresp.Status.Code = int(rep.Status.Code)
		return &mresp, nil

	}
	return
}
