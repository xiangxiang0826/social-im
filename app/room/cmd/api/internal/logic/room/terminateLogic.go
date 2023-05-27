package room

import (
	"context"

	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"
	"social-im/app/room/cmd/rpc/roomrpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type TerminateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTerminateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TerminateLogic {
	return &TerminateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TerminateLogic) Terminate(req *types.PartyTerminateReq) (*types.PartyTerminateResp, error) {
	// todo: add your logic here and delete this line
	terminateResp, err := l.svcCtx.RoomRpc.Terminate(l.ctx, &roomrpc.PartyTerminateReq{
		RoomId: req.RoomId,
		Uid:    req.Uid,
	})
	if err != nil {
		return nil, err
	}

	var resp types.PartyTerminateResp
	_ = copier.Copy(&resp, terminateResp)
	return &resp, nil
}
