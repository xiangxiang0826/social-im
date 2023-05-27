package admin

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/admin/cmd/api/internal/svc"
	"social-im/app/admin/cmd/api/internal/types"
	"social-im/app/admin/cmd/rpc/adminrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BackgroundImageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBackgroundImageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BackgroundImageListLogic {
	return &BackgroundImageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BackgroundImageListLogic) BackgroundImageList(req *types.BackgroundImgReq) (*types.BackgroundImgResp, error) {
	// todo: add your logic here and delete this line
	backgroundResp, err := l.svcCtx.AdminRpc.BackgroundImageList(l.ctx, &adminrpc.BackgroundImgReq{
		LastId:   req.LastId,
		PageSize: req.PageSize,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	var resp types.BackgroundImgResp
	_ = copier.Copy(&resp, backgroundResp)
	return &resp, nil
}
