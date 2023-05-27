package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResumeMicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResumeMicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResumeMicLogic {
	return &ResumeMicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResumeMicLogic) ResumeMic(req *types.ResumeMicReq) (resp *types.ResumeMicResp, err error) {
	// todo: add your logic here and delete this line
	var mresp types.ResumeMicResp
	rep, err := l.svcCtx.UserRpc.ResumeMic(l.ctx, &pb.ResumeMicReq{Room: req.Room, Uid: req.Uid, User: req.User})

	if rep != nil {
		mresp.Status.Status = rep.Status.Status
		mresp.Status.Code = int(rep.Status.Code)
		return &mresp, nil

	}
	return

}
