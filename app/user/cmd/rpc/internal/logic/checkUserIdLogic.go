package logic

import (
	"context"
	"fmt"

	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"
	errTypes "social-im/common/types"
	"social-im/common/xorm/errs"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewCheckUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserIdLogic {
	return &CheckUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *CheckUserIdLogic) CheckUserId(in *pb.CheckUserIdReq) (*pb.CheckUserIdResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.rep.UserModel.FindOneByMobile(l.ctx, in.Mobile)

	fmt.Printf("%#+v", user)

	if err != nil && !errs.RecordNotFound(err) {
		// return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "mobile:%s,err:%v", in.Mobile, err)
		return &pb.CheckUserIdResp{
			Iret:   errTypes.ErrSysError,
			Smsg:   err.Error(),
			UserId: 0,
		}, nil
	}

	if user != nil {
		//密码状态
		var pwdStatus int64
		if user.Password == "" {
			pwdStatus = 0
		} else {
			pwdStatus = 1
		}

		return &pb.CheckUserIdResp{
			Iret: 0,
			// Smsg:         err.New("已注册").Error(),
			Smsg:      "ok",
			UserId:    user.Id,
			PwdStatus: pwdStatus,
		}, nil
	}

	return &pb.CheckUserIdResp{
		Iret:   errTypes.ErrSysError,
		Smsg:   errTypes.ErrSysBusy.Error(),
		UserId: 0,
	}, nil
}
