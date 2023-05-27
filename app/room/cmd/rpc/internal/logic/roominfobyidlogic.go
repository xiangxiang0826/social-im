package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/room/cmd/rpc/internal/repository"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"

	"social-im/app/room/cmd/rpc/internal/svc"
	"social-im/app/room/cmd/rpc/pb"

	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type RoominfoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewRoominfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoominfoByIdLogic {
	return &RoominfoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *RoominfoByIdLogic) RoominfoById(in *pb.RoominfoReq) (*pb.RoominfoResq, error) {
	// todo: add your logic here and delete this line
	fmt.Println("-----11111111111111")
	room, err := l.rep.RoomModel.FindOne(l.ctx, in.Room)

	if err != nil && err != errs.ErrNotFound {
		return nil, xerr.NewErrMsg(err.Error())
	}
	var roomInfo pb.RoominfoResq
	if room != nil {
		fmt.Println("-----22222")
		_ = copier.Copy(&roomInfo, room)
	}
	fmt.Println("----3333333:", roomInfo)
	return &roomInfo, err
}
