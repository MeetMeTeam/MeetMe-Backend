package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
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

func (s bgService) GetBgShops(token string) (interface{}, error) {

	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.GetByEmail(email.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	bgs, err := s.bgRepo.GetAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.DataResponse{
				Data:    []int{},
				Message: "Get background shop success.",
			}, nil
		}
		return nil, errs.NewInternalError(err.Error())
	}

	bgResponses := []interfaces.BgShopResponse{}
	for _, bg := range bgs {
		_, err := s.inventoryRepo.GetByUserIdAndItemId(user.ID, bg.ID)
		isOwner := false
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				isOwner = false
			}
		} else {
			isOwner = true
		}
		bgResponse := interfaces.BgShopResponse{
			ID:      bg.ID.Hex(),
			Name:    bg.Name,
			Assets:  bg.Assets,
			Price:   bg.Price,
			IsOwner: isOwner,
		}
		bgResponses = append(bgResponses, bgResponse)
	}

	response := utils.DataResponse{
		Data:    bgResponses,
		Message: "Get background shop success.",
	}

	return response, nil
}
