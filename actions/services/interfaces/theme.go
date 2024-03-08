package interfaces

type ThemeResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Assets  string `json:"assets"`
	Price   int    `json:"price"`
	Song    string `json:"song"`
	IsOwner bool   `json:"isOwner"`
}

type ThemeCreateRequest struct {
	Name   string `json:"name" validate:"required"`
	Assets string `json:"assets" validate:"required"`
	Price  int    `json:"price" validate:"required"`
	Song   string `json:"song"`
}

type ThemeService interface {
	GetThemeShops(string) (interface{}, error)
	AddThemeShop(string, ThemeCreateRequest) (interface{}, error)
}
