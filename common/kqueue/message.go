// KqMessage
package kqueue

//第三方支付回调更改支付状态通知
type ThirdGiftSendNotifyMessage struct {
	From      int64 `json:"from"`
	To        int64 `json:"to"`
	GiftId    int64 `json:"giftId"`
	GiftCount int64 `json:"giftCount"`
	GiftAttr  int64 `json:"giftAttr"`
}

var gold = int64(1)
