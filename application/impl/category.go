package impl

import (
	"context"
	"fmt"
	"product-service/application"
	"product-service/domain/category"
	rep "product-service/infrastructure/sqlboiler/category"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CategoryServiceImpl struct {
	repository category.CategoryRepository
}

func (svr *CategoryServiceImpl) Categories() ([]category.Category, error) {
	ctx := context.Background()
	transaction, err := boil.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, application.NewServiceError(err.Error())
	}
	results, err := svr.repository.FindAll(ctx, transaction)
	if err != nil {
		return nil, err
	} else {
		return results, nil
	}
}

func (svr *CategoryServiceImpl) Category(categoryId *category.CategoryId) (*category.Category, error) {
	ctx := context.Background()
	transaction, err := boil.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, application.NewServiceError(err.Error())
	}
	result, err := svr.repository.FindById(ctx, transaction, categoryId)
	if err != nil {
		return nil, err
	} else {
		if result == nil {
			return nil, application.NewServiceError(fmt.Sprintf("カテゴリ番号:[%s]に該当するカテゴリが見つかりません。", categoryId.Value()))
		}
		return result, nil
	}
}

func NewCategoryServiceImpl() application.CategoryService {
	repository := rep.NewCategoryRepositoryImpl()
	return &CategoryServiceImpl{repository: repository}
}

/*
インスタンス生成
fx用
*/
func NewCategoryServiceImplFx(repository category.CategoryRepository) application.CategoryService {
	return &CategoryServiceImpl{repository: repository}
}
