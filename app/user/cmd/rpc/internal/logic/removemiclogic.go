package logic

import (
	"context"
	roomPb "social-im/app/room/cmd/rpc/pb"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"
	"strconv"

	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"social-im/common/agora"
)

type RemoveMicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewRemoveMicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveMicLogic {
	return &RemoveMicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *RemoveMicLogic) RemoveMic(in *pb.RemoveMicReq) (*pb.RemoveMicResp, error) { //删除 micstatus =1 的记录
	// todo: add your logic here and delete this line

	userMic, err := l.rep.RoomManagerOnMicer.GetUserMicStatus(l.ctx, int64(in.Room), in.User)
	if err != nil && !errs.RecordNotFound(err) { //排除没用户记录
		return nil, err
	}

	if (userMic != nil && userMic.UidType != 1) || (errs.RecordNotFound(err)) { //用户不是上麦的状态  或者没有麦位状态
		return nil, xerr.NewErrMsg("麦位状态有误， 无法进行下麦操作")
	}

	if userMic != nil && userMic.UidType == 1 && userMic.RoomId == in.Room {
		err = l.rep.RoomManagerOnMicer.DeleteUserMicStatus(l.ctx, l.rep.Mysql, in.User, in.Room)

		if err != nil {
			return nil, err
		}
		delKey := "cache:socialIm:roomId:micStatus:" + strconv.Itoa(int(in.Room)) + ":" + strconv.Itoa(int(in.User))
		l.svcCtx.Redis.Del(delKey)

		delKey = "cache:socialIm:roomId:onMic:" + strconv.Itoa(int(in.Room))
		l.svcCtx.Redis.Del(delKey)

		room, err := l.svcCtx.RoomRpc.RoominfoById(l.ctx, &roomPb.RoominfoReq{Room: in.Room})
		if err != nil {
			return nil, err
		}

		type uuser struct {
			UserId int64 `json:"user_id"`
		}
		var msgUser uuser
		msgUser.UserId = in.User
		msgdata, _ := json.Marshal(&msgUser)

		roomCreater := strconv.FormatInt(room.CreateUser, 10)
		ssuser := strconv.FormatInt(in.User, 10)
		//CreaterOrManager := strconv.FormatInt(in.Uid, 10)
		//	_, err = l.svcCtx.UserRpc.SendRtmChannel(l.ctx, &pb.SendRtmChannelReq{From: "1", ChannelName: room.Mark, MessageType: "4", MessageBody: string(msgdata)}) //给房主发消息
		_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: "1", To: ssuser, MessageType: agora.REMOVEMIC, MessageBody: string(msgdata)}) //给房主发消息
		_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: "1", To: roomCreater, MessageType: agora.REMOVEMIC, MessageBody: string(msgdata)})

		if err != nil {
			return nil, err
		}

		res := pb.MicStatus{Status: "ok", Code: 0}
		return &pb.RemoveMicResp{Status: &res}, err

	}
	return nil, xerr.NewErrMsg("用户麦位状态冲突， 请稍后")

}
