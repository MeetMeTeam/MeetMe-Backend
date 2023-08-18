package interfaces

type RegisterRequest struct {
	Firstname string `json:"firstname" validate:"required" example:"Kanyapat"`
	Lastname  string `json:"lastname" example:"Wittayamitkul"`
	Email     string `json:"email" validate:"required,email" example:"winner@mail.com"`
	Birthday  string `json:"birthday" example:"2023-08-12"`
	Password  string `json:"password" example:"winner"`
}

type RegisterResponse struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  string `json:"birthday"`
	Email     string `json:"email"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type UserService interface {
	//GetUsers() ([]RegisterResponse, error)
	// GetUserByLineId(string) (interface{}, error)
	CreateUser(RegisterRequest) (interface{}, error)
	Login(Login) (interface{}, error)
	// AddPoints(PointRequest, string) (interface{}, error)
	// EditUser(EditRequest, string) (interface{}, error)
}
