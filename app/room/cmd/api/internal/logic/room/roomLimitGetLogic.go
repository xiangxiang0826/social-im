package room

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/room/cmd/rpc/roomrpc"

	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoomLimitGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoomLimitGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoomLimitGetLogic {
	return &RoomLimitGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoomLimitGetLogic) RoomLimitGet(req *types.RoomLimitReq) (*types.RoomLimitResp, error) {
	// todo: add your logic here and delete this line
	roomLimitGetResp, err := l.svcCtx.RoomRpc.RoomLimitGet(l.ctx, &roomrpc.RoomLimitReq{
		Uid:    req.Uid,
		Type:   req.Type,
		DayNum: req.DayNum,
	})
	if err != nil {
		return nil, err
	}
	var resp types.RoomLimitResp
	_ = copier.Copy(&resp, roomLimitGetResp)
	return &resp, nil
}
