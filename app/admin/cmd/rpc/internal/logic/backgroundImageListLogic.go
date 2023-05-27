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

type BackgroundImageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewBackgroundImageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BackgroundImageListLogic {
	return &BackgroundImageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *BackgroundImageListLogic) BackgroundImageList(in *pb.BackgroundImgReq) (*pb.BackgroundImgResp, error) {
	// todo: add your logic here and delete this line
	total, list, err :=l.rep.BackgroundImageModel.BackGroundImageList(l.ctx, in.LastId, in.PageSize, in.Type)
	if err != nil && err != errs.ErrNotFound {
		return nil, xerr.NewErrMsg(err.Error())
	}
	var resp []*pb.BackImgConf
	if len(list) > 0 {
		for _, backGroundImageConf := range list {
			var pbBackImgConf pb.BackImgConf
			_ = copier.Copy(&pbBackImgConf, backGroundImageConf)
			pbBackImgConf.Name = backGroundImageConf.Name
			pbBackImgConf.Url = backGroundImageConf.Url
			pbBackImgConf.SmallUrl = backGroundImageConf.SmallUrl
			pbBackImgConf.Key = backGroundImageConf.Key
			pbBackImgConf.Tag = backGroundImageConf.Tag
			pbBackImgConf.Id = backGroundImageConf.Id
			pbBackImgConf.Type = backGroundImageConf.Type
			resp = append(resp, &pbBackImgConf)
		}
	}
	return &pb.BackgroundImgResp{Total: total, List: resp}, nil
}
