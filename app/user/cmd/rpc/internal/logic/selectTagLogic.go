package logic

import (
	"context"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"

	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewSelectTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectTagLogic {
	return &SelectTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *SelectTagLogic) SelectTag(in *pb.SelectTagReq) (*pb.SelectTagResp, error) {
	// todo: add your logic here and delete this line
	userBaseInfo, err := l.rep.UserBaseModel.FindOneByUid(l.ctx, in.Uid)
	if err != nil && !errs.RecordNotFound(err) {
		return nil, xerr.NewErrMsg(err.Error())
	}
	if len(in.DisplayFields) > 0 {
		userBaseInfo.DisplayFields = in.DisplayFields
	}
	if err == errs.ErrNotFound {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_DATA_EMPTY, "用户基本资料信息不存在:%d", in.Uid)
	}
	if len(in.DisplayFields) <= 0 { //用户没有填写任何信息直接返回成功
		return &pb.SelectTagResp{
			Id:  userBaseInfo.Id,
			Uid: userBaseInfo.Uid,
		}, nil
	}
	err = l.rep.UserBaseModel.Update(l.ctx, l.rep.Mysql, userBaseInfo)
	if err != nil {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "")
	}
	return &pb.SelectTagResp{Id: userBaseInfo.Id, Uid: userBaseInfo.Uid}, nil
}
