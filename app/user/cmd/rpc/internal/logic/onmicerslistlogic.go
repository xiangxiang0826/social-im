package logic

import (
	"context"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnMicersListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewOnMicersListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnMicersListLogic {
	return &OnMicersListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *OnMicersListLogic) OnMicersList(in *pb.OnMicersListReq) (*pb.OnMicersListResp, error) {
	// todo: add your logic here and delete this line
	var mresp pb.OnMicersListResp
	users, err := l.rep.RoomManagerOnMicer.FindUsersOnMic(l.ctx, in.Room)

	if err != nil {
		return nil, err
	}

	if users != nil {
		for _, v := range *users {
			var elm pb.UserMicTime
			elm.User = v.Uid
			elm.Time = v.JoinTime
			mresp.Users = append(mresp.Users, &elm)

		}
		return &mresp, nil
	}

	return nil, nil
}
