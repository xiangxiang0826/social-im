package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoomAdminListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoomAdminListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoomAdminListLogic {
	return &RoomAdminListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoomAdminListLogic) RoomAdminList(req *types.RoomAdminListReq) (resp *types.RoomAdminListResp, err error) {
	// todo: add your logic here and delete this line

	var mresp types.RoomAdminListResp
	users, err := l.svcCtx.UserRpc.RoomAdminList(l.ctx, &pb.RoomAdminListReq{Room: req.Room})

	if err != nil {
		return nil, err
	}

	if users != nil {
		for _, v := range users.Uids {
			mresp.Uids = append(mresp.Uids, v)
		}
		return &mresp, nil
	}

	return nil, nil

	return
}
