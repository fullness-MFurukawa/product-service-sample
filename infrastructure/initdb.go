package infrastructure

/*
データベース初期化処理を表すインターフェース
*/
type InitDB interface {
	Init(interface{}) interface{}
}
