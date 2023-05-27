package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/room/cmd/rpc/internal/repository"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"

	"social-im/app/room/cmd/rpc/internal/svc"
	"social-im/app/room/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PartyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewPartyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PartyListLogic {
	return &PartyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *PartyListLogic) PartyList(in *pb.PartyListReq) (*pb.PartyListResp, error) {
	// todo: add your logic here and delete this line
	total, list, err := l.rep.PartyList(l.ctx, l.Logger, in.OnlineNums, in.LastId, in.PageSize)
	if err != nil && err != errs.ErrNotFound {
		return nil, xerr.NewErrMsg(err.Error())
	}
    var pbPartyInfoList []*pb.PartyInfo
	if len(list) > 0 {
		_ = copier.Copy(&pbPartyInfoList, list)
	}
	return &pb.PartyListResp{Total: total, List: pbPartyInfoList}, nil
}
