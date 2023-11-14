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
	//ID        string `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  string `json:"birthday"`
	Email     string `json:"email"`
}

type ListUserResponse struct {
	ID        string `json:"id"`
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
	Refresh  string `json:"refreshToken"`
	Mail     string `json:"mail"`
	Username string `json:"username"`
	Id       string `json:"_id"`
	Image    string `json:"image"`
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
type UserService interface {
	GetUsers() (interface{}, error)
	//GetUserById(int) (interface{}, error)
	CreateUser(RegisterRequest) (interface{}, error)
	Login(Login) (interface{}, error)
	RefreshToken(string) (interface{}, error)
}
