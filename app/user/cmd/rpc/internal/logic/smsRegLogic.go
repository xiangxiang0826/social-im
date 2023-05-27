package logic

import (
	"context"
	"fmt"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/cmd/rpc/internal/sms"
	"social-im/app/user/cmd/rpc/internal/svc"
	cachedtoken "social-im/app/user/cmd/rpc/internal/token"
	"social-im/app/user/cmd/rpc/pb"
	"social-im/app/user/cmd/rpc/userrpc"
	errTypes "social-im/common/types"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type SmsRegLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep        *repository.Rep
	sms        *sms.Sms
	cacheToken *cachedtoken.CacheToken
}

func NewSmsRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsRegLogic {
	return &SmsRegLogic{
		ctx:        ctx,
		svcCtx:     svcCtx,
		Logger:     logx.WithContext(ctx),
		rep:        repository.NewRep(svcCtx),
		sms:        sms.NewSms(svcCtx),
		cacheToken: cachedtoken.NewCacheToken(svcCtx),
	}
}

func (l *SmsRegLogic) SmsReg(in *pb.SmsRegReq) (*pb.SmsRegResp, error) {
	// todo: add your logic here and delete this line
	var respUser userrpc.User
	user, err := l.rep.UserModel.FindOneByMobile(l.ctx, in.Mobile)

	fmt.Printf("%#+v", user)

	if err != nil && !errs.RecordNotFound(err) {
		// return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "mobile:%s,err:%v", in.Mobile, err)
		return &pb.SmsRegResp{
			Iret: errTypes.ErrSysError,
			Smsg: err.Error(),
		}, nil
	}

	accessToken, rtcToken, rtmToken := l.cacheToken.GetLoginToken(strconv.FormatInt(user.Id, 10), "0", user.Mobile)

	if user != nil && user.Id > 0 {
		_ = copier.Copy(&respUser, user)
		return &pb.SmsRegResp{
			Iret: errTypes.ErrCodeFailed,
			// Smsg:         err.New("已注册").Error(),
			Smsg:        errTypes.ErrUserRegistered.Error(),
			AccessToken: accessToken,
			RtmToken:    rtmToken,
			RtcToken:    rtcToken,
			UserInfo:    &respUser,
		}, nil
	}

	if len(user.NickName) == 0 {
		user.NickName = ""
	}

	// if len(in.Password) == 0 {
	// 	in.Password = "tayue"
	// }

	err = l.sms.VerifyCode(in.Mobile, in.Code, "reg")
	if err != nil {
		return &pb.SmsRegResp{
			Iret: errTypes.ErrCodeFailed,
			Smsg: err.Error(),
		}, nil
	}

	user.Mobile = in.Mobile
	// user.Password = encrypt.Md5(in.Password)
	err = l.rep.UserModel.Insert(l.ctx, l.rep.Mysql, user)

	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}
	l.rep.Mysql.Last(&user)

	_ = copier.Copy(&respUser, user)
	return &pb.SmsRegResp{
		AccessToken: accessToken,
		RtmToken:    rtmToken,
		RtcToken:    rtcToken,
		UserInfo:    &respUser,
	}, nil
}
