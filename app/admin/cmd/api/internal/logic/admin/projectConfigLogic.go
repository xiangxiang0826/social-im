package admin

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/admin/cmd/rpc/adminrpc"

	"social-im/app/admin/cmd/api/internal/svc"
	"social-im/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectConfigLogic {
	return &ProjectConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectConfigLogic) ProjectConfig(req *types.ProjectConfigDetailReq) (*types.ProjectConfigDetailResp, error) {
	// todo: add your logic here and delete this line
	projectConfigDetailResp, err := l.svcCtx.AdminRpc.ProjectConfigDetail(l.ctx, &adminrpc.ProjectConfigDetailReq{
		ConfigType: req.ConfigType,
		ConfigKey:  req.ConfigKey,
	})
	if err != nil {
		return nil, err
	}
	var resp types.ProjectConfigDetailResp
	if projectConfigDetailResp.ProjectConfigInfo != nil {
		_ = copier.Copy(&resp, projectConfigDetailResp.ProjectConfigInfo)
	}
	return &resp, nil
}
