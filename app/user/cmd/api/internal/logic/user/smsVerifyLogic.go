package user

import (
	"context"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
	"social-im/app/user/cmd/rpc/userrpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type SmsVerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSmsVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsVerifyLogic {
	return &SmsVerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SmsVerifyLogic) SmsVerify(req *types.SmsVerifyReq) (*types.SmsVerifyResp, error) {
	// todo: add your logic here and delete this line
	smsVerifyResp, err := l.svcCtx.UserRpc.SmsVerify(l.ctx, &userrpc.SmsVerifyReq{
		Mobile: req.Mobile,
		Type:   req.Type,
		Code:   req.Code,
	})

	if err != nil {
		return nil, err
	}

	var resp types.SmsVerifyResp
	_ = copier.Copy(&resp, smsVerifyResp)
	// fmt.Printf("smsverify is %v \n", resp)
	return &resp, nil
}
