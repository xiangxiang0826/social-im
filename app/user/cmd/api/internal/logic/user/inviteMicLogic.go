package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteMicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInviteMicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteMicLogic {
	return &InviteMicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InviteMicLogic) InviteMic(req *types.InviteMicReq) (resp *types.InviteMicResp, err error) {
	// todo: add your logic here and delete this line
	var mresp types.InviteMicResp
	rep, err := l.svcCtx.UserRpc.InviteMic(l.ctx, &pb.InviteMicReq{Room: req.Room, Uid: req.Uid, User: req.User})

	if rep != nil {
		mresp.Status.Status = rep.Status.Status
		mresp.Status.Code = int(rep.Status.Code)
		return &mresp, nil

	}

	return nil, err

}
