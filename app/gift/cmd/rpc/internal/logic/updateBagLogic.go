package logic

import (
	"context"
	"fmt"

	"social-im/app/gift/cmd/rpc/internal/repository"
	"social-im/app/gift/cmd/rpc/internal/svc"
	"social-im/app/gift/cmd/rpc/pb"
	"social-im/app/gift/model"
	"social-im/common/xorm/errs"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateBagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewUpdateBagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBagLogic {
	return &UpdateBagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *UpdateBagLogic) UpdateBag(in *pb.GiftUpdateBagReq) (*pb.GiftUpdateBagResp, error) {
	// todo: add your logic here and delete this line
	//1、获取flow表记录
	//2、获取giftbag表记录
	//3、获取admin表记录
	//4、事务更新对应的数据
	//5、失败之后的处理
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateBag: %+v\n", r)
		}
	}()

	//1、获取flow表记录
	flowData, err := l.rep.GiftFlowModel.FindOneByMessage(l.ctx, in.Uid, in.SendTo, in.GiftId, in.GiftNum)
	if err != nil {
		return nil, err
	}

	//2、获取giftbag表记录
	bagData, err := l.rep.GiftBagModel.FindOneByUidGiftId(l.ctx, in.SendTo, in.GiftId)
	if err != nil && !errs.RecordNotFound(err) {
		return nil, err
	}

	//4、事务更新对应的数据
	flowData.Status = 1
	err = l.rep.GiftBagModel.Transaction(l.ctx, func(db *gorm.DB) error {
		if errs.RecordNotFound(err) || bagData == nil {
			//走insert 流程
			bagData = &model.AppGiftBag{}
			bagData.Uid = in.SendTo
			bagData.GiftId = in.GiftId
			bagData.GiftCount = in.GiftNum
			err := l.rep.GiftBagModel.Insert(l.ctx, db, bagData)
			fmt.Printf("insert err is %v \n", err)
			if err != nil {
				fmt.Printf("l.rep.GiftBagModel.Insert error %v \n", err)
				return err
			}
		} else {
			//走update流程
			bagData.GiftCount += in.GiftNum
			err := l.rep.GiftBagModel.Update(l.ctx, db, bagData)
			if err != nil {
				fmt.Printf("l.rep.GiftBagModel.Update error %v \n", err)
				return err
			}
		}

		err = l.rep.GiftFlowModel.Update(l.ctx, db, flowData)
		if err != nil {
			fmt.Printf("l.rep.GiftFlowModel.Update error %v \n", err)
			return err
		}
		return nil
	})

	//5、失败之后的处理
	if err != nil {
		fmt.Printf("l.rep.GiftFlowModel.BatchInsert error %v \n", err)
		return nil, err
	}

	return &pb.GiftUpdateBagResp{}, nil
}
