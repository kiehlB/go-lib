package common

type AppError struct {
	Message interface{}
	Status  int
}

func NewAppError(status int, message interface{}) *AppError {
	return &AppError{
		Message: message,
		Status:  status,
	}
}