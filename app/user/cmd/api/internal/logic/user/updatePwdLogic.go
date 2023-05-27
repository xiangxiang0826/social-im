package user

import (
	"context"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
	"social-im/app/user/cmd/rpc/userrpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePwdLogic {
	return &UpdatePwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePwdLogic) UpdatePwd(req *types.UpdatePwdReq) (*types.UpdatePwdResp, error) {
	// todo: add your logic here and delete this line
	updatePwdResp, err := l.svcCtx.UserRpc.UpdatePwd(l.ctx, &userrpc.UpdatePwdReq{
		Mobile: req.Mobile,
		Pwd:    req.Pwd,
	})
	if err != nil {
		return nil, err
	}

	var resp types.UpdatePwdResp
	_ = copier.Copy(&resp, updatePwdResp)
	return &resp, nil
}
