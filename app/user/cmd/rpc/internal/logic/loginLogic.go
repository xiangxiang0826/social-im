package logic

import (
	"context"
	"fmt"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/cmd/rpc/internal/svc"
	cachedtoken "social-im/app/user/cmd/rpc/internal/token"
	"social-im/app/user/cmd/rpc/pb"
	"social-im/app/user/cmd/rpc/userrpc"
	errTypes "social-im/common/types"
	"social-im/common/utils/encrypt"
	"social-im/common/xorm/errs"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep        *repository.Rep
	cacheToken *cachedtoken.CacheToken
}

var tokenExpireHour = 24 * 7
var tokenExpireMinute = tokenExpireHour * 60

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:        ctx,
		svcCtx:     svcCtx,
		Logger:     logx.WithContext(ctx),
		rep:        repository.NewRep(svcCtx),
		cacheToken: cachedtoken.NewCacheToken(svcCtx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line
	var respUser userrpc.User
	user, err := l.rep.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	fmt.Printf("%#+v \n", user)
	// userId := user.Id
	if err != nil && !errs.RecordNotFound(err) {
		return &pb.LoginResp{
			Iret: errTypes.ErrCodeFailed,
			Smsg: errTypes.ErrUserNotfound.Error(),
		}, nil
	}

	if user.Password != encrypt.Md5(in.Password) {
		fmt.Println("输入密码是:", encrypt.Md5(in.Password))
		return &pb.LoginResp{
			Iret: errTypes.ErrCodeFailed,
			Smsg: errTypes.ErrPwdNotMatch.Error(),
		}, nil

	}

	accessToken, rtcToken, rtmToken := l.cacheToken.GetLoginToken(strconv.FormatInt(user.Id, 10), "0", user.Mobile)

	_ = copier.Copy(&respUser, user)
	return &pb.LoginResp{
		AccessToken: accessToken,
		RtmToken:    rtmToken,
		RtcToken:    rtcToken,
		UserInfo:    &respUser,
	}, nil
}
