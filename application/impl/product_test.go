package impl

import (
	"fmt"
	"product-service/domain/category"
	"product-service/domain/product"
	"product-service/infrastructure/sqlboiler"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProducts(t *testing.T) {
	// データベース接続とConenction Poolの生成
	_ = sqlboiler.SqlBiolderInitDB{}.Init(nil)
	// ProductServiceの生成
	service := NewProductServiceImpl()
	results, err := service.Products()
	if err != nil {
		assert.Fail(t, err.Error())
	}
	for _, result := range results {
		fmt.Println(result)
	}
	assert.True(t, true)
}

func TestSearchByKeyword(t *testing.T) {
	// データベース接続とConenction Poolの生成
	_ = sqlboiler.SqlBiolderInitDB{}.Init(nil)
	// ProductServiceの生成
	service := NewProductServiceImpl()
	results, err := service.SearchByKeyword("%ペン%")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	for _, result := range results {
		fmt.Println(result)
	}
	assert.NotNil(t, results)
	assert.Nil(t, err)

	results, err = service.SearchByKeyword("%おいうえお%")
	t.Log(err.Error())
	assert.Error(t, err)
	assert.Equal(t, 0, len(results))
}

func TestAdd(t *testing.T) {
	// データベース接続とConenction Poolの生成
	_ = sqlboiler.SqlBiolderInitDB{}.Init(nil)
	// ProductServiceの生成
	service := NewProductServiceImpl()
	category, _ := category.BuildCategory("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
	new_product, _ := product.NewProduct("商品-ABC", uint32(200), category)
	err := service.Add(new_product)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	// データベース接続とConenction Poolの生成
	_ = sqlboiler.SqlBiolderInitDB{}.Init(nil)
	// ProductServiceの生成
	service := NewProductServiceImpl()
	category, _ := category.BuildCategory("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
	update_product, _ := product.BuildProduct("db262872-857c-4824-9591-f397d9d0cbcf", "商品-XYZ", uint32(300), category)
	err := service.Update(update_product)
	assert.Nil(t, err)

	update_product, _ = product.BuildProduct("db262872-857c-4824-9591-f397d9d0cbc0", "商品-XYZ", uint32(300), category)
	err = service.Update(update_product)
	t.Log(err.Error())
	assert.Error(t, err)
}

func TestDelete(t *testing.T) {
	// データベース接続とConenction Poolの生成
	_ = sqlboiler.SqlBiolderInitDB{}.Init(nil)
	// ProductServiceの生成
	service := NewProductServiceImpl()
	product_id, _ := product.NewProductId("db262872-857c-4824-9591-f397d9d0cbcf")
	err := service.Delete(product_id)
	assert.Nil(t, err)
	product_id, _ = product.NewProductId("db262872-857c-4824-9591-f397d9d0cbc0")
	err = service.Delete(product_id)
	t.Log(err.Error())
	assert.Error(t, err)
}
