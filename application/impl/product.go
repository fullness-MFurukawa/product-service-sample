package impl

import (
	"context"
	"fmt"
	"product-service/application"
	"product-service/domain/product"

	rep "product-service/infrastructure/sqlboiler/product"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

/*
商品サービスインターフェースの実装
*/
type ProductServiceImpl struct {
	Prduct_Repository product.ProductRepositiry
}

// すべての商品を取得する
func (svr *ProductServiceImpl) Products() ([]product.Product, error) {
	ctx := context.Background()
	transaction, err := boil.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, application.NewServiceError(err.Error())
	}
	results, err := svr.Prduct_Repository.FindAll(ctx, transaction)
	if err != nil {
		return nil, err
	} else {
		return results, nil
	}
}

// キーワードで商品を検索する
func (svr *ProductServiceImpl) SearchByKeyword(keyword string) ([]product.Product, error) {
	ctx := context.Background()
	transaction, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, application.NewServiceError(err.Error())
	}
	results, err := svr.Prduct_Repository.FindByNameLike(ctx, transaction, keyword)
	if err != nil {
		return nil, err
	} else {
		if results == nil {
			return nil, application.NewServiceError(fmt.Sprintf("キーワード:[%s]が含まれる商品名が見つかりません。", keyword))
		}
		return results, nil
	}
}

// 新しい商品を追加する
func (svr *ProductServiceImpl) Add(product *product.Product) error {
	ctx := context.Background()
	transaction, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return application.NewServiceError(err.Error())
	}
	// 商品の存在チェック
	exist, err := svr.Prduct_Repository.Exist(ctx, transaction, product.ProductName())
	if err != nil {
		return err
	}
	if exist {
		return application.NewServiceError(fmt.Sprintf("商品:%sは既に登録済みです。", product.ProductName().Value()))
	}
	err = svr.Prduct_Repository.Create(ctx, transaction, product)
	if err != nil {
		return err
	}
	transaction.Commit()
	return nil
}

// 商品の内容を変更する
func (svr *ProductServiceImpl) Update(product *product.Product) error {
	ctx := context.Background()
	transaction, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return application.NewServiceError(err.Error())
	}
	result, err := svr.Prduct_Repository.UpdateById(ctx, transaction, product)
	if err != nil {
		return err
	}
	if !result { // 更新対象が見つからわない
		return application.NewServiceError(fmt.Sprintf("商品番号:%sの商品が見つかりませんでした。", product.ProductId().Value()))
	}
	transaction.Commit()
	return nil
}

// 商品を削除する
func (svr *ProductServiceImpl) Delete(productId *product.ProductId) error {
	ctx := context.Background()
	transaction, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return application.NewServiceError(err.Error())
	}
	result, err := svr.Prduct_Repository.DeleteById(ctx, transaction, productId)
	if err != nil {
		return err
	}
	if !result { // 削除対象が見つからわない
		return application.NewServiceError(fmt.Sprintf("商品番号:%sの商品が見つかりませんでした。", productId.Value()))
	}
	transaction.Commit()
	return nil
}

// コンストラクタ
func NewProductServiceImpl() application.ProductService {
	repository := rep.NewProductRepositoryImpl()
	return &ProductServiceImpl{
		Prduct_Repository: repository,
	}
}
