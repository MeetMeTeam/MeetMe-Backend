package interfaces

type RegisterRequest struct {
	Username    string `json:"username" example:"winnerkypt"`
	DisplayName string `json:"displayName" example:"winnerkypt"`
	Email       string `json:"email" validate:"required,email" example:"winner@mail.com"`
	Birthday    string `json:"birthday" example:"2023-08-12"`
	Password    string `json:"password" example:"winner" validate:"required"`
	Image       string `json:"image" validate:"required"`
}

type RegisterResponse struct {
	//ID        string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Birthday    string `json:"birthday"`
	Email       string `json:"email"`
}

type ListUserResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Birthday    string `json:"birthday"`
	Email       string `json:"email"`
	Image       string `json:"image"`
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
	Coin     int    `json:"coin"`
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Email struct {
	Email string `json:"email" validate:"required,email"`
}

type Password struct {
	Password string `json:"password" validate:"required"`
}

type UserService interface {
	GetUsers() (interface{}, error)
	//GetUserById(int) (interface{}, error)
	CreateUser(RegisterRequest) (interface{}, error)
	Login(Login) (interface{}, error)
	RefreshToken(string) (interface{}, error)
	ForgotPassword(Email) (interface{}, error)
	ResetPassword(string, Password) (interface{}, error)
	GetCoin(string) (interface{}, error)
	GetAvatars(string, string) (interface{}, error)
}
