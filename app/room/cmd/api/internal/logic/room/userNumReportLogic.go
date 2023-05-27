package room

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/room/cmd/rpc/roomrpc"

	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserNumReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserNumReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserNumReportLogic {
	return &UserNumReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserNumReportLogic) UserNumReport(req *types.UserNumReportReq) (*types.UserNumReportResp, error) {
	// todo: add your logic here and delete this line
	reportResp, err := l.svcCtx.RoomRpc.UserOnlineNumReport(l.ctx, &roomrpc.UserNumReportReq{
		Mark:       req.Mark,
		OnlineNums: req.OnlineNum,
	})
	if err != nil {
		return nil, err
	}
	var resp types.UserNumReportResp
	_ = copier.Copy(&resp, reportResp)
	return &resp, nil
}
