package room

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/room/cmd/rpc/roomrpc"

	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePartyBackgroundImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePartyBackgroundImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePartyBackgroundImgLogic {
	return &UpdatePartyBackgroundImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePartyBackgroundImgLogic) UpdatePartyBackgroundImg(req *types.PartyBackGroundImgUpdateReq) (*types.PartyBackGroundImgUpdateResp, error) {
	// todo: add your logic here and delete this line
	reportResp, err := l.svcCtx.RoomRpc.UpdatePartyBackgroundImg(l.ctx, &roomrpc.PartyBackGroundImgUpdateReq{
		Mark: req.Mark,
		Uid:  req.Uid,
		BackgroundUrl: req.BackgroundUrl,
		BackgroundSmallUrl: req.BackgroundSmallUrl,
	})
	if err != nil {
		return nil, err
	}
	var resp types.PartyBackGroundImgUpdateResp
	_ = copier.Copy(&resp, reportResp)
	return &resp, nil
}
