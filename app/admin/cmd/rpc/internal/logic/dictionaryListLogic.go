package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"social-im/app/admin/cmd/rpc/internal/repository"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"

	"social-im/app/admin/cmd/rpc/internal/svc"
	"social-im/app/admin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictionaryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictionaryListLogic {
	return &DictionaryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *DictionaryListLogic) DictionaryList(in *pb.DictionaryGetReq) (*pb.DictionaryGetResp, error) {
	// todo: add your logic here and delete this line
	list, err := l.rep.GetDictionarList(l.ctx, in.DictionaryType)
	if err != nil && err != errs.ErrNotFound {
		return nil, xerr.NewErrMsg(err.Error())
	}
	var resp []*pb.DictionaryDetailConf
	if len(list) > 0 {
		for _, dictionaryDetailConf := range list {
			var pbDictionaryDetailConf pb.DictionaryDetailConf
			_ = copier.Copy(&pbDictionaryDetailConf, dictionaryDetailConf)
			pbDictionaryDetailConf.Value = dictionaryDetailConf.Value.Int64
			pbDictionaryDetailConf.Label = dictionaryDetailConf.Label.String
			pbDictionaryDetailConf.DictionaryId = dictionaryDetailConf.SysDictionaryId.Int64
			resp = append(resp, &pbDictionaryDetailConf)
		}
	}
	return &pb.DictionaryGetResp{List: resp}, nil
}
