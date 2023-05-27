package logic

import (
	"context"
	roomPb "social-im/app/room/cmd/rpc/pb"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/model"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"
	"strconv"
	"time"

	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"

	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"social-im/common/agora"
)

type InviteMicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewInviteMicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteMicLogic {
	return &InviteMicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *InviteMicLogic) InviteMic(in *pb.InviteMicReq) (*pb.InviteMicResp, error) {
	// todo: add your logic here and delete this line
	onMicers, err := l.rep.RoomManagerOnMicer.FindUsersOnMic(l.ctx, int64(in.Room))
	if err != nil && !errs.RecordNotFound(err) { //数据库错误
		return nil, err
	}

	if len(*onMicers) >= 10 {
		return nil, xerr.NewErrMsg("麦位已满") //麦位已满
	}
	// 查询用户麦位信息， 如果已经有麦位相关信息，则返回不用操作
	userMic, err := l.rep.RoomManagerOnMicer.GetUserMicStatus(l.ctx, int64(in.Room), in.User)
	if err != nil && !errs.RecordNotFound(err) { //排除没用户记录
		return nil, err
	}

	if userMic != nil { //用户已经有麦位状态
		return nil, xerr.NewErrMsg("用户已经有麦位状态， 请勿重复操作")
	}

	err = l.rep.RoomManagerOnMicer.Insert(l.ctx, l.rep.Mysql, &model.RoomManagerOnmicer{
		RoomId:   int64(in.Room),
		Type:     1,
		Uid:      in.User,
		UidType:  3,
		JoinTime: time.Now().Unix(),
	})

	if err != nil {
		return nil, err
	}
	//需要删除影响到的redis Key 记录

	delKey := "cache:socialIm:roomId:invitedMic:" + strconv.Itoa(int(in.Room))
	l.svcCtx.Redis.Del(delKey)

	delKey = "cache:socialIm:roomId:micStatus:" + strconv.Itoa(int(in.Room)) + ":" + strconv.Itoa(int(in.User))
	l.svcCtx.Redis.Del(delKey)

	sFrom := strconv.FormatInt(in.Uid, 10)
	sUser := strconv.FormatInt(in.User, 10)

	room, err := l.svcCtx.RoomRpc.RoominfoById(l.ctx, &roomPb.RoominfoReq{Room: in.Room})
	if err != nil {
		return nil, err
	} else {
		roomCreater := strconv.FormatInt(room.CreateUser, 10)

		type uuser struct {
			UserId int64 `json:"user_id"`
		}
		var msgUser uuser
		msgUser.UserId = in.User
		msgdata, _ := json.Marshal(&msgUser)

		_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: "1", To: roomCreater, MessageType: agora.INVITEMIC, MessageBody: string(msgdata)}) //给房主发消息
		_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: "1", To: sUser, MessageType: agora.INVITEMIC, MessageBody: string(msgdata)})
		_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: "1", To: sFrom, MessageType: agora.INVITEMIC, MessageBody: string(msgdata)}) //给房主发消息

		if err != nil {
			return nil, err
		}

	}

	res := pb.MicStatus{Status: "ok", Code: 0}
	return &pb.InviteMicResp{Status: &res}, err

}
