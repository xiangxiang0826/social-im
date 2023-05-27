package user

import (
	"context"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
	"social-im/app/user/cmd/rpc/userrpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRtcTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRtcTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRtcTokenLogic {
	return &GetRtcTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRtcTokenLogic) GetRtcToken(req *types.GetRtcTokenReq) (*types.GetRtcTokenResp, error) {
	// todo: add your logic here and delete this line
	getRtcTokenResp, err := l.svcCtx.UserRpc.GetRtcToken(l.ctx, &userrpc.GetRtcTokenReq{
		Uid:         req.Uid,
		ChannelName: req.ChannelName,
	})

	if err != nil {
		return nil, err
	}

	var resp types.GetRtcTokenResp
	_ = copier.Copy(&resp, getRtcTokenResp)
	return &resp, nil
}
