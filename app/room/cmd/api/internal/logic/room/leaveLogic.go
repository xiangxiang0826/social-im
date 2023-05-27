package room

import (
	"context"

	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"
	"social-im/app/room/cmd/rpc/roomrpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LeaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLeaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LeaveLogic {
	return &LeaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LeaveLogic) Leave(req *types.PartyLeaveReq) (*types.PartyLeaveResp, error) {
	// todo: add your logic here and delete this line
	leaveResp, err := l.svcCtx.RoomRpc.Leave(l.ctx, &roomrpc.PartyLeaveReq{
		RoomId: req.RoomId,
		Uid:    req.Uid,
	})
	if err != nil {
		return nil, err
	}

	var resp types.PartyLeaveResp
	_ = copier.Copy(&resp, leaveResp)
	return &resp, nil
}
