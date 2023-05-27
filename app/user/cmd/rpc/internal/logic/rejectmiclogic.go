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

type RejectMicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewRejectMicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RejectMicLogic {
	return &RejectMicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *RejectMicLogic) RejectMic(in *pb.RejectMicReq) (*pb.RejectMicResp, error) { //管理员房主拒绝 上麦申请
	// todo: add your logic here and delete this line

	userMic, err := l.rep.RoomManagerOnMicer.GetUserMicStatus(l.ctx, int64(in.Room), in.User)
	if err != nil && !errs.RecordNotFound(err) { //排除没用户记录
		return nil, err
	}

	if userMic != nil && userMic.UidType == 2 && userMic.RoomId == in.Room { // 用户是申请上麦的状态 ，进行上麦拒绝

		err := l.rep.RoomManagerOnMicer.DeleteUserMicStatus(l.ctx, l.rep.Mysql, in.User, in.Room)
		if err != nil {
			return nil, err
		}
		//上麦成功， 广播消息
		//需要删除影响到的redis Key 记录
		delKey := "cache:socialIm:roomId:applyingMic:" + strconv.Itoa(int(in.Room))
		l.svcCtx.Redis.Del(delKey)

		delKey = "cache:socialIm:roomId:micStatus:" + strconv.Itoa(int(in.Room)) + ":" + strconv.Itoa(int(in.User))
		l.svcCtx.Redis.Del(delKey)

		room, err := l.svcCtx.RoomRpc.RoominfoById(l.ctx, &roomPb.RoominfoReq{Room: in.Room})
		if err != nil {
			return nil, err
		}
		//
		roomCreater := strconv.FormatInt(room.CreateUser, 10)
		user := strconv.FormatInt(in.User, 10)
		CreaterOrManager := strconv.FormatInt(in.Uid, 10)

		type uuser struct {
			UserId int64 `json:"user_id"`
		}
		var msgUser uuser
		msgUser.UserId = in.User
		msgdata, _ := json.Marshal(&msgUser)

		_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: "1", To: user, MessageType: agora.REJECTMIC, MessageBody: string(msgdata)})             //给房主发消息
		_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: "1", To: roomCreater, MessageType: agora.REJECTMIC, MessageBody: string(msgdata)})      //给房主发消息
		_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: "1", To: CreaterOrManager, MessageType: agora.REJECTMIC, MessageBody: string(msgdata)}) //给房主发消息

		if err != nil {
			return nil, err
		}

		res := pb.MicStatus{Status: "ok", Code: 0}
		return &pb.RejectMicResp{Status: &res}, err

	}
	return nil, xerr.NewErrMsg("用户麦位状态冲突， 请稍后")
}
