package user

import (
	"context"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
	"social-im/app/user/cmd/rpc/userrpc"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (*types.UserInfoResp, error) {
	// todo: add your logic here and delete this line
	detailResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userrpc.GetUserInfoReq{
		Id: req.Uid,
	})

	if err != nil {
		return nil, err
	}

	var resp types.UserInfoResp
	// fmt.Printf("AutoReg resp is %v", resp)
	_ = copier.Copy(&resp, detailResp)
	return &resp, nil
}
