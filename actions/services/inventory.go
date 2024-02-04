package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"meetme/be/actions/repositories"
	"meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
)

type inventoryService struct {
	inventoryRepo repositories.InventoryRepository
	userRepo      repositories.UserRepository
	avatarRepo    repositories.AvatarRepository
}

func NewInventoryService(inventoryRepo repositories.InventoryRepository, userRepo repositories.UserRepository, avatarRepo repositories.AvatarRepository) interfaces.InventoryService {
	return inventoryService{inventoryRepo: inventoryRepo, userRepo: userRepo, avatarRepo: avatarRepo}
}

func (s inventoryService) GetInventory(token string) (interface{}, error) {
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

	inventories, err := s.inventoryRepo.GetByUserId(user.ID)
	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}
	if inventories == nil {
		return utils.DataResponse{
			Data:    []string{},
			Message: "Get inventory list of " + user.Username + " success.",
		}, nil
	}

	response := []interfaces.AvatarResponse{}
	for _, inventory := range inventories {
		if inventory.Type == "avatar" {
			avatar, err := s.avatarRepo.GetById(inventory.Item)
			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {
					return nil, errs.NewBadRequestError("Avatar not found.")
				}
				return nil, errs.NewInternalError(err.Error())
			}

			avatarResponse := interfaces.AvatarResponse{
				Name:    avatar.Name,
				Assets:  avatar.Assets,
				Preview: avatar.Preview,
			}
			response = append(response, avatarResponse)
		}

	}
	return utils.DataResponse{
		Data:    response,
		Message: "Get inventory list of " + user.Username + " success.",
	}, nil
}
