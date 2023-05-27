package gift

import (
	"context"
	"strconv"

	"social-im/app/gift/cmd/api/internal/svc"
	"social-im/app/gift/cmd/api/internal/types"
	"social-im/app/gift/cmd/rpc/giftrpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendLogic) Send(req *types.GiftSendReq) (*types.GiftSendResp, error) {
	// todo: add your logic here and delete this line
	Uid, _ := strconv.ParseInt(l.ctx.Value("uid").(string), 10, 64)
	sendResp, err := l.svcCtx.GiftRpc.Send(l.ctx, &giftrpc.GiftSendReq{
		Uid:      Uid,
		GiftId:   req.GiftId,
		GiftNum:  req.GiftNum,
		SendTo:   req.SendTo,
		RoomMark: req.RoomMark,
	})
	if err != nil {
		return nil, err
	}

	var resp types.GiftSendResp
	_ = copier.Copy(&resp, sendResp)
	return &resp, nil
}
