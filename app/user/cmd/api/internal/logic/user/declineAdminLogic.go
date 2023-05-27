package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeclineAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeclineAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeclineAdminLogic {
	return &DeclineAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeclineAdminLogic) DeclineAdmin(req *types.DeclineAdminReq) (resp *types.DeclineAdminResp, err error) {
	// todo: add your logic here and delete this line
	var mresp types.DeclineAdminResp
	rep, err := l.svcCtx.UserRpc.DeclineAdmin(l.ctx, &pb.DeclineAdminReq{Room: req.Room, User: req.User})

	if rep != nil {
		mresp.Status.Iret = int64(rep.Resp.Code)
		mresp.Status.Smsg = rep.Resp.Status
		return &mresp, nil
	}

	return nil, err
}
