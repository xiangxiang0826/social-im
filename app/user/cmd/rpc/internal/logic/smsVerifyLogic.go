package logic

import (
	"context"

	"social-im/app/user/cmd/rpc/internal/sms"
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"
	errTypes "social-im/common/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SmsVerifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	sms *sms.Sms
}

func NewSmsVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsVerifyLogic {
	return &SmsVerifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		sms:    sms.NewSms(svcCtx),
	}
}

func (l *SmsVerifyLogic) SmsVerify(in *pb.SmsVerifyReq) (*pb.SmsVerifyResp, error) {
	// todo: add your logic here and delete this line
	err := l.sms.VerifyCode(in.Mobile, in.Code, in.Type)
	if err != nil {
		return &pb.SmsVerifyResp{
			Iret: errTypes.ErrCodeFailed,
			Smsg: errTypes.ErrCodeNotMatch.Error(),
			Type: in.Type,
		}, nil
	}

	return &pb.SmsVerifyResp{
		Type: in.Type,
	}, nil
}
