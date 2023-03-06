package application

/*
業務処理エラーを表すエラー型
*/
type ServiceError struct {
	message string
}

// エラーメッセージを提供する
func (e *ServiceError) Error() string {
	return e.message
}

// コンストラクタ
func NewServiceError(message string) error {
	return &ServiceError{message: message}
}
