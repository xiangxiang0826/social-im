package logic

import (
	"context"

	"social-im/app/user/cmd/rpc/internal/svc"
	cachedtoken "social-im/app/user/cmd/rpc/internal/token"
	"social-im/app/user/cmd/rpc/pb"
	errTypes "social-im/common/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRtcTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	cachedtoken *cachedtoken.CacheToken
}

func NewGetRtcTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRtcTokenLogic {
	return &GetRtcTokenLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		cachedtoken: cachedtoken.NewCacheToken(svcCtx),
	}
}

func (l *GetRtcTokenLogic) GetRtcToken(in *pb.GetRtcTokenReq) (*pb.GetRtcTokenResp, error) {
	// todo: add your logic here and delete this line
	rtcToken, err := l.cachedtoken.GetRtcTokenWithCache(l.svcCtx.Config.AgoraConf.AppId, l.svcCtx.Config.AgoraConf.AppCertificate, in.ChannelName, in.Uid)

	if err != nil {
		return &pb.GetRtcTokenResp{
			Iret:     errTypes.ErrCodeFailed,
			Smsg:     errTypes.ErrSysBusy.Error(),
			RtcToken: "",
		}, nil
	}
	return &pb.GetRtcTokenResp{
		RtcToken: rtcToken,
	}, nil
}
