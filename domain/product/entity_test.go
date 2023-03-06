package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("商品-ABC", uint32(200), nil)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.NotNil(t, product)
}
func TestBuildProduct(t *testing.T) {
	id := "ac413f22-0cf1-490a-9635-7e9ca810e544"
	product, err := BuildProduct(id, "商品-ABC", uint32(300), nil)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.NotNil(t, product)
}
func TestEquals(t *testing.T) {
	id := "ac413f22-0cf1-490a-9635-7e9ca810e544"
	product1, err := BuildProduct(id, "商品-ABC", uint32(300), nil)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	product2, err := BuildProduct(id, "商品-ABC", uint32(300), nil)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	result, _ := product1.Equals(product2)
	assert.Equal(t, result, true)

	id = "ac413f22-0cf1-490a-9635-7e9ca810e545"
	product3, err := BuildProduct(id, "商品-ABC", uint32(300), nil)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	result, _ = product1.Equals(product3)
	assert.NotEqual(t, result, true)
}
