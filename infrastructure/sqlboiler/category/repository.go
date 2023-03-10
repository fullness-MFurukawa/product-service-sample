package category

import (
	"context"
	"database/sql"
	"product-service/domain"
	"product-service/domain/category"
	"product-service/infrastructure"
	"product-service/infrastructure/sqlboiler/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

/*
CategoryテーブルアクセスRepository
*/
type CategoryRepositoryImpl struct {
	adapter domain.EntityAdapter // Categoryデータ変換Adapter
}

/*
カテゴリを全件取得する
*/
func (rep *CategoryRepositoryImpl) FindAll(ctx context.Context, tran *sql.Tx) ([]category.Category, error) {
	results, err := models.Categories().All(ctx, tran)
	if err != nil {
		return nil, infrastructure.NewInternalError(err.Error())
	}
	var categories []category.Category
	for _, result := range results {
		category_inf, err := rep.adapter.Restore(result)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category_inf.(category.Category))
	}
	return categories, nil
}

/*
指定された番号のカテゴリを取得する
*/
func (rep *CategoryRepositoryImpl) FindById(ctx context.Context, tran *sql.Tx, id *category.CategoryId) (*category.Category, error) {
	result, err := models.Categories(qm.Where("category.obj_id = ?", id.Value())).One(ctx, tran)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		} else {
			return nil, infrastructure.NewInternalError(err.Error())
		}
	}
	category_inf, err := rep.adapter.Restore(result)
	if err != nil {
		return nil, err
	}
	category := category_inf.(category.Category)
	return &category, nil
}
func NewCategoryRepositoryImpl() category.CategoryRepository {
	adapter := NewCategoryAdapterImpl()
	return &CategoryRepositoryImpl{adapter: adapter}
}
