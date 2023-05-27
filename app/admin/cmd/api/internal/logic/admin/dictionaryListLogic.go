package admin

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/admin/cmd/api/internal/svc"
	"social-im/app/admin/cmd/api/internal/types"
	"social-im/app/admin/cmd/rpc/adminrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictionaryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictionaryListLogic {
	return &DictionaryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictionaryListLogic) DictionaryList(req *types.DictionaryGetReq) (*types.DictionaryGetResp, error) {
	// todo: add your logic here and delete this line
	backgroundResp, err := l.svcCtx.AdminRpc.DictionaryList(l.ctx, &adminrpc.DictionaryGetReq{
		DictionaryType: req.DictionaryType,
	})
	if err != nil {
		return nil, err
	}
	var resp types.DictionaryGetResp
	_ = copier.Copy(&resp, backgroundResp)
	return &resp, nil
}
