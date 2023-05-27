package logic

import (
	"context"
	roomPb "social-im/app/room/cmd/rpc/pb"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/common/agora"
	"social-im/common/rediskey"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"
	"strconv"

	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"

	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeclineAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewDeclineAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeclineAdminLogic {
	return &DeclineAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *DeclineAdminLogic) DeclineAdmin(in *pb.DeclineAdminReq) (*pb.DeclineAdminResq, error) {
	// todo: add your logic here and delete this line
	// 查询用户麦位信息， 如果已经有麦位相关信息，则返回不用操作
	userAdmin, err := l.rep.RoomManagerOnMicer.GetUserAdminStatus(l.ctx, int64(in.Room), in.User)
	if err != nil && !errs.RecordNotFound(err) { //排除没用户记录
		return nil, err
	}

	if userAdmin != nil && userAdmin.UidType == -2 && userAdmin.RoomId == in.Room { //删除用户管理员状态
		err := l.rep.RoomManagerOnMicer.DeleteUserAdminStatus(l.ctx, l.rep.Mysql, in.User, in.Room)
		//err = l.rep.RoomManagerOnMicer.UpdateMicStatus(l.ctx, l.rep.Mysql, in.User, in.Room, int64(micStatus))
		if err != nil {
			return nil, err
		}
		//上麦成功， 广播消息
		//需要删除影响到的redis Key 记录
		delKey := rediskey.CacheSocialImRoomUserAdminStatusPrefix + strconv.Itoa(int(in.Room)) + ":" + strconv.Itoa(int(in.User))
		l.svcCtx.Redis.Del(delKey)

		room, err := l.svcCtx.RoomRpc.RoominfoById(l.ctx, &roomPb.RoominfoReq{Room: in.Room})
		if err != nil {
			return nil, err
		}
		//	roomCreater := strconv.FormatInt(room.CreateUser, 10)
		//user := strconv.FormatInt(in.User, 10)
		type uuser struct {
			UserId int64 `json:"user_id"`
		}
		var msgUser uuser
		msgUser.UserId = in.User
		msgdata, _ := json.Marshal(&msgUser)

		_, err = l.svcCtx.UserRpc.SendRtmChannel(l.ctx, &pb.SendRtmChannelReq{From: "1", ChannelName: room.Mark, MessageType: agora.DECLINEADMIN, MessageBody: string(msgdata)}) //给房主发消息
		//_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: CreaterOrManager, To: roomCreater, MessageType: "approve mic ok", MessageBody: "approve mic ok"}) //给房主发消息
		//_, err = UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: CreaterOrManager, To: CreaterOrManager, MessageType: "approve mic ok", MessageBody: "approve mic ok"}) //给房主发消息
		if err != nil {
			return nil, err
		}

		res := pb.CommonResp{Status: "ok", Code: 0}
		return &pb.DeclineAdminResq{Resp: &res}, err

	}
	return nil, xerr.NewErrMsg("用户管理员状态冲突， 请稍后")
}
