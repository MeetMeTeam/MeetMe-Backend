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
type AvatarService interface {
	GetAvatarShops(string) (interface{}, error)
}
