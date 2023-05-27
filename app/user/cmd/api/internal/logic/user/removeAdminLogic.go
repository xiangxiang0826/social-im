package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveAdminLogic {
	return &RemoveAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveAdminLogic) RemoveAdmin(req *types.RemoveAdminReq) (resp *types.RemoveAdminResp, err error) {
	// todo: add your logic here and delete this line
	var mresp types.RemoveAdminResp
	rep, err := l.svcCtx.UserRpc.RemoveAdmin(l.ctx, &pb.RemoveAdminReq{Room: req.Room, Uid: req.Uid, User: req.User})

	if rep != nil {
		mresp.Status.Smsg = rep.Resp.Status
		mresp.Status.Iret = int64(rep.Resp.Code)
		return &mresp, nil

	}

	return nil, err

}
