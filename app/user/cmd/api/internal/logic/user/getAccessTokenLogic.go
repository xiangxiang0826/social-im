package user

import (
	"context"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
	"social-im/app/user/cmd/rpc/userrpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccessTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAccessTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccessTokenLogic {
	return &GetAccessTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAccessTokenLogic) GetAccessToken(req *types.GetAccessTokenReq) (*types.GetAccessTokenResp, error) {
	// todo: add your logic here and delete this line
	getAccessTokenResp, err := l.svcCtx.UserRpc.GetAccessToken(l.ctx, &userrpc.GetAccessTokenReq{
		Uid: req.Uid,
	})

	if err != nil {
		return nil, err
	}

	var resp types.GetAccessTokenResp
	_ = copier.Copy(&resp, getAccessTokenResp)
	return &resp, nil
}
