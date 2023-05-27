package logic

import (
	"context"
	"fmt"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/cmd/rpc/internal/sms"
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"
	errTypes "social-im/common/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
	sms *sms.Sms
}

func NewSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsLogic {
	return &SmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
		sms:    sms.NewSms(svcCtx),
	}
}

func (l *SmsLogic) Sms(in *pb.SmsReq) (*pb.SmsResp, error) {
	verifyCode, err := l.sms.SendSms(l.ctx, in.Mobile, in.Type)
	if err != nil {
		fmt.Println("短信发送频率限制")
		return &pb.SmsResp{
			Iret: errTypes.ErrCodeFailed,
			Smsg: err.Error(),
		}, nil
	}

	return &pb.SmsResp{
		Iret: 0,
		Smsg: "",
		Msg:  verifyCode,
	}, nil
}
