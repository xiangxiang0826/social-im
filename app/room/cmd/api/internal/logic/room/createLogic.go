package room

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"
	"social-im/app/room/cmd/rpc/roomrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.PartyCreateReq) (*types.PartyCreateResp, error) {
	// todo: add your logic here and delete this line
	createResp, err := l.svcCtx.RoomRpc.Create(l.ctx, &roomrpc.PartyCreateReq{
		Uid:           req.Uid,
		Name:          req.Name,
		BackgroundUrl: req.BackgroundUrl,
		PartyType:     req.PartyType,
		BackgroundSmallUrl: req.BackgroundSmallUrl,
	})
	if err != nil {
		return nil, err
	}
	var resp types.PartyCreateResp
	_ = copier.Copy(&resp, createResp)
	return &resp, nil
}
