package application

import "product-service/domain/category"

/*
カテゴリサービスインターフェース
*/
type CategoryService interface {
	// 全件取得する
	Categories() ([]category.Category, error)
	// 指定されたカテゴリを取得する
	Category(categoryId *category.CategoryId) (*category.Category, error)
}
