package room

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/room/cmd/rpc/roomrpc"

	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PartyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPartyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartyListLogic {
	return &PartyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PartyListLogic) PartyList(req *types.PartyListReq) (*types.PartyListResp, error) {
	// todo: add your logic here and delete this line
	listResp, err := l.svcCtx.RoomRpc.PartyList(l.ctx, &roomrpc.PartyListReq{
		LastId:   req.LastId,
		PageSize: req.PageSize,
        OnlineNums: req.OnlineNums,
	})
	if err != nil {
		return nil, err
	}
	var resp types.PartyListResp
	_ = copier.Copy(&resp, listResp)
	return &resp, nil
}
