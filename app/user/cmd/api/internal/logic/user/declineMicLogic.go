package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeclineMicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeclineMicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeclineMicLogic {
	return &DeclineMicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeclineMicLogic) DeclineMic(req *types.DeclineMicReq) (resp *types.DeclineMicResp, err error) {
	// todo: add your logic here and delete this line
	var mresp types.DeclineMicResp
	rep, err := l.svcCtx.UserRpc.DeclineMic(l.ctx, &pb.DeclineMicReq{Room: req.Room, User: req.User})

	if rep != nil {
		mresp.Status.Status = rep.Status.Status
		mresp.Status.Code = int(rep.Status.Code)
		return &mresp, nil

	}

	return nil, err

}
