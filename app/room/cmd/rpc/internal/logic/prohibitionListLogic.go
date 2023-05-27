package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"social-im/app/room/cmd/rpc/internal/repository"
	"social-im/app/room/cmd/rpc/internal/svc"
	"social-im/app/room/cmd/rpc/pb"
	"social-im/app/room/model"
	userPb "social-im/app/user/cmd/rpc/pb"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"
	"sort"
)

type ProhibitionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewProhibitionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProhibitionListLogic {
	return &ProhibitionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *ProhibitionListLogic) ProhibitionList(in *pb.ProhibitionListReq) (*pb.ProhibitionListResp, error) {
	// todo: add your logic here and delete this line
	total, list, err := l.rep.RoomProhibitionUsersModel.RoomProhibitionUsersList(l.ctx, in.RoomId, in.RoomType, in.LastId, in.PageSize)
	if err != nil && err != errs.ErrNotFound {
		return nil, xerr.NewErrMsg(err.Error())
	}
	var listResp []*pb.ProhibitionInfo
	if len(list) > 0 {
		r, err := mr.MapReduce(func(source chan<- *model.AppRoomProhibitionUsers) {
			for _, eachProhibitionUsersData := range list {
				source <- eachProhibitionUsersData
			}
		}, func(item *model.AppRoomProhibitionUsers, writer mr.Writer[*pb.ProhibitionInfo], cancel func(error)) {
			prohibitionInfoData := item
			var pbProhibitionInfo pb.ProhibitionInfo
			_ = copier.Copy(&pbProhibitionInfo, prohibitionInfoData)
			pbProhibitionInfo.Id = prohibitionInfoData.Id
			pbProhibitionInfo.Status = prohibitionInfoData.Status
			pbProhibitionInfo.RoomId = prohibitionInfoData.RoomId
			pbProhibitionInfo.RoomType = prohibitionInfoData.RoomType
			pbProhibitionInfo.Uid = prohibitionInfoData.Uid
			pbProhibitionInfo.OperatorUser = prohibitionInfoData.OperatorUser
			userInfo, _ := l.rep.UserRpc.GetUserInfo(l.ctx, &userPb.GetUserInfoReq{
				Id: prohibitionInfoData.Uid,
			})
			if userInfo.Iret == 0 && userInfo.UserInfo != nil {
				pbProhibitionInfo.Avatar = userInfo.UserInfo.Avatar
				pbProhibitionInfo.UserName = userInfo.UserInfo.NickName
			}
			pbProhibitionInfo.CreatedAt = prohibitionInfoData.CreateTime.Unix()
			writer.Write(&pbProhibitionInfo)
		}, func(pipe <-chan *pb.ProhibitionInfo, writer mr.Writer[[]*pb.ProhibitionInfo], cancel func(error)) {
			var resp []*pb.ProhibitionInfo
			for p := range pipe {
				resp = append(resp, p)
			}
			writer.Write(resp)
		})
		if err != nil {
			return nil, xerr.NewErrMsg(err.Error())
		}
		listResp = r
	}
	sort.SliceStable(listResp, func(i, j int) bool {
		return listResp[i].Id > listResp[j].Id
	})
	return &pb.ProhibitionListResp{Total: total, List: listResp}, nil
}
