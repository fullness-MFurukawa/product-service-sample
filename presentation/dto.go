package presentation

/*
カテゴリ情報を扱うDTO
*/
type CategoryDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

/*
商品情報を扱うDTO
*/
type ProductDto struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

/*
商品情報一覧
*/
type ProductsDto struct {
	Products []*ProductDto `json:"products"`
}

/*
エラー情報
*/
type ErrorDto struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
}

/*
サンプルDTO
*/
type SampleDto struct {
	Message string `json:"message"`
}
