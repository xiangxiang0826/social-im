package user

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/user/cmd/rpc/userrpc"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBackgroundImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBackgroundImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBackgroundImageLogic {
	return &UpdateBackgroundImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBackgroundImageLogic) UpdateBackgroundImage(req *types.UpdateBackgroundImageReq) (*types.UpdateBackgroundImageResp, error) {
	// todo: add your logic here and delete this line
	updateBackgroundImageResp, err := l.svcCtx.UserRpc.UpdateBackgroundImage(l.ctx, &userrpc.UpdateBackgroundImageReq{
		Uid:     req.Uid,
		BackgroundUrl: req.BackgroundUrl,
		BackgroundSmallUrl: req.BackgroundSmallUrl,
	})
	if err != nil {
		return nil, err
	}
	var resp types.UpdateBackgroundImageResp
	_ = copier.Copy(&resp, updateBackgroundImageResp)
	return &resp, nil
}
