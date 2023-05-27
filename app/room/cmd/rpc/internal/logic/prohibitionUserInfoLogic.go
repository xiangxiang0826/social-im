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

type ProhibitionUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewProhibitionUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProhibitionUserInfoLogic {
	return &ProhibitionUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *ProhibitionUserInfoLogic) ProhibitionUserInfo(in *pb.ProhibitionGetReq) (*pb.ProhibitionGetResp, error) {
	// todo: add your logic here and delete this line
	prohibitionGetResp, err := l.rep.RoomProhibitionUsersModel.FindOneByRoomTypeRoomIdUid(l.ctx, in.RoomType, in.RoomId, in.Uid)
	if err != nil && err != errs.ErrNotFound {
		return nil, xerr.NewErrMsg(err.Error())
	}
	var (
		prohibitionId int64
		status        int64
	)
	if prohibitionGetResp != nil {
		prohibitionId = prohibitionGetResp.Id
		status = prohibitionGetResp.Status
	}
	return &pb.ProhibitionGetResp{
		Id:       prohibitionId,
		Status:   status,
		RoomId:   in.RoomId,
		RoomType: in.RoomType,
		Uid:      in.Uid,
	}, nil
}
