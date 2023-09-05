package interfaces

type RegisterRequest struct {
	Firstname string `json:"firstname" validate:"required" example:"Kanyapat"`
	Lastname  string `json:"lastname" example:"Wittayamitkul"`
	Username  string `json:"username" example:"winnerkypt"`
	Email     string `json:"email" validate:"required,email" example:"winner@mail.com"`
	Birthday  string `json:"birthday" example:"2023-08-12"`
	Password  string `json:"password" example:"winner" validate:"required"`
	Image     string `json:"image" validate:"required"`
}

type RegisterResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  string `json:"birthday"`
	Email     string `json:"email"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email" example:"winner@mail.com"`
	Password string `json:"password" validate:"required" example:"winner"`
}

type LoginResponse struct {
	UserDetails interface{} `json:"userDetails"`
}

type UserDetails struct {
	Token    string `json:"token"`
	Mail     string `json:"mail"`
	Username string `json:"username"`
	Id       int    `json:"_id"`
}
type UserService interface {
	GetUsers() (interface{}, error)
	//GetUserById(int) (interface{}, error)
	CreateUser(RegisterRequest) (interface{}, error)
	Login(Login) (interface{}, error)
	// AddPoints(PointRequest, string) (interface{}, error)
	// EditUser(EditRequest, string) (interface{}, error)
}
