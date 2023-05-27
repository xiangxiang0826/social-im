package room

import (
	"context"
	"fmt"
	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"
	"social-im/app/room/cmd/rpc/roomrpc"
	"social-im/app/user/cmd/rpc/userrpc"
	"social-im/common/agora"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req *types.PartyRemoveReq) (*types.PartyRemoveResp, error) {
	// todo: add your logic here and delete this line
	removeResp, err := l.svcCtx.RoomRpc.Remove(l.ctx, &roomrpc.PartyRemoveReq{
		RoomId: req.RoomId,
		Uid:    req.Uid,
	})
	if err != nil {
		return nil, err
	}

	var resp types.PartyRemoveResp
	_ = copier.Copy(&resp, removeResp)

	//新增发送管理员消息
	l.svcCtx.UserRpc.SendRtm(l.ctx, &userrpc.SendRtmReq{
		From:        agora.ADMINUSER,
		To:          strconv.FormatInt(req.Uid, 10),
		MessageType: agora.REMOVEROOM,
		MessageBody: fmt.Sprintf(agora.REMOVEROOMMSG, removeResp.Name),
	})
	return &resp, nil
}
