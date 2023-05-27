package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"social-im/app/gift/cmd/rpc/internal/repository"
	"social-im/app/gift/cmd/rpc/internal/svc"
	"social-im/app/gift/cmd/rpc/pb"
	"social-im/app/gift/model"
	"social-im/common/xorm/errs"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateAttrLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewUpdateAttrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAttrLogic {
	return &UpdateAttrLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *UpdateAttrLogic) UpdateAttr(in *pb.GiftUpdateAttrReq) (*pb.GiftUpdateAttrResp, error) {
	// todo: add your logic here and delete this line
	//1、获取flow表记录
	//2、获取giftbag表记录
	//3、解析Attr数值
	//4、事务更新对应的数据
	//5、失败之后的处理
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateBag: %+v\n", r)
		}
	}()

	//1、获取flow表记录
	flowData, err := l.rep.GiftFlowModel.FindOneByAttrStatus(l.ctx, in.Uid, in.SendTo, in.GiftId, in.GiftNum)
	if err != nil {
		return nil, err
	}
	flowData.AttrStatus = 1

	//2、获取balance表记录
	_, balanceData, err := l.rep.BalanceModel.FindListByUid(l.ctx, in.SendTo)
	if err != nil && !errs.RecordNotFound(err) {
		return nil, err
	}

	//3、json字符串解析
	attrStr := flowData.GiftAttr
	fmt.Printf("update attr flowdata is %v \n", flowData)
	if len(attrStr) == 0 {
		err := l.rep.GiftFlowModel.Update(l.ctx, l.rep.Mysql, flowData)
		if err != nil {
			fmt.Printf("update attr flowdata update err is %v \n", err)
		}
		return nil, errors.New("wrong attr")
	}
	var attrStruct model.AppGiftAttr
	err = json.Unmarshal([]byte(attrStr), &attrStruct)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("json.Unmarshal error")
	}

	//4、处理金币
	err = l.rep.GiftFlowModel.Transaction(l.ctx, func(db *gorm.DB) error {
		if attrStruct.Gold > 0 {
			balanceItem, _ := l.getData(balanceData, 1)
			if balanceItem == nil {
				balanceItem = &model.AppUserBalance{}
				balanceItem.Uid = flowData.ToUid
				balanceItem.Balance = attrStruct.Gold
				balanceItem.Type = 1
				err := l.rep.BalanceModel.Insert(l.ctx, l.rep.Mysql, balanceItem)
				if err != nil {
					fmt.Printf("l.rep.BalanceModel.Insert gold error %v \n", err)
					return err
				}
			} else {
				balanceItem.Balance += attrStruct.Gold
				err := l.rep.BalanceModel.Update(l.ctx, l.rep.Mysql, balanceItem)
				if err != nil {
					fmt.Printf("l.rep.BalanceModel.update gold error %v \n", err)
					return err
				}
			}

			err = l.rep.GiftFlowModel.Update(l.ctx, db, flowData)
			if err != nil {
				fmt.Printf("l.rep.GiftFlowModel.Update error %v \n", err)
				return err
			}
		}

		//4、处理钻石
		if attrStruct.Diamond > 0 {
			balanceItem, _ := l.getData(balanceData, 2)
			// err = l.rep.GiftFlowModel.Transaction(l.ctx, func(db *gorm.DB) error {
			if balanceItem == nil {
				balanceItem = &model.AppUserBalance{}
				balanceItem.Uid = flowData.ToUid
				balanceItem.Balance = attrStruct.Diamond
				balanceItem.Type = 2
				err := l.rep.BalanceModel.Insert(l.ctx, l.rep.Mysql, balanceItem)
				if err != nil {
					fmt.Printf("l.rep.BalanceModel.Insert gold error %v \n", err)
					return err
				}
			} else {
				balanceItem.Balance += attrStruct.Diamond
				err := l.rep.BalanceModel.Update(l.ctx, l.rep.Mysql, balanceItem)
				if err != nil {
					fmt.Printf("l.rep.BalanceModel.update gold error %v \n", err)
					return err
				}
			}

			err = l.rep.GiftFlowModel.Update(l.ctx, db, flowData)
			if err != nil {
				fmt.Printf("l.rep.GiftFlowModel.Update error %v \n", err)
				return err
			}
			// return nil
		}

		//4、处理魅力值
		if attrStruct.Charm > 0 {
			balanceItem, _ := l.getData(balanceData, 3)
			// err = l.rep.GiftFlowModel.Transaction(l.ctx, func(db *gorm.DB) error {
			if balanceItem == nil {
				balanceItem = &model.AppUserBalance{}
				balanceItem.Uid = flowData.ToUid
				balanceItem.Balance = attrStruct.Charm
				balanceItem.Type = 3
				err := l.rep.BalanceModel.Insert(l.ctx, l.rep.Mysql, balanceItem)
				if err != nil {
					fmt.Printf("l.rep.BalanceModel.Insert gold error %v \n", err)
					return err
				}
			} else {
				balanceItem.Balance += attrStruct.Charm
				err := l.rep.BalanceModel.Update(l.ctx, l.rep.Mysql, balanceItem)
				if err != nil {
					fmt.Printf("l.rep.BalanceModel.update gold error %v \n", err)
					return err
				}
			}

			err = l.rep.GiftFlowModel.Update(l.ctx, db, flowData)
			if err != nil {
				fmt.Printf("l.rep.GiftFlowModel.Update error %v \n", err)
				return err
			}
			// return nil
		}
		return nil
	})

	return &pb.GiftUpdateAttrResp{}, nil
}

func (l *UpdateAttrLogic) getData(balanceData []*model.AppUserBalance, attr int64) (*model.AppUserBalance, error) {
	if len(balanceData) == 0 {
		return nil, nil
	}

	for _, v := range balanceData {
		if v.Type == attr {
			return v, nil
		}
	}
	return nil, nil
}
