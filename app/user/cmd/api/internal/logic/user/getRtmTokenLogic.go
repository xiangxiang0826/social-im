package user

import (
	"context"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
	"social-im/app/user/cmd/rpc/userrpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRtmTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRtmTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRtmTokenLogic {
	return &GetRtmTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRtmTokenLogic) GetRtmToken(req *types.GetRtmTokenReq) (*types.GetRtmTokenResp, error) {
	// todo: add your logic here and delete this line
	getRtmTokenResp, err := l.svcCtx.UserRpc.GetRtmToken(l.ctx, &userrpc.GetRtmTokenReq{
		Uid: req.Uid,
	})

	if err != nil {
		return nil, err
	}

	var resp types.GetRtmTokenResp
	_ = copier.Copy(&resp, getRtmTokenResp)
	return &resp, nil
}
