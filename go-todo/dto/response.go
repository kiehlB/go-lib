package dto

type SuccessResponse struct {
	Status  int
	Message string
	Data    interface{}
}

type ErrorResponse struct {
	Status  int
	Message string
}
