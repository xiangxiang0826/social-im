package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInviteAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteAdminLogic {
	return &InviteAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InviteAdminLogic) InviteAdmin(req *types.InviteAdminReq) (resp *types.InviteAdminResp, err error) {
	// todo: add your logic here and delete this line
	var mresp types.InviteAdminResp
	rep, err := l.svcCtx.UserRpc.InviteAdmin(l.ctx, &pb.InviteAdminReq{Room: req.Room, Uid: req.Uid, User: req.User})

	if rep != nil {
		mresp.Status.Iret = int64(rep.Resp.Code)
		mresp.Status.Smsg = rep.Resp.Status
		return &mresp, nil

	}

	return nil, err

}
