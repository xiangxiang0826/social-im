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
	"social-im/common/xerr"
	"social-im/common/xorm/errs"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep        *repository.Rep
	cacheToken *cachedtoken.CacheToken
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:        ctx,
		svcCtx:     svcCtx,
		Logger:     logx.WithContext(ctx),
		rep:        repository.NewRep(svcCtx),
		cacheToken: cachedtoken.NewCacheToken(svcCtx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// todo: add your logic here and delete this line
	var respUser userrpc.User
	user, err := l.rep.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	fmt.Printf("%#+v \n", user)
	// userId := user.Id
	if err != nil && !errs.RecordNotFound(err) {
		return nil, errTypes.ErrSysBusy
	}

	accessToken, rtcToken, rtmToken := l.cacheToken.GetLoginToken(strconv.FormatInt(user.Id, 10), "0", user.Mobile)

	if user != nil && user.Id > 0 {
		_ = copier.Copy(&respUser, user)
		return &pb.RegisterResp{
			Iret:        errTypes.ErrCodeFailed,
			Smsg:        errTypes.ErrUserRegistered.Error(),
			AccessToken: accessToken,
			RtmToken:    rtmToken,
			RtcToken:    rtcToken,
			UserInfo:    &respUser,
		}, nil
	}
	// if len(user.NickName) == 0 {
	// 	user.NickName = "tayue"
	// }
	user.Mobile = in.Mobile

	if len(in.Password) > 0 {
		user.Password = encrypt.Md5(in.Password)
	} else {
		user.Password = ""
	}

	err = l.rep.UserModel.Insert(l.ctx, l.rep.Mysql, user)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}
	l.rep.Mysql.Last(&user)

	_ = copier.Copy(&respUser, user)
	return &pb.RegisterResp{
		UserInfo:    &respUser,
		AccessToken: accessToken,
		RtmToken:    rtmToken,
		RtcToken:    rtcToken,
	}, nil
}
