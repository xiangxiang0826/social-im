package logic

import (
	"context"
	"fmt"
	"strconv"

	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/cmd/rpc/internal/sms"
	"social-im/app/user/cmd/rpc/internal/svc"
	cachedtoken "social-im/app/user/cmd/rpc/internal/token"
	"social-im/app/user/cmd/rpc/pb"
	"social-im/app/user/cmd/rpc/userrpc"
	errTypes "social-im/common/types"
	"social-im/common/xorm/errs"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type SmsLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep        *repository.Rep
	sms        *sms.Sms
	cacheToken *cachedtoken.CacheToken
}

func NewSmsLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsLoginLogic {
	return &SmsLoginLogic{
		ctx:        ctx,
		svcCtx:     svcCtx,
		Logger:     logx.WithContext(ctx),
		rep:        repository.NewRep(svcCtx),
		sms:        sms.NewSms(svcCtx),
		cacheToken: cachedtoken.NewCacheToken(svcCtx),
	}
}

func (l *SmsLoginLogic) SmsLogin(in *pb.SmsLoginReq) (*pb.SmsLoginResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.rep.UserModel.FindOneByMobile(l.ctx, in.Mobile)

	//系统异常
	if err != nil && !errs.RecordNotFound(err) {
		return &pb.SmsLoginResp{
			Iret: errTypes.ErrCodeFailed,
			Smsg: errTypes.ErrUserNotfound.Error(),
		}, nil
	}

	//校验验证码
	err = l.sms.VerifyCode(in.Mobile, in.Code, "login")
	fmt.Printf("SmsLoginLogic verify code : %d", in.Code)
	if err != nil {
		return &pb.SmsLoginResp{
			Iret: errTypes.ErrCodeFailed,
			Smsg: err.Error(),
		}, nil
	}

	accessToken, rtcToken, rtmToken := l.cacheToken.GetLoginToken(strconv.FormatInt(user.Id, 10), "0", user.Mobile)

	var respUser userrpc.User
	_ = copier.Copy(&respUser, user)
	return &pb.SmsLoginResp{
		AccessToken: accessToken,
		RtcToken:    rtcToken,
		RtmToken:    rtmToken,
		UserInfo:    &respUser,
	}, nil
}
