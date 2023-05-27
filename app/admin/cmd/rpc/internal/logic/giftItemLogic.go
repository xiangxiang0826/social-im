package logic

import (
	"context"
	"fmt"

	"social-im/app/admin/cmd/rpc/internal/repository"
	"social-im/app/admin/cmd/rpc/internal/svc"
	"social-im/app/admin/cmd/rpc/pb"
	errTypes "social-im/common/types"
	"social-im/common/xorm/errs"

	"github.com/zeromicro/go-zero/core/logx"
)

type GiftItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewGiftItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GiftItemLogic {
	return &GiftItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *GiftItemLogic) GiftItem(in *pb.GiftItemReq) (*pb.GiftItemResp, error) {
	// todo: add your logic here and delete this line
	item, err := l.rep.GiftMoel.FindOne(l.ctx, in.GiftId)
	if err != nil && err != errs.ErrNotFound {
		return &pb.GiftItemResp{
			Iret: errTypes.ErrSysError,
			Smsg: err.Error(),
		}, nil
	}

	fmt.Printf("GiftItem is %v \n", item)

	return &pb.GiftItemResp{
		Id:        item.Id,
		GiftName:  item.GiftName,
		ImgUrl:    item.ImgUrl,
		Currency:  item.Currency,
		ChargeNum: item.ChargeNum,
		GiftAttr:  item.GiftAttr,
	}, nil
}
