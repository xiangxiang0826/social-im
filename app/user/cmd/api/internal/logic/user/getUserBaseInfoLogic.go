package user

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/user/cmd/rpc/userrpc"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserBaseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserBaseInfoLogic {
	return &GetUserBaseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserBaseInfoLogic) GetUserBaseInfo(req *types.GetUserBaseReq) (*types.GetUserBaseResp, error) {
	// todo: add your logic here and delete this line
	getUserBaseInfoResp, err := l.svcCtx.UserRpc.GetUserBaseInfo(l.ctx, &userrpc.GetUserBaseReq{
		Uid: req.Uid,
	})
	if err != nil {
		return nil, err
	}
	var resp types.GetUserBaseResp
	_ = copier.Copy(&resp, getUserBaseInfoResp)
	return &resp, nil
}
