package user

import (
	"context"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
	"social-im/app/user/cmd/rpc/userrpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type SmsLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSmsLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsLoginLogic {
	return &SmsLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SmsLoginLogic) SmsLogin(req *types.SmsLoginReq) (*types.SmsLoginResp, error) {
	// todo: add your logic here and delete this line

	smsLoginResp, err := l.svcCtx.UserRpc.SmsLogin(l.ctx, &userrpc.SmsLoginReq{
		Mobile: req.Mobile,
		Code:   req.Code,
	})

	if err != nil {
		return nil, err
	}

	var resp types.SmsLoginResp
	_ = copier.Copy(&resp, smsLoginResp)
	return &resp, nil
}
