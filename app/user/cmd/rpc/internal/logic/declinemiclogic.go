package logic

import (
	"context"
	roomPb "social-im/app/room/cmd/rpc/pb"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"
	"strconv"

	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"

	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"social-im/common/agora"
)

type DeclineMicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewDeclineMicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeclineMicLogic {
	return &DeclineMicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *DeclineMicLogic) DeclineMic(in *pb.DeclineMicReq) (*pb.DeclineMicResq, error) { //用户拒绝上麦邀请
	// todo: add your logic here and delete this line

	// 查询用户麦位信息， 如果已经有麦位相关信息，则返回不用操作
	userMic, err := l.rep.RoomManagerOnMicer.GetUserMicStatus(l.ctx, int64(in.Room), in.User)
	if err != nil && !errs.RecordNotFound(err) { //排除没用户记录
		return nil, err
	}

	if userMic != nil && userMic.UidType == 3 && userMic.RoomId == in.Room { //删除用户麦位状态

		err := l.rep.RoomManagerOnMicer.DeleteUserMicStatus(l.ctx, l.rep.Mysql, in.User, in.Room)
		//err = l.rep.RoomManagerOnMicer.UpdateMicStatus(l.ctx, l.rep.Mysql, in.User, in.Room, int64(micStatus))
		if err != nil {
			return nil, err
		}
		//上麦成功， 广播消息
		//需要删除影响到的redis Key 记录
		delKey := "cache:socialIm:roomId:invitedMic:" + strconv.Itoa(int(in.Room))
		l.svcCtx.Redis.Del(delKey)

		delKey = "cache:socialIm:roomId:micStatus:" + strconv.Itoa(int(in.Room)) + ":" + strconv.Itoa(int(in.User))
		l.svcCtx.Redis.Del(delKey)

		room, err := l.svcCtx.RoomRpc.RoominfoById(l.ctx, &roomPb.RoominfoReq{Room: in.Room})
		if err != nil {
			return nil, err
		}
		roomCreater := strconv.FormatInt(room.CreateUser, 10)
		user := strconv.FormatInt(in.User, 10)

		type uuser struct {
			UserId int64 `json:"user_id"`
		}
		var msgUser uuser
		msgUser.UserId = in.User
		msgdata, _ := json.Marshal(&msgUser)

		_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: "1", To: roomCreater, MessageType: agora.DECLINEMIC, MessageBody: string(msgdata)}) //给房主发消息
		_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: "1", To: user, MessageType: agora.DECLINEMIC, MessageBody: string(msgdata)})
		//_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: CreaterOrManager, To: roomCreater, MessageType: "approve mic ok", MessageBody: "approve mic ok"}) //给房主发消息
		//_, err = UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: CreaterOrManager, To: CreaterOrManager, MessageType: "approve mic ok", MessageBody: "approve mic ok"}) //给房主发消息
		if err != nil {
			return nil, err
		}

		res := pb.MicStatus{Status: "ok", Code: 0}
		return &pb.DeclineMicResq{Status: &res}, err
	}
	return nil, xerr.NewErrMsg("用户麦位状态冲突， 请稍后")
}
