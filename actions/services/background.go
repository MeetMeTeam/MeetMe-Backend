package services

import (
	"meetme/be/actions/repositories"
	repoInt "meetme/be/actions/repositories/interfaces"
	"meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
)

type bgService struct {
	bgRepo        repositories.BgRepository
	userRepo      repositories.UserRepository
	inventoryRepo repositories.InventoryRepository
}

func NewBgService(bgRepo repositories.BgRepository, userRepo repositories.UserRepository, inventoryRepo repositories.InventoryRepository) interfaces.BgService {
	return bgService{bgRepo: bgRepo, userRepo: userRepo, inventoryRepo: inventoryRepo}
}

func (s bgService) AddBgShop(token string, request interfaces.BgRequest) (interface{}, error) {
	result, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	if result.IsAdmin == false {
		return nil, errs.NewForbiddenError("You don't have permission.")
	}
	if request.Price < 0 {
		return nil, errs.NewForbiddenError("The price must not be negative.")
	}
	newBg := repoInt.Background{
		Name:   request.Name,
		Price:  request.Price,
		Assets: request.Assets,
	}
	resultBg, err := s.bgRepo.Create(newBg)
	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}

	return utils.DataResponse{
		Data: interfaces.BgRequest{
			Name:   resultBg.Name,
			Assets: resultBg.Assets,
			Price:  resultBg.Price,
		},
		Message: "Create " + resultBg.Name + " success.",
	}, nil
}
