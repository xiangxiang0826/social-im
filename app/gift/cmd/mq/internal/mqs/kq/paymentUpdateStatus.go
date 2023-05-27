package kq

import (
	"context"
	"encoding/json"
	"fmt"
	"social-im/app/gift/cmd/mq/internal/svc"
	"social-im/app/gift/cmd/rpc/giftrpc"
	"social-im/common/kqueue"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

/*
*
Listening to the gift flow status change notification message queue
*/
type PaymentUpdateStatusMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentUpdateStatusMq(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentUpdateStatusMq {
	fmt.Println("message in NewPaymentUpdateStatusMq")
	return &PaymentUpdateStatusMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentUpdateStatusMq) Consume(_, val string) error {
	fmt.Printf("val in PaymentUpdateStatusMq %v \n", val)
	var message kqueue.ThirdGiftSendNotifyMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("PaymentUpdateStatusMq->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("PaymentUpdateStatusMq->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

func (l *PaymentUpdateStatusMq) execService(message kqueue.ThirdGiftSendNotifyMessage) error {
	//1、获取flow表记录
	//2、获取balance表记录
	//3、获取admin表记录
	//4、事务更新对应的数据
	//5、失败之后的处理
	fmt.Println(" message in execService")
	_, err := l.svcCtx.GiftRpc.UpdateBag(l.ctx, &giftrpc.GiftUpdateBagReq{
		Uid:     message.From,
		SendTo:  message.To,
		GiftId:  message.GiftId,
		GiftNum: message.GiftCount,
	})

	//处理礼物的属性，增加延迟避免行锁竞争
	if message.GiftAttr > 0 {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()

			time.Sleep(1 * time.Second)
			_, err := l.svcCtx.GiftRpc.UpdateAttr(l.ctx, &giftrpc.GiftUpdateAttrReq{
				Uid:     message.From,
				SendTo:  message.To,
				GiftId:  message.GiftId,
				GiftNum: message.GiftCount,
			})
			fmt.Printf("mq update attr err is %v \n", err)
		}()
		wg.Wait()
	}

	//todo 失败需要增加jobs的重试逻辑
	if err != nil {
		fmt.Printf(" exec update service err %v", err)
		return err
	}
	return nil
}
