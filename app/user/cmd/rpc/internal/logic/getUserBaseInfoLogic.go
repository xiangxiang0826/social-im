package logic

import (
	"context"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"
	"strings"

	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewGetUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserBaseInfoLogic {
	return &GetUserBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *GetUserBaseInfoLogic) GetUserBaseInfo(in *pb.GetUserBaseReq) (*pb.GetUserBaseResp, error) {
	// todo: add your logic here and delete this line
	userBaseInfo, err := l.rep.UserBaseModel.FindOneByUid(l.ctx, in.Uid)
	if err != nil && !errs.RecordNotFound(err) {
		return nil, xerr.NewErrMsg(err.Error())
	}
	if errs.RecordNotFound(err) { //为空返回
		return &pb.GetUserBaseResp{}, nil
	}
	var notEmptyBaseFields []string
	if len(userBaseInfo.BodyHeight) > 0 {
		notEmptyBaseFields = append(notEmptyBaseFields, "body_height")
	}
	if len(userBaseInfo.Constellation) > 0 {
		notEmptyBaseFields = append(notEmptyBaseFields, "constellation")
	}
	if len(userBaseInfo.Education) > 0 {
		notEmptyBaseFields = append(notEmptyBaseFields, "education")
	}
	if len(userBaseInfo.School) > 0 {
		notEmptyBaseFields = append(notEmptyBaseFields, "school")
	}
	if len(userBaseInfo.Career) > 0 {
		notEmptyBaseFields = append(notEmptyBaseFields, "career")
	}
	if len(userBaseInfo.Hobby) > 0 {
		notEmptyBaseFields = append(notEmptyBaseFields, "hobby")
	}
	if len(userBaseInfo.Address) > 0 {
		notEmptyBaseFields = append(notEmptyBaseFields, "address")
	}
	if len(userBaseInfo.WorkAddress) > 0 {
		notEmptyBaseFields = append(notEmptyBaseFields, "work_address")
	}
	return &pb.GetUserBaseResp{
		Id:                 userBaseInfo.Id,
		Uid:                userBaseInfo.Uid,
		Constellation:      userBaseInfo.Constellation,
		BodyHeight:         userBaseInfo.BodyHeight,
		Education:          userBaseInfo.Education,
		School:             userBaseInfo.School,
		Career:             userBaseInfo.Career,
		Hobby:              userBaseInfo.Hobby,
		Address:            userBaseInfo.Address,
		WorkAddress:        userBaseInfo.WorkAddress,
		ProvinceId:         userBaseInfo.ProvinceId,
		CityId:             userBaseInfo.CityId,
		DistrictId:         userBaseInfo.DistrictId,
		WorkProvinceId:     userBaseInfo.WorkProvinceId,
		WorkCityId:         userBaseInfo.WorkCityId,
		WorkDistrictId:     userBaseInfo.WorkDistrictId,
		AboutMe:            userBaseInfo.AboutMe,
		BackgroundUrl:      userBaseInfo.BackgroundUrl,
		BackgroundSmallUrl: userBaseInfo.BackgroundSmallUrl,
		DisplayBaseFields:  strings.Split(userBaseInfo.DisplayFields, ","),
		NotEmptyBaseFields: notEmptyBaseFields,
	}, nil
}
