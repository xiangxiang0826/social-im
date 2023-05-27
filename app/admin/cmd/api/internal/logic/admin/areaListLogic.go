package admin

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/admin/cmd/rpc/adminrpc"

	"social-im/app/admin/cmd/api/internal/svc"
	"social-im/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AreaListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAreaListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AreaListLogic {
	return &AreaListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AreaListLogic) AreaList(req *types.AreaListReq) (*types.AreaListResp, error) {
	backgroundResp, err := l.svcCtx.AdminRpc.AreaList(l.ctx, &adminrpc.AreaListReq{
		LastId:   req.LastId,
		PageSize: req.PageSize,
		Level: req.Level,
		Pid: req.Pid,
	})
	if err != nil {
		return nil, err
	}
	var resp types.AreaListResp
	_ = copier.Copy(&resp, backgroundResp)
	return &resp, nil
}
