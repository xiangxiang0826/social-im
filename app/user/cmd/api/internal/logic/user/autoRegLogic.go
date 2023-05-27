package user

import (
	"context"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
	"social-im/app/user/cmd/rpc/userrpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AutoRegLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAutoRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoRegLogic {
	return &AutoRegLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AutoRegLogic) AutoReg(req *types.AutoRegReq) (*types.AutoRegResp, error) {
	// todo: add your logic here and delete this line
	autoRegResp, err := l.svcCtx.UserRpc.AutoReg(l.ctx, &userrpc.AutoRegReq{
		Mobile: req.Mobile,
	})

	if err != nil {
		return nil, err
	}

	var resp types.AutoRegResp
	// fmt.Printf("AutoReg resp is %v", resp)
	_ = copier.Copy(&resp, autoRegResp)
	return &resp, nil
}
