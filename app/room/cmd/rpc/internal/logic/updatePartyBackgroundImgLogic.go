package logic

import (
	"context"
	"social-im/app/room/cmd/rpc/internal/repository"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"

	"social-im/app/room/cmd/rpc/internal/svc"
	"social-im/app/room/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePartyBackgroundImgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewUpdatePartyBackgroundImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePartyBackgroundImgLogic {
	return &UpdatePartyBackgroundImgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *UpdatePartyBackgroundImgLogic) UpdatePartyBackgroundImg(in *pb.PartyBackGroundImgUpdateReq) (*pb.PartyBackGroundImgUpdateResp, error) {
	// todo: add your logic here and delete this line
	partyResp, err := l.rep.RoomModel.FindOneByMark(l.ctx, in.Mark)
	if err != nil && err != errs.ErrNotFound {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "")
	}
	if err == errs.ErrNotFound {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_DATA_EMPTY, "派对表:"+in.Mark+",数据为空.")
	}
	partyResp.BackgroundUrl = in.BackgroundUrl
	partyResp.BackgroundSmallUrl = in.BackgroundSmallUrl
	err = l.rep.RoomModel.Update(l.ctx, l.rep.Mysql, partyResp)
	if err != nil {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "派对表更新失败:"+in.Mark+",错误信息:")
	}
	return &pb.PartyBackGroundImgUpdateResp{Id: partyResp.Id, Mark: partyResp.Mark, BackgroundUrl: in.BackgroundUrl, BackgroundSmallUrl: in.BackgroundSmallUrl}, nil
}
