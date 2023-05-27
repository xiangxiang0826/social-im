package user

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
	"social-im/app/user/cmd/rpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	// todo: add your logic here and delete this line
	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &userrpc.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
    var resp types.RegisterResp
	_ = copier.Copy(&resp, registerResp)
	return &resp, nil
}
