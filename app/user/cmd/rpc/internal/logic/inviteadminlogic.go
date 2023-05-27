package logic

import (
	"context"
	roomPb "social-im/app/room/cmd/rpc/pb"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/model"
	"social-im/common/rediskey"
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

type InviteAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewInviteAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteAdminLogic {
	return &InviteAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *InviteAdminLogic) InviteAdmin(in *pb.InviteAdminReq) (*pb.InviteAdminResq, error) {
	// todo: add your logic here and delete this line
	admins, err := l.rep.RoomManagerOnMicer.FindUsersIsAdmin(l.ctx, int64(in.Room))
	if err != nil && !errs.RecordNotFound(err) { //数据库错误
		return nil, err
	}

	if len(*admins) >= 3 {
		return nil, xerr.NewErrMsg("管理员人数已满") //麦位已满
	}
	// 查询用户麦位信息， 如果已经有麦位相关信息，则返回不用操作
	userAdmin, err := l.rep.RoomManagerOnMicer.GetUserAdminStatus(l.ctx, int64(in.Room), in.User)
	if err != nil && !errs.RecordNotFound(err) { //排除没用户记录
		return nil, err
	}

	if userAdmin != nil { //用户已经有管理员状态
		return nil, xerr.NewErrMsg("用户已被邀请或者已是管理员， 请勿重复操作")
	}

	err = l.rep.RoomManagerOnMicer.Insert(l.ctx, l.rep.Mysql, &model.RoomManagerOnmicer{
		RoomId:   int64(in.Room),
		Type:     1,
		Uid:      in.User,
		UidType:  -2,
		JoinTime: time.Now().Unix(),
	})

	if err != nil {
		return nil, err

	}
	//需要删除影响到的redis Key 记录
	delKey := rediskey.CacheSocialImRoomUserAdminStatusPrefix + strconv.Itoa(int(in.Room)) + ":" + strconv.Itoa(int(in.User))
	l.svcCtx.Redis.Del(delKey)

	//sFrom := strconv.FormatInt(in.Uid, 10)
	sUser := strconv.FormatInt(in.User, 10)

	room, err := l.svcCtx.RoomRpc.RoominfoById(l.ctx, &roomPb.RoominfoReq{Room: in.Room})
	if err != nil {
		return nil, err
	}

	adminss, err := l.svcCtx.UserRpc.RoomAdminList(l.ctx, &pb.RoomAdminListReq{Room: in.Room})

	if err != nil {
		return nil, err
	}

	roomCreater := strconv.FormatInt(room.CreateUser, 10)

	type uuser struct {
		UserId int64 `json:"user_id"`
	}
	var msgUser uuser
	msgUser.UserId = in.User
	msgdata, _ := json.Marshal(&msgUser)

	for _, v := range adminss.Uids {
		amdinStr := strconv.FormatInt(v, 10)
		_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: "1", To: amdinStr, MessageType: agora.INVITEADMIN, MessageBody: string(msgdata)}) //给管理员发消息

	}
	_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: "1", To: roomCreater, MessageType: agora.INVITEADMIN, MessageBody: string(msgdata)}) //给房主发消息
	_, err = l.svcCtx.UserRpc.SendRtm(l.ctx, &pb.SendRtmReq{From: "1", To: sUser, MessageType: agora.INVITEADMIN, MessageBody: string(msgdata)})       //给房主发消息

	if err != nil {
		return nil, err
	}

	res := pb.CommonResp{Status: "ok", Code: 0}
	return &pb.InviteAdminResq{Resp: &res}, err

}
