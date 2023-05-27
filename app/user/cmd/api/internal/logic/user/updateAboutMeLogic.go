package user

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/user/cmd/rpc/userrpc"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAboutMeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAboutMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAboutMeLogic {
	return &UpdateAboutMeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAboutMeLogic) UpdateAboutMe(req *types.UpdateAboutMeReq) (*types.UpdateAboutMeResp, error) {
	// todo: add your logic here and delete this line
	updateAboutMeResp, err := l.svcCtx.UserRpc.UpdateAboutMe(l.ctx, &userrpc.UpdateAboutMeReq{
		Uid:     req.Uid,
		AboutMe: req.AboutMe,
	})
	if err != nil {
		return nil, err
	}
	var resp types.UpdateAboutMeResp
	_ = copier.Copy(&resp, updateAboutMeResp)
	return &resp, nil
}
