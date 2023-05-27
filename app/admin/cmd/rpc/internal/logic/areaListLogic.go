package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/admin/cmd/rpc/internal/repository"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"

	"social-im/app/admin/cmd/rpc/internal/svc"
	"social-im/app/admin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AreaListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewAreaListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AreaListLogic {
	return &AreaListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *AreaListLogic) AreaList(in *pb.AreaListReq) (*pb.AreaListResp, error) {
	// todo: add your logic here and delete this line
	total, list, err :=l.rep.AreaModel.AreaList(l.ctx, in.Pid, in.Level, in.LastId, in.PageSize)
	if err != nil && err != errs.ErrNotFound {
		return nil, xerr.NewErrMsg(err.Error())
	}
	var resp []*pb.AreaInfo
	if len(list) > 0 {
		for _, areaConf := range list {
			var pbAreaConf pb.AreaInfo
			_ = copier.Copy(&pbAreaConf, areaConf)
			pbAreaConf.Name = areaConf.Name
			pbAreaConf.Id = areaConf.Id
			pbAreaConf.Pid = areaConf.Pid
			pbAreaConf.Shortname = areaConf.Shortname
			pbAreaConf.Latitude = areaConf.Latitude
			pbAreaConf.Longitude = areaConf.Longitude
			pbAreaConf.Level = areaConf.Level
			resp = append(resp, &pbAreaConf)
		}
	}
	return &pb.AreaListResp{Total: total, List: resp}, nil
}
