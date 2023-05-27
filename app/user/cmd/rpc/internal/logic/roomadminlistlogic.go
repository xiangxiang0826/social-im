package logic

import (
	"context"
	"social-im/app/user/cmd/rpc/internal/repository"

	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoomAdminListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewRoomAdminListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoomAdminListLogic {
	return &RoomAdminListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *RoomAdminListLogic) RoomAdminList(in *pb.RoomAdminListReq) (*pb.RoomAdminListResp, error) {
	// todo: add your logic here and delete this line
	var mresp pb.RoomAdminListResp
	users, err := l.rep.RoomManagerOnMicer.FindUsersIsAdmin(l.ctx, in.Room)

	if err != nil {
		return nil, err
	}

	if users != nil {
		for _, v := range *users {
			mresp.Uids = append(mresp.Uids, v)

		}
		return &mresp, nil
	}

	return &pb.RoomAdminListResp{}, nil
}
