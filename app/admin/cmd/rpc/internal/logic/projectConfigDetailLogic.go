package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/admin/cmd/rpc/internal/repository"
	"social-im/app/admin/cmd/rpc/internal/svc"
	"social-im/app/admin/cmd/rpc/pb"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectConfigDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewProjectConfigDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectConfigDetailLogic {
	return &ProjectConfigDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *ProjectConfigDetailLogic) ProjectConfigDetail(in *pb.ProjectConfigDetailReq) (*pb.ProjectConfigDetailResp, error) {
	// todo: add your logic here and delete this line
	configResp, err := l.rep.ProjectConfigModel.FindOneByConfigKeyConfigType(l.ctx, in.ConfigKey, in.ConfigType)
	if err != nil && err != errs.ErrNotFound {
		return nil, xerr.NewErrMsg(err.Error())
	}
	var resp pb.ProjectConfigInfo
	if configResp != nil {
		_ = copier.Copy(&resp, configResp)
	}
	return &pb.ProjectConfigDetailResp{ProjectConfigInfo: &resp}, nil
}
