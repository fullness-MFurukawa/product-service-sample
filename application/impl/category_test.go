package impl

import (
	"product-service/domain/category"
	"product-service/infrastructure/sqlboiler"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategories(t *testing.T) {
	// データベース接続とConenction Poolの生成
	_ = sqlboiler.SqlBiolderInitDB{}.Init(nil)
	// CategoryServiceの生成
	service := NewCategoryServiceImpl()
	results, err := service.Categories()
	for _, result := range results {
		t.Log(result.String())
	}
	assert.NotNil(t, results)
	assert.Nil(t, err)
}

func TestCategory(t *testing.T) {
	// データベース接続とConenction Poolの生成
	_ = sqlboiler.SqlBiolderInitDB{}.Init(nil)
	// CategoryServiceの生成
	service := NewCategoryServiceImpl()

	category_id, err := category.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	result, err := service.Category(category_id)
	t.Log(result.String())
	assert.Nil(t, err)
	assert.NotNil(t, result)

	category_id, err = category.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c1")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	result, err = service.Category(category_id)
	t.Log(err.Error())
	assert.NotNil(t, err)
	assert.Nil(t, result)
}
