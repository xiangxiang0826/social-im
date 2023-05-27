package logic

import (
	"context"

	"social-im/app/user/cmd/rpc/internal/svc"
	cachedtoken "social-im/app/user/cmd/rpc/internal/token"
	"social-im/app/user/cmd/rpc/pb"
	errTypes "social-im/common/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccessTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	cachedtoken *cachedtoken.CacheToken
}

func NewGetAccessTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccessTokenLogic {
	return &GetAccessTokenLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		cachedtoken: cachedtoken.NewCacheToken(svcCtx),
	}
}

func (l *GetAccessTokenLogic) GetAccessToken(in *pb.GetAccessTokenReq) (*pb.GetAccessTokenResp, error) {
	// todo: add your logic here and delete this line
	accessToken, err := l.cachedtoken.GetJwtTokenWithCache(in.Uid, "0", in.Mobile, l.svcCtx.Config.JwtAuth.AccessSecret)

	if err != nil {
		return &pb.GetAccessTokenResp{
			Iret:        errTypes.ErrCodeFailed,
			Smsg:        errTypes.ErrSysBusy.Error(),
			AccessToken: "",
		}, nil
	}
	return &pb.GetAccessTokenResp{
		AccessToken: accessToken,
	}, nil
}
