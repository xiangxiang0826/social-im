package room

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/room/cmd/rpc/roomrpc"

	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProhibitionUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProhibitionUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProhibitionUserInfoLogic {
	return &ProhibitionUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProhibitionUserInfoLogic) ProhibitionUserInfo(req *types.ProhibitionGetReq) (*types.ProhibitionGetResp, error) {
	prohibitionGetResp, err := l.svcCtx.RoomRpc.ProhibitionUserInfo(l.ctx, &roomrpc.ProhibitionGetReq{
		Uid:            req.Uid,
		RoomId:         req.RoomId,
		RoomType:       req.RoomType,
	})
	if err != nil {
		return nil, err
	}
	var resp types.ProhibitionGetResp
	_ = copier.Copy(&resp, prohibitionGetResp)
	return &resp, nil
}
