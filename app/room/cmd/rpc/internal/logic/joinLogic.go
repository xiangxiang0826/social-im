package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"social-im/app/room/cmd/rpc/internal/repository"
	"social-im/app/room/cmd/rpc/internal/svc"
	"social-im/app/room/cmd/rpc/pb"
	"social-im/app/room/model"
	"social-im/app/user/cmd/rpc/userrpc"
	"social-im/common/rediskey"
	errTypes "social-im/common/types"
	"social-im/common/xorm/errs"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/logx"
	bloomRedis "github.com/zeromicro/go-zero/core/stores/redis"
)

type JoinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

type User struct {
	Avatar string `json:"Avatar"`
}

func NewJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinLogic {
	return &JoinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *JoinLogic) Join(in *pb.PartyJoinReq) (*pb.PartyJoinResp, error) {
	// todo: add your logic here and delete this line
	//1 判断roomid是否存在，判断userid是否存在
	//2 取roomid的最后一次加入记录，判断冷却时间是否ok
	//3 加入房间
	//4 写入redisList玩家的头像

	//1 判断roomid是否存在
	roomRes, err := l.rep.RoomModel.FindOne(l.ctx, in.RoomId) //判断该用户当前是否已经有在开启的派对房间
	if err != nil && err != errs.ErrNotFound {
		fmt.Println("RoomModel.FindOne err")
		return &pb.PartyJoinResp{
			Iret: 1,
			Smsg: errTypes.ErrSysBusy.Error(),
		}, nil
	}
	if err == errs.ErrNotFound {
		return &pb.PartyJoinResp{
			Iret: 1,
			Smsg: errTypes.ErrPartyNotfound.Error(),
		}, nil
	}

	_ = l.saveRoomAvatar(l.rep.Redis, rediskey.CacheSocialImRoomAvatarList+strconv.FormatInt(in.RoomId, 10), strconv.FormatInt(in.Uid, 10))

	//计算房间uv
	redisConf := bloomRedis.RedisConf{
		Host: l.svcCtx.Config.RedisConf.Host,
		Type: l.svcCtx.Config.RedisConf.Host,
		Pass: l.svcCtx.Config.RedisConf.Pass,
	}
	redisBloomConn := bloomRedis.MustNewRedis(redisConf, func(r *bloomRedis.Redis) {
		r.Type = l.svcCtx.Config.RedisConf.Type
		r.Pass = l.svcCtx.Config.RedisConf.Pass
	})
	filter := bloom.New(redisBloomConn, rediskey.CacheSocialImRoomBloom+strconv.FormatInt(in.RoomId, 10), 64)
	if isExist, _ := filter.Exists([]byte(strconv.FormatInt(in.Uid, 10))); isExist == false {
		filter.Add([]byte(strconv.FormatInt(in.Uid, 10)))
		l.rep.Redis.Incr(l.ctx, rediskey.CacheSocialImRoomUV+strconv.FormatInt(in.RoomId, 10)).Err()
	}
	// defer redisBloomConn.Close()

	//2 取roomid的最后一次加入记录，判断冷却时间是否ok
	oldRoomUser, err := l.rep.RoomUserModel.FindOneByRoomId(l.ctx, in.RoomId, in.Uid)
	fmt.Printf("oldroomuser is %v \n", oldRoomUser)
	if err != nil && err != errs.ErrNotFound {
		fmt.Println("RoomModel.FindOneByRoomId err %v \n", err)
		return &pb.PartyJoinResp{
			Iret: 1,
			Smsg: errTypes.ErrSysBusy.Error(),
		}, nil
	}

	if oldRoomUser != nil && oldRoomUser.CoolAt.Unix() > 0 {
		if oldRoomUser.CoolAt.Unix() > time.Now().Unix() {
			return &pb.PartyJoinResp{
				Iret: 2,
				Smsg: errTypes.ErrPartyCoolTime.Error(),
			}, nil
		}
	}

	//插入新的记录
	roomUser := &model.AppRoomUser{}
	roomUser.PartyId = in.RoomId
	roomUser.UserId = in.Uid
	if err := l.rep.RoomUserModel.Insert(l.ctx, l.rep.Mysql, roomUser); err != nil {
		return &pb.PartyJoinResp{
			Iret: 1,
			Smsg: errTypes.ErrPartyNotfound.Error(),
		}, nil
	}

	// roomId := strconv.FormatInt(in.RoomId, 10)
	// uid := strconv.FormatInt(in.Uid, 10)
	// //写入加入房间记录
	// err = l.rep.Redis.HSetNX(l.ctx, rediskey.CacheSocialImRoomOnLineHash+roomId, uid, 1).Err()
	// if err != nil {
	// 	fmt.Println("Redis.HSetNX err")
	// 	return &pb.PartyJoinResp{
	// 		Iret: 1,
	// 		Smsg: err.Error(),
	// 	}, nil
	// }
	rtcToken, _ := l.rep.UserRpc.GetRtcToken(l.ctx, &userrpc.GetRtcTokenReq{
		Uid:         strconv.FormatInt(in.Uid, 10),
		ChannelName: roomRes.Mark,
	})
	fmt.Printf("roomres is %v \n", roomRes.CreatedAt.Unix())
	return &pb.PartyJoinResp{
		RoomId:        roomRes.Id,
		Mark:          roomRes.Mark,
		Name:          roomRes.Name,
		BackgroundUrl: roomRes.BackgroundUrl,
		PartyType:     roomRes.PartyType,
		CreateUser:    roomRes.CreateUser,
		CreateAt:      roomRes.CreatedAt.Unix(),
		RtcToken:      rtcToken.RtcToken,
	}, nil
}

func (l *JoinLogic) saveRoomAvatar(redisConn redis.UniversalClient, redisListKey, userId string) error {
	userString := redisConn.Get(l.ctx, rediskey.CacheSocialImUserIdPrefix+userId).Val()

	fmt.Printf("userstring is %v \n", userString)
	if len(userString) == 0 {
		fmt.Println("error saveRoomAvatar len")
		return nil
	}
	var userAvatar User
	err := json.Unmarshal([]byte(userString), &userAvatar)
	if err != nil {
		fmt.Println("error saveRoomAvatar Unmarshal")
		return err
	}

	err = redisConn.RPush(l.ctx, redisListKey, userAvatar.Avatar).Err()
	if err != nil {
		fmt.Println("error saveRoomAvatar RPush")
		return nil
	}

	// 如果队列长度大于5，则删除队列头部元素
	length, err := redisConn.LLen(l.ctx, redisListKey).Result()
	if err != nil {
		fmt.Println("error saveRoomAvatar LLen")
		return err
	}
	if length > 5 {
		err = redisConn.LTrim(l.ctx, redisListKey, 1, -1).Err()
		if err != nil {
			fmt.Println("error saveRoomAvatar LTrim")
			return err
		}
	}
	return nil
}
