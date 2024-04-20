package interfaces

type BgRequest struct {
	Name   string `json:"name" validate:"required"`
	Assets string `json:"assets" validate:"required"`
	Price  int    `json:"price" validate:"required"`
}

type BgShopResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Assets  string `json:"img"`
	Price   int    `json:"price"`
	IsOwner bool   `json:"isOwner"`
}
type BgService interface {
	GetBgShops(string) (interface{}, error)
	AddBgShop(string, BgRequest) (interface{}, error)
}
