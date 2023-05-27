package room

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/room/cmd/rpc/roomrpc"

	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePartyNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePartyNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePartyNameLogic {
	return &UpdatePartyNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePartyNameLogic) UpdatePartyName(req *types.PartyNameUpdateReq) (*types.PartyNameUpdateResp, error) {
	reportResp, err := l.svcCtx.RoomRpc.UpdatePartyName(l.ctx, &roomrpc.PartyNameUpdateReq{
		Mark: req.Mark,
		Uid:  req.Uid,
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	var resp types.PartyNameUpdateResp
	_ = copier.Copy(&resp, reportResp)
	return &resp, nil
}
