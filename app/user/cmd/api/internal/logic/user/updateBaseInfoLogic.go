package user

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/user/cmd/rpc/userrpc"

	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBaseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBaseInfoLogic {
	return &UpdateBaseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBaseInfoLogic) UpdateBaseInfo(req *types.UpdateBaseReq) (*types.UpdateBaseResp, error) {
	// todo: add your logic here and delete this line
	updateBaseInfoResp, err := l.svcCtx.UserRpc.UpdateBaseInfo(l.ctx, &userrpc.UpdateBaseReq{
		Uid:            req.Uid,
		Constellation:  req.Constellation,
		BodyHeight:     req.BodyHeight,
		Education:      req.Education,
		School:         req.School,
		Career:         req.Career,
		Hobby:          req.Hobby,
		Address:        req.Address,
		WorkAddress:    req.WorkAddress,
		ProvinceId:     req.ProvinceId,
		CityId:         req.CityId,
		DistrictId:     req.DistrictId,
		WorkProvinceId: req.WorkProvinceId,
		WorkCityId:     req.WorkCityId,
		WorkDistrictId: req.WorkDistrictId,
	})
	if err != nil {
		return nil, err
	}
	var resp types.UpdateBaseResp
	_ = copier.Copy(&resp, updateBaseInfoResp)
	return &resp, nil
}
