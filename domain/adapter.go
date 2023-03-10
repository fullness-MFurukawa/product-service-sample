package domain

/*
Entityと他のモデルを変換するAdapterインターフェース
*/
type EntityAdapter interface {
	// 任意のEntityから任意のModelへ変換する
	Convert(entity interface{}) (interface{}, error)
	// 任意のModeleから任意のEntityへ変換する
	Restore(model interface{}) (interface{}, error)
}
