package user

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/user/cmd/rpc/userrpc"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectTagLogic {
	return &SelectTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectTagLogic) SelectTag(req *types.SelectTagReq) (*types.SelectTagResp, error) {
	// todo: add your logic here and delete this line
	updateSelectTagResp, err := l.svcCtx.UserRpc.SelectTag(l.ctx, &userrpc.SelectTagReq{
		Uid:           req.Uid,
		DisplayFields: req.DisplayFields,
	})
	if err != nil {
		return nil, err
	}
	var resp types.SelectTagResp
	_ = copier.Copy(&resp, updateSelectTagResp)
	return &resp, nil
}
