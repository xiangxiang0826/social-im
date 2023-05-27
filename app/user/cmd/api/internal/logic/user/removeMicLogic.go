package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveMicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveMicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveMicLogic {
	return &RemoveMicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveMicLogic) RemoveMic(req *types.RemoveMicReq) (resp *types.RemoveMicResp, err error) {
	// todo: add your logic here and delete this line
	var mresp types.RemoveMicResp
	rep, err := l.svcCtx.UserRpc.RemoveMic(l.ctx, &pb.RemoveMicReq{Room: req.Room, Uid: req.Uid, User: req.User})

	if rep != nil {
		mresp.Status.Status = rep.Status.Status
		mresp.Status.Code = int(rep.Status.Code)
		return &mresp, nil

	}

	return nil, err

}
