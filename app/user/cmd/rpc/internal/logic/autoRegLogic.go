package logic

import (
	"context"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/cmd/rpc/internal/svc"
	cachedtoken "social-im/app/user/cmd/rpc/internal/token"
	"social-im/app/user/cmd/rpc/pb"
	"social-im/app/user/cmd/rpc/userrpc"
	errTypes "social-im/common/types"
	"social-im/common/xorm/errs"
	"strconv"

	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type AutoRegLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep        *repository.Rep
	cacheToken *cachedtoken.CacheToken
}

func NewAutoRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoRegLogic {
	return &AutoRegLogic{
		ctx:        ctx,
		svcCtx:     svcCtx,
		Logger:     logx.WithContext(ctx),
		rep:        repository.NewRep(svcCtx),
		cacheToken: cachedtoken.NewCacheToken(svcCtx),
	}
}

func (l *AutoRegLogic) AutoReg(in *pb.AutoRegReq) (*pb.AutoRegResp, error) {
	// todo: add your logic here and delete this line
	var respUser userrpc.User
	user, err := l.rep.UserModel.FindOneByMobile(l.ctx, in.Mobile)

	// sys error
	if err != nil && !errs.RecordNotFound(err) {
		return &pb.AutoRegResp{
			Iret: errTypes.ErrSysError,
			Smsg: err.Error(),
		}, nil
	}

	// accessToken, rtcToken, rtmToken := l.cacheToken.GetLoginToken(strconv.FormatInt(user.Id, 10), "0")
	accessToken, rtcToken, rtmToken := l.cacheToken.GetLoginToken(strconv.FormatInt(user.Id, 10), "0", user.Mobile)

	// already registered
	if user != nil && user.Id > 0 {
		_ = copier.Copy(&respUser, user)
		return &pb.AutoRegResp{
			Iret:        errTypes.ErrCodeFailed,
			Smsg:        errTypes.ErrUserRegistered.Error(),
			UserInfo:    &respUser,
			AccessToken: accessToken,
			RtmToken:    rtmToken,
			RtcToken:    rtcToken,
		}, nil
	}

	//new register
	if len(user.NickName) == 0 {
		user.NickName = ""
	}

	user.Mobile = in.Mobile
	err = l.rep.UserModel.Insert(l.ctx, l.rep.Mysql, user)

	if err != nil {
		return nil, errTypes.ErrSysBusy
	}
	l.rep.Mysql.Last(&user)

	_ = copier.Copy(&respUser, user)

	return &pb.AutoRegResp{
		AccessToken: accessToken,
		RtmToken:    rtmToken,
		RtcToken:    rtcToken,
		UserInfo:    &respUser,
	}, nil
}
