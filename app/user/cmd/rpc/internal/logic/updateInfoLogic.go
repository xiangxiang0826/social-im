package logic

import (
	"context"
	"time"

	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"
	"social-im/app/user/cmd/rpc/userrpc"
	errTypes "social-im/common/types"
	"social-im/common/xorm/errs"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *UpdateInfoLogic) UpdateInfo(in *pb.UpdateInfoReq) (*pb.UpdateInfoResp, error) {
	// todo: add your logic here and delete this line
	var respUser userrpc.User
	user, err := l.rep.UserModel.FindOne(l.ctx, in.Uid)
	if err != nil && !errs.RecordNotFound(err) {
		return nil, errTypes.ErrSysBusy
	}

	if errs.RecordNotFound(err) {
		return &pb.UpdateInfoResp{
			Iret: errTypes.ErrCodeFailed,
			Smsg: errTypes.ErrUserNotfound.Error(),
		}, nil
	}
	if in.Sex > 0 {
		user.Sex = in.Sex
	}
	if len(in.Avatar) > 0 {
		user.Avatar = in.Avatar
	}
	if in.Birthday > 0 {
		user.Birthday = time.Unix(in.Birthday, 0)
	}
	if len(in.NickName) > 0 {
		user.NickName = in.NickName
	}
	err = l.rep.UserModel.Update(l.ctx, l.rep.Mysql, user)
	if err != nil {
		return &pb.UpdateInfoResp{
			Iret: errTypes.ErrCodeFailed,
			Smsg: errTypes.ErrSysBusy.Error(),
		}, nil
	}
	_ = copier.Copy(&respUser, user)
	return &pb.UpdateInfoResp{
		UserInfo: &respUser,
	}, nil
}
