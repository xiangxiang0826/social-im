package logic

import (
	"context"
	"fmt"

	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"
	errTypes "social-im/common/types"
	"social-im/common/utils/encrypt"
	"social-im/common/xorm/errs"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewUpdatePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePwdLogic {
	return &UpdatePwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *UpdatePwdLogic) UpdatePwd(in *pb.UpdatePwdReq) (*pb.UpdatePwdResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.rep.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && !errs.RecordNotFound(err) {
		return nil, errTypes.ErrSysBusy
	}
	if errs.RecordNotFound(err) {
		return &pb.UpdatePwdResp{
			Iret: errTypes.ErrCodeFailed,
			Smsg: errTypes.ErrUserNotfound.Error(),
		}, nil
	}

	user.Password = encrypt.Md5(in.Pwd)
	err = l.rep.UserModel.Update(l.ctx, l.rep.Mysql, user)
	if err != nil {
		fmt.Printf("UpdatePwdLogic error is %v \n", err)
		return &pb.UpdatePwdResp{
			Iret: errTypes.ErrCodeFailed,
			Smsg: errTypes.ErrSysBusy.Error(),
		}, nil
	}

	return &pb.UpdatePwdResp{}, nil
}
