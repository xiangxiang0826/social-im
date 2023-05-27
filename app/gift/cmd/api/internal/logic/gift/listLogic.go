package gift

import (
	"context"
	"fmt"

	"social-im/app/admin/cmd/rpc/adminrpc"
	"social-im/app/gift/cmd/api/internal/svc"
	"social-im/app/gift/cmd/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.GiftListReq) (*types.GiftListResp, error) {
	// todo: add your logic here and delete this line
	giftResp, err := l.svcCtx.AdminRpc.GiftList(l.ctx, &adminrpc.GiftListReq{
		LastId:   req.LastId,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("List res is %v \n", giftResp)
	var resp types.GiftListResp
	_ = copier.Copy(&resp, giftResp)
	return &resp, nil
}
