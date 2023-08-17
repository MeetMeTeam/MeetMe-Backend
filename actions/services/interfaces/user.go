package interfaces

type RegisterRequest struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" validate:"required,email"`
	Birthday  string `json:"birthday"`
	Password  string `json:"password"`
}

type RegisterResponse struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  string `json:"birthday"`
	Email     string `json:"email"`
}

type UserService interface {
	//GetUsers() ([]RegisterResponse, error)
	// GetUserByLineId(string) (interface{}, error)
	CreateUser(RegisterRequest) (interface{}, error)
	// AddPoints(PointRequest, string) (interface{}, error)
	// EditUser(EditRequest, string) (interface{}, error)
}
