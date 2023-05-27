package logic

import (
	"context"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/model"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"

	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBackgroundImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewUpdateBackgroundImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBackgroundImageLogic {
	return &UpdateBackgroundImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *UpdateBackgroundImageLogic) UpdateBackgroundImage(in *pb.UpdateBackgroundImageReq) (*pb.UpdateBackgroundImageResp, error) {
	// todo: add your logic here and delete this line
	userBaseInfo, err := l.rep.UserBaseModel.FindOneByUid(l.ctx, in.Uid)
	if err != nil && !errs.RecordNotFound(err) {
		return nil, xerr.NewErrMsg(err.Error())
	}
	var currentUserBaseId int64
	if errs.RecordNotFound(err) { //如果不存在进行初始化
		userBaseInfo = &model.AppUserBase{}
	} else {
		currentUserBaseId = userBaseInfo.Id
	}
	userBaseInfo.Uid = in.Uid
	userBaseInfo.BackgroundUrl = in.BackgroundUrl
	userBaseInfo.BackgroundSmallUrl = in.BackgroundSmallUrl
	if currentUserBaseId > 0 {
		err = l.rep.UserBaseModel.Update(l.ctx, l.rep.Mysql, userBaseInfo)
	} else {
		if len(in.BackgroundUrl) <= 0 || len(in.BackgroundSmallUrl) <= 0 { //用户没有填写任何信息直接返回成功
			return &pb.UpdateBackgroundImageResp{
				Id:  currentUserBaseId,
				Uid: in.Uid,
			}, nil
		}
		err = l.rep.UserBaseModel.Insert(l.ctx, l.rep.Mysql, userBaseInfo)
	}
	if err != nil {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "")
	}
	return &pb.UpdateBackgroundImageResp{
		Id:  userBaseInfo.Id,
		Uid: userBaseInfo.Uid,
	}, nil
}
