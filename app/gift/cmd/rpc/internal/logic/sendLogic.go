package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"social-im/app/admin/cmd/rpc/adminrpc"
	"social-im/app/gift/cmd/rpc/internal/repository"
	"social-im/app/gift/cmd/rpc/internal/svc"
	"social-im/app/gift/cmd/rpc/pb"
	"social-im/app/gift/model"
	"social-im/app/user/cmd/rpc/userrpc"
	"social-im/common/agora"
	"social-im/common/kqueue"
	errTypes "social-im/common/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type SendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *SendLogic) Send(in *pb.GiftSendReq) (*pb.GiftSendResp, error) {
	// todo: add your logic here and delete this line
	//1、查找用户余额是否满足
	//2、事务 （扣除发送者余额，写入发货流水表）
	//3、发kafka消息

	//1、查找用户余额是否满足
	item, err := l.rep.AdminRpc.GiftItem(l.ctx, &adminrpc.GiftItemReq{
		GiftId: in.GiftId,
	})
	fmt.Printf("item is %v \n", item)
	if err != nil {
		fmt.Println("l.rep.AdminRpc.GiftItem error")
		return &pb.GiftSendResp{
			Iret: errTypes.ErrSysError,
			Smsg: err.Error(),
		}, nil
	}
	toList := strings.Split(in.SendTo, ",")

	//增送免费礼物
	if item.ChargeNum == 0 {
		err = l.sendFree(in, item, toList)
	} else { //增送付费礼物
		err = l.sendPaid(in, item, toList)
	}

	if err != nil {
		fmt.Printf("l.rep.GiftFlowModel.BatchInsert error %v \n", err)
		return &pb.GiftSendResp{
			Iret: errTypes.ErrCodeSysBusy,
			Smsg: errTypes.ErrSysBusy.Error(),
		}, nil
	}

	//TODO 3、发kafka消息
	for i := 0; i < len(toList); i++ {
		toUid, _ := strconv.ParseInt(toList[i], 10, 64)
		if err := l.pubKqSendSuccess(in.Uid, in.GiftId, in.GiftNum, toUid, item.GiftAttr); err != nil {
			fmt.Printf("l.pubKqPaySuccess : %+v", err)
		}
	}

	//发频道消息
	messageBody, _ := json.Marshal(in)
	l.rep.UserRpc.SendRtmChannel(l.ctx, &userrpc.SendRtmChannelReq{
		From:        agora.ADMINUSER,
		ChannelName: in.RoomMark,
		MessageType: agora.ToUsersGift,
		MessageBody: string(messageBody),
	})
	return &pb.GiftSendResp{}, nil
}

func (l *SendLogic) pubKqSendSuccess(Uid, giftId, giftNum, toUid int64, attr string) error {
	isAttr := int64(0)
	if len(attr) > 0 {
		isAttr = 1
	}
	m := kqueue.ThirdGiftSendNotifyMessage{
		From:      Uid,
		To:        toUid,
		GiftId:    giftId,
		GiftCount: giftNum,
		GiftAttr:  isAttr,
	}

	body, err := json.Marshal(m)
	if err != nil {
		return err
	}
	fmt.Printf("kqueue send is %v \n", string(body))
	return l.svcCtx.KqGiftSendClient.Push(string(body))
}

// 发送免费礼物
func (l *SendLogic) sendFree(in *pb.GiftSendReq, item *adminrpc.GiftItemResp, toList []string) error {
	fmt.Println("l.sendFree gift")
	if len(toList) < 1 {
		return errTypes.ErrToList
	}
	flowList := []*model.AppGiftFlow{}
	for i := 0; i < len(toList); i++ {
		var err error
		flowItem := &model.AppGiftFlow{}
		flowItem.FromUid = in.Uid
		flowItem.ToUid, err = strconv.ParseInt(toList[i], 10, 64)
		if err != nil {
			return errTypes.ErrToList
		}
		flowItem.GiftCount = in.GiftNum
		flowItem.GiftAttr = item.GiftAttr
		flowItem.GiftId = in.GiftId
		flowList = append(flowList, flowItem)
	}

	err := l.rep.GiftFlowModel.BatchInsert(l.ctx, l.rep.Mysql, flowList)
	if err != nil {
		return err
	}
	return nil
}

// 发送付费礼物
func (l *SendLogic) sendPaid(in *pb.GiftSendReq, item *adminrpc.GiftItemResp, toList []string) error {
	fmt.Println("l.sendPaid gift")
	balance, err := l.rep.BalanceModel.FindOneByUidType(l.ctx, in.Uid, item.Currency)
	if err != nil {
		fmt.Println("l.rep.BalanceModel.FindOneByUidType")
		return err
	}
	money := item.ChargeNum * in.GiftNum * int64(len(toList))
	if money > balance.Balance {
		fmt.Println("balance not enough")
		return err
	}

	//2、事务 （扣除发送者余额，写入发货流水表）
	fmt.Printf("balance is %v, cost monney is %v \n", balance.Balance, money)
	balance.Balance = balance.Balance - money

	flowList := []*model.AppGiftFlow{}
	for i := 0; i < len(toList); i++ {
		flowItem := &model.AppGiftFlow{}
		flowItem.FromUid = in.Uid
		flowItem.ToUid, _ = strconv.ParseInt(toList[i], 10, 64)
		flowItem.GiftCount = in.GiftNum
		flowItem.GiftAttr = item.GiftAttr
		flowItem.GiftId = in.GiftId
		flowList = append(flowList, flowItem)
	}

	err = l.rep.BalanceModel.Transaction(l.ctx, func(db *gorm.DB) error {
		err := l.rep.BalanceModel.Update(l.ctx, db, balance)
		if err != nil {
			fmt.Printf("l.rep.BalanceModel.Update error %v \n", err)
			return err
		}

		err = l.rep.GiftFlowModel.BatchInsert(l.ctx, db, flowList)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
