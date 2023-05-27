package repository

import (
	"context"
	"social-im/app/admin/model"
)

func (rep *Rep) GetDictionarList(ctx context.Context, dictionaryType string) ([]*model.SysDictionaryDetails, error) {
	dictionaryDetail, err := rep.DictionariesModel.FindOneByType(ctx, dictionaryType)
	if err != nil {
		return nil, err
	}
	return rep.DictionariesDetailsModel.DictionaryDetailList(ctx, dictionaryDetail.Id)
}
