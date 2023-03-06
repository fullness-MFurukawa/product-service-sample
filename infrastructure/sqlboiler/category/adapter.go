package category

import (
	"product-service/domain"
	"product-service/domain/category"
	"product-service/infrastructure/sqlboiler/models"
)

type CategoryAdapterImpl struct{}

// Entity CategoryからsqlboilerのModel Categoryへ変換する
func (adapter *CategoryAdapterImpl) Convert(entity interface{}) (interface{}, error) {
	source, ok := entity.(category.Category)
	if !ok {
		return nil, domain.NewDomainError("指定されたEntityはCategoryではありません。")
	}
	category := models.Category{ObjID: source.CategoryId().Value(), Name: source.CategoryName().Value()}
	return category, nil
}

// sqlboilerのModel Categoryから任意のEntity Categoryへ変換する
func (adapter *CategoryAdapterImpl) Restore(model interface{}) (interface{}, error) {
	source, ok := model.(*models.Category)
	if !ok {
		return nil, domain.NewDomainError("指定されたmodelはCategoryではありません。")
	}
	category, err := category.BuildCategory(source.ObjID, source.Name)
	if err != nil {
		return nil, err
	}
	return *category, nil
}

// コンストラクタ
func NewCategoryAdapterImpl() domain.EntityAdapter {
	return &CategoryAdapterImpl{}
}
