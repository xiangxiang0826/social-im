package room

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/room/cmd/rpc/roomrpc"

	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProhibitionUserRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProhibitionUserRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProhibitionUserRemoveLogic {
	return &ProhibitionUserRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProhibitionUserRemoveLogic) ProhibitionUserRemove(req *types.ProhibitionRemoveReq) (*types.ProhibitionRemoveResp, error) {
	// todo: add your logic here and delete this line
	removeResp, err := l.svcCtx.RoomRpc.ProhibitionUserRemove(l.ctx, &roomrpc.ProhibitionRemoveReq{
		Id:            req.Id,
		RoomId:         req.RoomId,
		RoomType:       req.RoomType,
	})
	if err != nil {
		return nil, err
	}
	var resp types.ProhibitionRemoveResp
	_ = copier.Copy(&resp, removeResp)
	return &resp, nil
}
