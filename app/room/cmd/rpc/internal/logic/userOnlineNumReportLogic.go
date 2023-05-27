package logic

import (
	"context"
	"social-im/app/room/cmd/rpc/internal/repository"
	"social-im/app/room/model"
	"social-im/common/utils/periodLimit"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"

	"social-im/app/room/cmd/rpc/internal/svc"
	"social-im/app/room/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOnlineNumReportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewUserOnlineNumReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOnlineNumReportLogic {
	return &UserOnlineNumReportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *UserOnlineNumReportLogic) UserOnlineNumReport(in *pb.UserNumReportReq) (*pb.UserNumReportResp, error) {
	// todo: add your logic here and delete this line
	isPeriodLimit, err :=periodLimit.IsPeriodLimit(l.rep.RateLimiter, "UserOnlineNumReport:" + in.Mark)
	if isPeriodLimit{ //限流了
		return nil, xerr.NewErrWithFormatMsg(xerr.PERIOD_LIMIT_ERROR, "UserOnlineNumReport 派对标识:"+in.Mark)
	}
	if err != nil{
		l.Logger.WithContext(l.ctx).Error(err.Error())
	}
	_, err =l.rep.RoomModel.FindOneByMark(l.ctx, in.Mark)
	if err != nil && err != errs.ErrNotFound {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "")
	}
	if err == errs.ErrNotFound {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_DATA_EMPTY, "派对表:"+in.Mark+",数据为空.")
	}
	roomPartyModel := &model.AppRoomMic{}
	roomPartyModel.Mark = in.Mark
	roomPartyModel.OnlineNums = in.OnlineNums
	err = l.rep.RoomModel.UpsertUserOnlineNums(l.ctx, l.rep.Mysql, roomPartyModel)
	if err != nil {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "")
	}
	return &pb.UserNumReportResp{Mark: in.Mark}, nil
}
