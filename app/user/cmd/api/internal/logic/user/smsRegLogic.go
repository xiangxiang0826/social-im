package user

import (
	"context"
	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
	"social-im/app/user/cmd/rpc/userrpc"

	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type SmsRegLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSmsRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsRegLogic {
	return &SmsRegLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SmsRegLogic) SmsReg(req *types.SmsRegReq) (*types.SmsRegResp, error) {
	smsRegResp, err := l.svcCtx.UserRpc.SmsReg(l.ctx, &userrpc.SmsRegReq{
		Mobile: req.Mobile,
		Code:   req.Code,
	})
	if err != nil {
		return nil, err
	}
	var resp types.SmsRegResp
	_ = copier.Copy(&resp, smsRegResp)
	return &resp, nil
}
