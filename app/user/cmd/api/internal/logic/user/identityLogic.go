package user

import (
	"context"
	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IdentityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIdentityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IdentityLogic {
	return &IdentityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IdentityLogic) Identity(req *types.IdentityReq) (resp *types.IdentityResp, err error) {
	// todo: add your logic here and delete this line

	return
}
