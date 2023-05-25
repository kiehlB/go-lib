package authdto

type LoginResponse struct {
	ID       int
	Name     string
	Email    string
	Password bool
	IsAdmin  bool
	Token    string
}
