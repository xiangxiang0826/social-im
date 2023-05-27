package logic

import (
	"context"
	roomPb "social-im/app/room/cmd/rpc/pb"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/common/rediskey"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"
	"strconv"

	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"

	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"social-im/common/agora"
)

type AcceptAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewAcceptAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AcceptAdminLogic {
	return &AcceptAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *AcceptAdminLogic) AcceptAdmin(in *pb.AcceptAdminReq) (*pb.AcceptAdminResq, error) {
	// todo: add your logic here and delete this line

	admins, err := l.rep.RoomManagerOnMicer.FindUsersIsAdmin(l.ctx, int64(in.Room))
	if err != nil && !errs.RecordNotFound(err) { //数据库错误
		return nil, err
	}
	if admins != nil && len(*admins) >= 3 {
		return nil, xerr.NewErrMsg("管理员人数已满") //麦位已满
	}

	// 查询用户麦位信息， 如果已经有麦位相关信息，则返回不用操作
	userAdmin, err := l.rep.RoomManagerOnMicer.GetUserAdminStatus(l.ctx, int64(in.Room), in.User)
	if err != nil && !errs.RecordNotFound(err) { //排除没用户记录
		return nil, err
	}

	if userAdmin != nil && userAdmin.UidType == -2 && userAdmin.RoomId == in.Room { // 用户是被邀请状态，同意
		adminStatus := -1
		err = l.rep.RoomManagerOnMicer.UpdateAdminStatus(l.ctx, l.rep.Mysql, in.User, in.Room, int64(adminStatus))
		if err != nil {
			return nil, err
		}
		//上麦成功， 广播消息
		//需要删除影响到的redis Key 记录
		delKey := rediskey.CacheSocialImRoomUserAdminStatusPrefix + strconv.Itoa(int(in.Room)) + ":" + strconv.Itoa(int(in.User))
		l.svcCtx.Redis.Del(delKey)

		delKey = rediskey.CacheSocialImRoomAdminPrefix + strconv.Itoa(int(in.Room))
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

		_, err = l.svcCtx.UserRpc.SendRtmChannel(l.ctx, &pb.SendRtmChannelReq{From: "1", ChannelName: room.Mark, MessageType: agora.ACCEPTADMIN, MessageBody: string(msgdata)}) //给房主发消息
		//_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: CreaterOrManager, To: roomCreater, MessageType: "approve mic ok", MessageBody: "approve mic ok"}) //给房主发消息
		//_, err = UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: CreaterOrManager, To: CreaterOrManager, MessageType: "approve mic ok", MessageBody: "approve mic ok"}) //给房主发消息

		if err != nil {
			return nil, err
		}

		res := pb.CommonResp{Status: "ok", Code: 0}
		return &pb.AcceptAdminResq{Resp: &res}, err
	}
	return nil, xerr.NewErrMsg("用户管理员状态冲突， 请稍后")
}
