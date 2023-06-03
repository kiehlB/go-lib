package postsdto

type PostCreateRequest struct {
	Content string `json:"content"`
	UserID  int    `json:"userId"`
}
