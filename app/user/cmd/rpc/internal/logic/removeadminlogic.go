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

type RemoveAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewRemoveAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveAdminLogic {
	return &RemoveAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *RemoveAdminLogic) RemoveAdmin(in *pb.RemoveAdminReq) (*pb.RemoveAdminResp, error) {
	// todo: add your logic here and delete this line
	userAdmin, err := l.rep.RoomManagerOnMicer.GetUserAdminStatus(l.ctx, int64(in.Room), in.User)
	if err != nil && !errs.RecordNotFound(err) { //排除没用户记录
		return nil, err
	}

	if (userAdmin != nil && userAdmin.UidType != -1) || (errs.RecordNotFound(err)) { //用户不是上麦的状态  或者没有麦位状态
		return nil, xerr.NewErrMsg("管理员状态有误， 无法进行删除管理员操作")
	}

	if userAdmin != nil && userAdmin.UidType == -1 && userAdmin.RoomId == in.Room {
		err = l.rep.RoomManagerOnMicer.DeleteUserAdminStatus(l.ctx, l.rep.Mysql, in.User, in.Room)

		if err != nil {
			return nil, err
		}

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
		//roomCreater := strconv.FormatInt(room.CreateUser, 10)
		//ssuser := strconv.FormatInt(in.User, 10)
		//CreaterOrManager := strconv.FormatInt(in.Uid, 10)
		//	_, err = l.svcCtx.UserRpc.SendRtmChannel(l.ctx, &pb.SendRtmChannelReq{From: "1", ChannelName: room.Mark, MessageType: "4", MessageBody: string(msgdata)}) //给房主发消息
		_, err = l.svcCtx.UserRpc.SendRtmChannel(l.ctx, &pb.SendRtmChannelReq{From: "1", ChannelName: room.Mark, MessageType: agora.REMOVEADMIN, MessageBody: string(msgdata)}) //给房主发

		if err != nil {
			return nil, err
		}

		res := pb.CommonResp{Status: "ok", Code: 0}
		return &pb.RemoveAdminResp{Resp: &res}, err

	}
	return nil, xerr.NewErrMsg("用户管理员状态冲突， 请稍后")

}
