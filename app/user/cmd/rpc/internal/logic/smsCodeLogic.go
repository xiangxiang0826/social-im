package logic

import (
	"context"
	"errors"
	"fmt"
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"
	"social-im/common/rediskey"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
)

type SmsCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSmsCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsCodeLogic {
	return &SmsCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SmsCodeLogic) SmsCode(in *pb.SmsCodeReq) (*pb.SmsCodeResp, error) {
	// todo: add your logic here and delete this line
	redisConn := l.svcCtx.Redis
	verifyKey := rediskey.CacheSocialImVerifyCodePrefix + in.Mobile
	codeInCache, err := redisConn.GetCtx(l.ctx, verifyKey)
	fmt.Printf("error is %+v", err)

	if err != redis.Nil && err != nil {
		return &pb.SmsCodeResp{
			Msg: "sys busy",
		}, errors.New("sys busy")
	}

	fmt.Printf("cache is %+v, in is %+v \n", codeInCache, in.Code)
	if codeInCache != in.Code {
		return &pb.SmsCodeResp{
			Msg: "wrong code",
		}, errors.New("wrong code")
	}
	return &pb.SmsCodeResp{
		Msg: "ok",
	}, nil
}
