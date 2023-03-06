package category

import (
	"context"
	"database/sql"
)

/*
カテゴリをアクセスするRepositoryインターフェース
*/
type CategoryRepository interface {
	FindAll(ctx context.Context, tran *sql.Tx) ([]Category, error)
	FindById(ctx context.Context, tran *sql.Tx, id *CategoryId) (*Category, error)
}
