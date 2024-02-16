package interfaces

type AvatarResponse struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Assets  []string `json:"assets"`
	Preview string   `json:"preview"`
	Type    string   `json:"type"`
}
type AvatarShopResponse struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Assets  []string `json:"assets"`
	Preview string   `json:"preview"`
	Price   int      `json:"price"`
	IsOwner bool     `json:"isOwner"`
	Type    string   `json:"type"`
}

type AvatarRequest struct {
	Name    string   `json:"name"`
	Assets  []string `json:"assets"`
	Preview string   `json:"preview"`
	Price   int      `json:"price"`
	Type    string   `json:"type"`
}

type CreateResponse struct {
	Name    string   `json:"name"`
	Assets  []string `json:"assets"`
	Preview string   `json:"preview"`
	Price   int      `json:"price"`
	Type    string   `json:"type"`
}

type AvatarService interface {
	GetAvatarShops(string) (interface{}, error)
	AddAvatarShop(string, AvatarRequest) (interface{}, error)
}
