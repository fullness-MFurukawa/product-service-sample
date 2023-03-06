package category

import (
	"fmt"
	"product-service/domain/category"
	"product-service/domain/product"
	"product-service/infrastructure/sqlboiler/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	category, err := category.BuildCategory("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	adapter := NewCategoryAdapterImpl()
	model, _ := adapter.Convert(category)
	fmt.Println(model)
	assert.True(t, true)

	id := "ac413f22-0cf1-490a-9635-7e9ca810e544"
	product, err := product.BuildProduct(id, "商品-ABC", uint32(300), nil)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	_, err = adapter.Convert(product)
	assert.Equal(t, "指定されたEntityはCategoryではありません。", err.Error())
}

func TestRestore(t *testing.T) {
	model := &models.Category{ID: 0, ObjID: "b1524011-b6af-417e-8bf2-f449dd58b5c0", Name: "文房具"}
	adapter := NewCategoryAdapterImpl()
	category, err := adapter.Restore(model)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	fmt.Println(category)
	assert.True(t, true)

	p_model := models.Product{ID: 0, ObjID: "b1524011-b6af-417e-8bf2-f449dd58b5c0",
		Name: "商品-ABC", Price: 100}
	_, err = adapter.Restore(p_model)
	assert.Equal(t, "指定されたmodelはCategoryではありません。", err.Error())
}
