package category

import (
	"context"
	"fmt"
	"product-service/domain/category"
	"product-service/infrastructure/sqlboiler"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestFindAll(t *testing.T) {
	// データベース接続とConenction Poolの生成
	_ = sqlboiler.SqlBiolderInitDB{}.Init(nil)
	ctx := context.Background()
	transaction, tran_err := boil.BeginTx(ctx, nil)
	if tran_err != nil {
		assert.Fail(t, tran_err.Error())
	}
	defer transaction.Rollback()
	repository := NewCategoryRepositoryImpl()
	categories, err := repository.FindAll(ctx, transaction)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	for _, category := range categories {
		fmt.Println(category.String())
	}
	assert.Equal(t, 3, len(categories))
}

func TestFindById(t *testing.T) {
	id, err := category.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	// データベース接続とConenction Poolの生成
	_ = sqlboiler.SqlBiolderInitDB{}.Init(nil)
	ctx := context.Background()
	transaction, tran_err := boil.BeginTx(ctx, nil)
	if tran_err != nil {
		assert.Fail(t, tran_err.Error())
	}
	defer transaction.Rollback()
	repository := NewCategoryRepositoryImpl()
	result, r_err := repository.FindById(ctx, transaction, id)
	if r_err != nil {
		assert.Fail(t, err.Error())
	}
	fmt.Println(result)
	assert.True(t, true)

	// 存在しないカテゴリID
	id, err = category.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c1")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	result, r_err = repository.FindById(ctx, transaction, id)
	assert.Nil(t, result)
	assert.Nil(t, err)
}
