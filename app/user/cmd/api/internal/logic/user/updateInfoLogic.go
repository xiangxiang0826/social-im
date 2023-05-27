package user

import (
	"context"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
	"social-im/app/user/cmd/rpc/userrpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInfoLogic) UpdateInfo(req *types.UpdateInfoReq) (*types.UpdateInfoResp, error) {
	// todo: add your logic here and delete this line
	updateInfoResp, err := l.svcCtx.UserRpc.UpdateInfo(l.ctx, &userrpc.UpdateInfoReq{
		Uid: req.Uid,
		Mobile:   req.Mobile,
		NickName: req.Nickname,
		Sex:      req.Sex,
		Avatar:   req.Avatar,
		Birthday: req.Birthday,
	})
	if err != nil {
		return nil, err
	}
	var resp types.UpdateInfoResp
	_ = copier.Copy(&resp, updateInfoResp)
	return &resp, nil
}
