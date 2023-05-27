package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"
	"social-im/app/user/model"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"
)

type UpdateBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewUpdateBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBaseInfoLogic {
	return &UpdateBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *UpdateBaseInfoLogic) UpdateBaseInfo(in *pb.UpdateBaseReq) (*pb.UpdateBaseResp, error) {
	// todo: add your logic here and delete this line
	userBaseInfo, err := l.rep.UserBaseModel.FindOneByUid(l.ctx, in.Uid)
	if err != nil && !errs.RecordNotFound(err) {
		return nil, xerr.NewErrMsg(err.Error())
	}
	var currentUserBaseId int64
	var displayBaseFields []string
	if errs.RecordNotFound(err) { //如果不存在进行初始化
		userBaseInfo = &model.AppUserBase{}
	} else {
		currentUserBaseId = userBaseInfo.Id
	}
	userBaseInfo.Uid = in.Uid
	if len(in.BodyHeight) > 0 {
		userBaseInfo.BodyHeight = in.BodyHeight
		displayBaseFields = append(displayBaseFields, "body_height")
	}
	if len(in.Constellation) > 0 {
		userBaseInfo.Constellation = in.Constellation
		displayBaseFields = append(displayBaseFields, "constellation")
	}
	if len(in.Education) > 0 {
		userBaseInfo.Education = in.Education
		displayBaseFields = append(displayBaseFields, "education")
	}
	if len(in.School) > 0 {
		userBaseInfo.School = in.School
		displayBaseFields = append(displayBaseFields, "school")
	}
	if len(in.Career) > 0 {
		userBaseInfo.Career = in.Career
		displayBaseFields = append(displayBaseFields, "career")
	}
	if len(in.Hobby) > 0 {
		userBaseInfo.Hobby = in.Hobby
		displayBaseFields = append(displayBaseFields, "hobby")
	}
	if len(in.Address) > 0 {
		userBaseInfo.Address = in.Address
		displayBaseFields = append(displayBaseFields, "address")
	}
	if len(in.WorkAddress) > 0 {
		userBaseInfo.WorkAddress = in.WorkAddress
		displayBaseFields = append(displayBaseFields, "work_address")
	}
	if in.ProvinceId > 0 {
		userBaseInfo.ProvinceId = in.ProvinceId
	}
	if in.CityId > 0 {
		userBaseInfo.CityId = in.CityId
	}
	if in.DistrictId > 0 {
		userBaseInfo.DistrictId = in.DistrictId
	}
	if in.WorkProvinceId > 0 {
		userBaseInfo.WorkProvinceId = in.WorkProvinceId
	}
	if in.WorkCityId > 0 {
		userBaseInfo.WorkCityId = in.WorkCityId
	}
	if in.WorkDistrictId > 0 {
		userBaseInfo.WorkDistrictId = in.WorkDistrictId
	}
	if currentUserBaseId > 0 {
		err = l.rep.UserBaseModel.Update(l.ctx, l.rep.Mysql, userBaseInfo)
	} else {
		if len(displayBaseFields) <= 0 { //用户没有填写任何信息直接返回成功
			return &pb.UpdateBaseResp{
				Id:  currentUserBaseId,
				Uid: in.Uid,
			}, nil
		}
		//userBaseInfo.DisplayFields = strings.Join(displayBaseFields,",")
		err = l.rep.UserBaseModel.Insert(l.ctx, l.rep.Mysql, userBaseInfo)
	}
	if err != nil {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "")
	}
	return &pb.UpdateBaseResp{
		Id:  userBaseInfo.Id,
		Uid: userBaseInfo.Uid,
	}, nil
}
