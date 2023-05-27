package room

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/room/cmd/rpc/roomrpc"

	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProhibitionUserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProhibitionUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProhibitionUserAddLogic {
	return &ProhibitionUserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProhibitionUserAddLogic) ProhibitionUserAdd(req *types.ProhibitionCreateReq) (*types.ProhibitionCreateResp, error) {
	// todo: add your logic here and delete this line
	createResp, err := l.svcCtx.RoomRpc.ProhibitionUserAdd(l.ctx, &roomrpc.ProhibitionCreateReq{
		Uid:            req.Uid,
		ProhibitionUid: req.ProhibitionUid,
		RoomId:         req.RoomId,
		RoomType:       req.RoomType,
	})
	if err != nil {
		return nil, err
	}
	var resp types.ProhibitionCreateResp
	_ = copier.Copy(&resp, createResp)
	return &resp, nil
}
