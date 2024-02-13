package interfaces

type AvatarResponse struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Assets  []string `json:"assets"`
	Preview string   `json:"preview"`
}
type AvatarShopResponse struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Assets  []string `json:"assets"`
	Preview string   `json:"preview"`
	Price   int      `json:"price"`
	IsOwner bool     `json:"isOwner"`
}

type AvatarRequest struct {
	Name    string   `json:"name"`
	Assets  []string `json:"assets"`
	Preview string   `json:"preview"`
	Price   int      `json:"price"`
}

type CreateResponse struct {
	Name    string   `json:"name"`
	Assets  []string `json:"assets"`
	Preview string   `json:"preview"`
	Price   int      `json:"price"`
}

type AvatarService interface {
	GetAvatarShops(string) (interface{}, error)
	AddAvatarShop(string, AvatarRequest) (interface{}, error)
}
