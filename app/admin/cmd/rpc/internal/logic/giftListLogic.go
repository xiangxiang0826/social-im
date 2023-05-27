package logic

import (
	"context"
	"fmt"

	"social-im/app/admin/cmd/rpc/internal/repository"
	"social-im/app/admin/cmd/rpc/internal/svc"
	"social-im/app/admin/cmd/rpc/pb"
	errTypes "social-im/common/types"
	"social-im/common/xorm/errs"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GiftListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewGiftListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GiftListLogic {
	return &GiftListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *GiftListLogic) GiftList(in *pb.GiftListReq) (*pb.GiftListResp, error) {
	// todo: add your logic here and delete this line
	total, list, err := l.rep.GiftMoel.GiftList(l.ctx, in.LastId, in.PageSize)
	if err != nil && err != errs.ErrNotFound {
		return &pb.GiftListResp{
			Iret: errTypes.ErrSysError,
			Smsg: err.Error(),
		}, nil
	}

	var resp []*pb.GiftInfo
	if len(list) > 0 {
		for _, item := range list {
			var giftInfo pb.GiftInfo
			_ = copier.Copy(&giftInfo, item)
			resp = append(resp, &giftInfo)
		}
	}

	fmt.Printf("adminrpc GiftList is %v \n", resp)
	return &pb.GiftListResp{
		Total: total,
		List:  resp}, nil
}
