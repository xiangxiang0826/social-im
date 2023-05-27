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

type AcceptMicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewAcceptMicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AcceptMicLogic {
	return &AcceptMicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *AcceptMicLogic) AcceptMic(in *pb.AcceptMicReq) (*pb.AcceptMicResq, error) {
	// todo: add your logic here and delete this line

	onMicers, err := l.rep.RoomManagerOnMicer.FindUsersOnMic(l.ctx, int64(in.Room))
	if err != nil && !errs.RecordNotFound(err) { //数据库错误
		return nil, err
	}
	if onMicers != nil && len(*onMicers) >= 10 {
		return nil, xerr.NewErrMsg("麦位已满") //麦位已满
	}

	// 查询用户麦位信息， 如果已经有麦位相关信息，则返回不用操作
	userMic, err := l.rep.RoomManagerOnMicer.GetUserMicStatus(l.ctx, int64(in.Room), in.User)
	if err != nil && !errs.RecordNotFound(err) { //排除没用户记录
		return nil, err
	}

	if userMic != nil && userMic.UidType == 3 && userMic.RoomId == in.Room { // 用户是申请上麦的状态 ，进行上麦批准

		micStatus := 1
		err = l.rep.RoomManagerOnMicer.UpdateMicStatus(l.ctx, l.rep.Mysql, in.User, in.Room, int64(micStatus))
		if err != nil {
			return nil, err
		}
		//上麦成功， 广播消息
		//需要删除影响到的redis Key 记录
		delKey := "cache:socialIm:roomId:invitedMic:" + strconv.Itoa(int(in.Room))
		l.svcCtx.Redis.Del(delKey)

		delKey = "cache:socialIm:roomId:micStatus:" + strconv.Itoa(int(in.Room)) + ":" + strconv.Itoa(int(in.User))
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

		//	roomCreater := strconv.FormatInt(room.CreateUser, 10)
		//user := strconv.FormatInt(in.User, 10)

		_, err = l.svcCtx.UserRpc.SendRtmChannel(l.ctx, &pb.SendRtmChannelReq{From: "1", ChannelName: room.Mark, MessageType: agora.ACCEPTMIC, MessageBody: string(msgdata)}) //给房主发消息
		//_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: CreaterOrManager, To: roomCreater, MessageType: "approve mic ok", MessageBody: "approve mic ok"}) //给房主发消息
		//_, err = UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: CreaterOrManager, To: CreaterOrManager, MessageType: "approve mic ok", MessageBody: "approve mic ok"}) //给房主发消息

		if err != nil {
			return nil, err
		}

		res := pb.MicStatus{Status: "ok", Code: 0}
		return &pb.AcceptMicResq{Status: &res}, err
	}
	return nil, xerr.NewErrMsg("用户麦位状态冲突， 请稍后")

}
