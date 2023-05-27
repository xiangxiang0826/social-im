package gift

import (
	"context"

	"social-im/app/gift/cmd/api/internal/svc"
	"social-im/app/gift/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetItemLogic {
	return &GetItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetItemLogic) GetItem(req *types.GetItemReq) (resp *types.GetItemResp, err error) {
	// todo: add your logic here and delete this line

	return
}
