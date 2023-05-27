package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApproveMicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApproveMicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApproveMicLogic {
	return &ApproveMicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApproveMicLogic) ApproveMic(req *types.ApproveMicReq) (resp *types.ApproveMicResp, err error) {
	// todo: add your logic here and delete this line
	var mresp types.ApproveMicResp
	rep, err := l.svcCtx.UserRpc.ApproveMic(l.ctx, &pb.ApproveMicReq{Room: req.Room, Uid: req.Uid, User: req.User})

	if rep != nil {
		mresp.Status.Status = rep.Status.Status
		mresp.Status.Code = int(rep.Status.Code)
		return &mresp, nil

	}
	return nil, err
}
