package user

import (
	"context"
	"social-im/app/user/cmd/rpc/pb"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnMicListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOnMicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnMicListLogic {
	return &OnMicListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OnMicListLogic) OnMicList(req *types.OnMicListReq) (resp *types.OnMicListResp, err error) {
	// todo: add your logic here and delete this line
	var mresp types.OnMicListResp
	users, err := l.svcCtx.UserRpc.OnMicersList(l.ctx, &pb.OnMicersListReq{Room: req.Room})

	if err != nil {
		return nil, err
	}

	if users != nil {
		for _, v := range users.Users {
			mresp.Users = append(mresp.Users, v.User)
		}
		return &mresp, nil
	}

	return nil, nil

}
