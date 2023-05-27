package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AcceptAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAcceptAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AcceptAdminLogic {
	return &AcceptAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AcceptAdminLogic) AcceptAdmin(req *types.AcceptAdminReq) (resp *types.AcceptAdminResp, err error) {
	// todo: add your logic here and delete this line

	var mresp types.AcceptAdminResp
	rep, err := l.svcCtx.UserRpc.AcceptAdmin(l.ctx, &pb.AcceptAdminReq{Room: req.Room, User: req.User})

	if rep != nil {
		mresp.Status.Iret = int64(rep.Resp.Code)
		mresp.Status.Smsg = rep.Resp.Status
		return &mresp, nil

	}

	return nil, err

}
