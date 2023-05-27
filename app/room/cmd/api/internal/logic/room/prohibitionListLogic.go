package room

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/room/cmd/rpc/roomrpc"

	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProhibitionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProhibitionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProhibitionListLogic {
	return &ProhibitionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProhibitionListLogic) ProhibitionList(req *types.ProhibitionListReq) (*types.ProhibitionListResp, error) {
	// todo: add your logic here and delete this line
	listResp, err := l.svcCtx.RoomRpc.ProhibitionList(l.ctx, &roomrpc.ProhibitionListReq{
		LastId:   req.LastId,
		PageSize: req.PageSize,
		RoomId: req.RoomId,
		RoomType: req.RoomType,
	})
	if err != nil {
		return nil, err
	}
	var resp types.ProhibitionListResp
	_ = copier.Copy(&resp, listResp)
	return &resp, nil
}
