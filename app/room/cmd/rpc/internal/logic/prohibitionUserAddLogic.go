package logic

import (
	"context"
	"encoding/json"
	"social-im/app/room/cmd/rpc/internal/repository"
	"social-im/app/room/model"
	"social-im/app/user/cmd/rpc/userrpc"
	"social-im/common/agora"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"
	"time"

	"social-im/app/room/cmd/rpc/internal/svc"
	"social-im/app/room/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProhibitionUserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewProhibitionUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProhibitionUserAddLogic {
	return &ProhibitionUserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *ProhibitionUserAddLogic) ProhibitionUserAdd(in *pb.ProhibitionCreateReq) (*pb.ProhibitionCreateResp, error) {
	roomProhibitionUsersModel := &model.AppRoomProhibitionUsers{}
	roomProhibitionUsersModel.RoomId = in.RoomId
	roomProhibitionUsersModel.RoomType = in.RoomType
	roomProhibitionUsersModel.Uid = in.ProhibitionUid
	roomProhibitionUsersModel.OperatorUser = in.Uid
	roomProhibitionUsersModel.Status = 1
	roomProhibitionUsersModel.UpdateTime = time.Now()
	roomProhibitionUsersModel.CreateTime = time.Now()
	err := l.rep.RoomProhibitionUsersModel.UpsertRoomProhibitionUsersStatus(l.ctx, l.rep.Mysql, roomProhibitionUsersModel)
	if err != nil {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "")
	}
	prohibitionUsersResp, err := l.rep.RoomProhibitionUsersModel.FindOneByRoomTypeRoomIdUid(l.ctx, in.RoomType, in.RoomId, in.ProhibitionUid)
	if err != nil && err != errs.ErrNotFound {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "")
	}
	if err == errs.ErrNotFound {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_DATA_EMPTY, "房间禁言用户数据数据为空.")
	}
	roomGetResp, err := l.rep.RoomModel.FindOne(l.ctx, in.RoomId)
	if err != nil && err != errs.ErrNotFound {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "")
	}
	if err == errs.ErrNotFound {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_DATA_EMPTY, "房间派对表%d数据为空.", in.RoomId)
	}
	var userProhibitionChannelMsg repository.UserProhibitionChannelMsg
	userProhibitionChannelMsg.UserId = in.ProhibitionUid
	userProhibitionChannelMsg.Status = 1
	userProhibitionChannelMsg.Id = prohibitionUsersResp.Id
	msgdata, _ := json.Marshal(&userProhibitionChannelMsg)
	//发送频道消息
	l.rep.UserRpc.SendRtmChannel(l.ctx, &userrpc.SendRtmChannelReq{
		From:        agora.ADMINUSER,
		ChannelName: roomGetResp.Mark,
		MessageType: agora.SHUTUPUSER,
		MessageBody: string(msgdata),
	})
	return &pb.ProhibitionCreateResp{
		Id:           prohibitionUsersResp.Id,
		Status:       1,
		RoomId:       in.RoomId,
		RoomType:     in.RoomType,
		Uid:          in.ProhibitionUid,
		OperatorUser: in.Uid,
	}, nil
}
