package domain

/*
ドメインルール違反を表すエラー型
*/
type DomainError struct {
	message string
}

// エラーメッセージを提供する
func (e *DomainError) Error() string {
	return e.message
}

// コンストラクタ
func NewDomainError(message string) error {
	return &DomainError{message: message}
}
