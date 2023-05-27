package logic

import (
	"context"

	"social-im/app/user/cmd/rpc/internal/svc"
	cachedtoken "social-im/app/user/cmd/rpc/internal/token"
	"social-im/app/user/cmd/rpc/pb"
	errTypes "social-im/common/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRtmTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	cachedtoken *cachedtoken.CacheToken
}

func NewGetRtmTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRtmTokenLogic {
	return &GetRtmTokenLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		cachedtoken: cachedtoken.NewCacheToken(svcCtx),
	}
}

func (l *GetRtmTokenLogic) GetRtmToken(in *pb.GetRtmTokenReq) (*pb.GetRtmTokenResp, error) {
	// todo: add your logic here and delete this line
	rtmToken, err := l.cachedtoken.GetRtmTokenWithCache(l.svcCtx.Config.AgoraConf.AppId, l.svcCtx.Config.AgoraConf.AppCertificate, in.Uid)

	if err != nil {
		return &pb.GetRtmTokenResp{
			Iret:     errTypes.ErrCodeFailed,
			Smsg:     errTypes.ErrSysBusy.Error(),
			RtmToken: "",
		}, nil
	}
	return &pb.GetRtmTokenResp{
		RtmToken: rtmToken,
	}, nil

	return &pb.GetRtmTokenResp{
		RtmToken: rtmToken,
	}, nil
}
