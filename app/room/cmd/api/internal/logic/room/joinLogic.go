package room

import (
	"context"
	"fmt"

	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"
	"social-im/app/room/cmd/rpc/roomrpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type JoinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinLogic {
	return &JoinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinLogic) Join(req *types.PartyJoinReq) (*types.PartyJoinResp, error) {
	// todo: add your logic here and delete this line
	joinResp, err := l.svcCtx.RoomRpc.Join(l.ctx, &roomrpc.PartyJoinReq{
		RoomId: req.RoomId,
		Uid:    req.Uid,
	})
	if err != nil {
		return nil, err
	}

	var resp types.PartyJoinResp
	_ = copier.Copy(&resp, joinResp)
	fmt.Printf("api join res is %v \n", resp)
	return &resp, nil
}
