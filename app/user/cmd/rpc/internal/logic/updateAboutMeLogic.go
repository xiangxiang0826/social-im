package logic

import (
	"context"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"
	"social-im/app/user/model"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAboutMeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewUpdateAboutMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAboutMeLogic {
	return &UpdateAboutMeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *UpdateAboutMeLogic) UpdateAboutMe(in *pb.UpdateAboutMeReq) (*pb.UpdateAboutMeResp, error) {
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
	userBaseInfo.AboutMe = in.AboutMe
	if currentUserBaseId > 0 {
		err = l.rep.UserBaseModel.Update(l.ctx, l.rep.Mysql, userBaseInfo)
	} else {
		if len(in.AboutMe) <= 0 { //用户没有填写任何信息直接返回成功
			return &pb.UpdateAboutMeResp{
				Id:  currentUserBaseId,
				Uid: in.Uid,
			}, nil
		}
		err = l.rep.UserBaseModel.Insert(l.ctx, l.rep.Mysql, userBaseInfo)
	}
	if err != nil {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "")
	}
	return &pb.UpdateAboutMeResp{
		Id:  userBaseInfo.Id,
		Uid: userBaseInfo.Uid,
	}, nil
}
