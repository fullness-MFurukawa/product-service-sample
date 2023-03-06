package application

import "product-service/domain/product"

/*
商品サービスインターフェース
*/
type ProductService interface {
	// すべての商品を取得する
	Products() ([]product.Product, error)
	// キーワードで商品を検索する
	SearchByKeyword(keyword string) ([]product.Product, error)
	// 新しい商品を追加する
	Add(product *product.Product) error
	// 商品の内容を変更する
	Update(product *product.Product) error
	// 商品を削除する
	Delete(productId *product.ProductId) error
}
