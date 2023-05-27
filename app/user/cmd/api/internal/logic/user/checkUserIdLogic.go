package user

import (
	"context"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
	"social-im/app/user/cmd/rpc/userrpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserIdLogic {
	return &CheckUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckUserIdLogic) CheckUserId(req *types.CheckUserIdReq) (*types.CheckUserIdResp, error) {
	// todo: add your logic here and delete this line
	checkUserIdResp, err := l.svcCtx.UserRpc.CheckUserId(l.ctx, &userrpc.CheckUserIdReq{
		Mobile: req.Mobile,
	})

	if err != nil {
		return nil, err
	}

	// fmt.Printf("CheckUserId resp is %v", resp)
	var resp types.CheckUserIdResp
	_ = copier.Copy(&resp, checkUserIdResp)
	return &resp, nil
}
