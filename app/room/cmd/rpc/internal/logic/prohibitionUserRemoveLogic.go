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
	"strconv"
	"time"

	"social-im/app/room/cmd/rpc/internal/svc"
	"social-im/app/room/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProhibitionUserRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewProhibitionUserRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProhibitionUserRemoveLogic {
	return &ProhibitionUserRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *ProhibitionUserRemoveLogic) ProhibitionUserRemove(in *pb.ProhibitionRemoveReq) (*pb.ProhibitionRemoveResp, error) {
	// todo: add your logic here and delete this line
	roomProhibitionUsersData, err := l.rep.RoomProhibitionUsersModel.FindOne(l.ctx, in.Id)
	if err != nil && err != errs.ErrNotFound {
		return nil, xerr.NewErrMsg(err.Error())
	}
	if err == errs.ErrNotFound {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_DATA_EMPTY, "房间禁言用户数据:"+strconv.FormatInt(in.Id, 10)+",数据为空.")
	}
	roomProhibitionUsersModel := &model.AppRoomProhibitionUsers{}
	roomProhibitionUsersModel.RoomId = in.RoomId
	roomProhibitionUsersModel.RoomType = in.RoomType
	roomProhibitionUsersModel.Uid = roomProhibitionUsersData.Uid
	roomProhibitionUsersModel.OperatorUser = roomProhibitionUsersData.OperatorUser
	roomProhibitionUsersModel.Status = 0
	roomProhibitionUsersModel.UpdateTime = time.Now()
	err = l.rep.RoomProhibitionUsersModel.UpsertRoomProhibitionUsersStatus(l.ctx, l.rep.Mysql, roomProhibitionUsersModel)
	if err != nil {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "")
	}
	roomGetResp, err :=l.rep.RoomModel.FindOne(l.ctx, in.RoomId)
	if err != nil && err != errs.ErrNotFound {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "")
	}
	if err == errs.ErrNotFound {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_DATA_EMPTY, "房间派对表%d数据为空.", in.RoomId)
	}
	var userProhibitionChannelMsg repository.UserProhibitionChannelMsg
	userProhibitionChannelMsg.UserId = roomProhibitionUsersModel.Uid
	userProhibitionChannelMsg.Status = 0
	userProhibitionChannelMsg.Id = roomProhibitionUsersData.Id
	msgdata, _ := json.Marshal(&userProhibitionChannelMsg)
	//发送频道消息
	l.rep.UserRpc.SendRtmChannel(l.ctx, &userrpc.SendRtmChannelReq{
		From:        agora.ADMINUSER,
		ChannelName: roomGetResp.Mark,
		MessageType: agora.UNSHUTUPUSER,
		MessageBody: string(msgdata),
	})
	return &pb.ProhibitionRemoveResp{
		Id:           roomProhibitionUsersData.Id,
		Status:       roomProhibitionUsersModel.Status,
		RoomId:       roomProhibitionUsersModel.RoomId,
		RoomType:     roomProhibitionUsersModel.RoomType,
		Uid:          roomProhibitionUsersModel.Uid,
		OperatorUser: roomProhibitionUsersModel.OperatorUser,
	}, nil
}
