package user

import (
	"context"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
	"social-im/app/user/cmd/rpc/userrpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type SmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsLogic {
	return &SmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SmsLogic) Sms(req *types.SmsReq) (*types.SmsResp, error) {
	// todo: add your logic here and delete this line
	// todo: add your logic here and delete this line
	smsResp, err := l.svcCtx.UserRpc.Sms(l.ctx, &userrpc.SmsReq{
		Mobile: req.Mobile,
		Type:   req.Type,
	})

	if err != nil {
		return nil, err
	}

	var resp types.SmsResp
	_ = copier.Copy(&resp, smsResp)
	// _ = copier.C
	return &resp, nil
}
