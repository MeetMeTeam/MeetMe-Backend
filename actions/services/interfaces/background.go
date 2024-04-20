package interfaces

type BgRequest struct {
	Name   string `json:"name" validate:"required"`
	Assets string `json:"assets" validate:"required"`
	Price  int    `json:"price" validate:"required"`
}

type BgService interface {
	//GetAvatarShops(string, string) (interface{}, error)
	AddBgShop(string, BgRequest) (interface{}, error)
}
