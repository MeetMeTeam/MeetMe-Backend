package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	themeRepo     repositories.ThemeRepository
}

func NewInventoryService(inventoryRepo repositories.InventoryRepository, userRepo repositories.UserRepository, avatarRepo repositories.AvatarRepository, themeRepo repositories.ThemeRepository) interfaces.InventoryService {
	return inventoryService{inventoryRepo: inventoryRepo, userRepo: userRepo, avatarRepo: avatarRepo, themeRepo: themeRepo}
}

func (s inventoryService) GetInventory(token string, itemType string) (interface{}, error) {
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

	inventories, err := s.inventoryRepo.GetByUserIdAndItemType(user.ID, itemType)
	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}
	if inventories == nil {
		return utils.DataResponse{
			Data:    []string{},
			Message: "Get inventory list of " + user.Username + " success.",
		}, nil
	}

	var response interface{}
	if itemType == "avatar" {
		avatarRes := []interfaces.AvatarResponse{}
		for _, inventory := range inventories {

			avatar, err := s.avatarRepo.GetById(inventory.Item)
			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {
					return nil, errs.NewBadRequestError("Avatar not found.")
				}
				return nil, errs.NewInternalError(err.Error())
			}

			avatarResponse := interfaces.AvatarResponse{
				ID:      avatar.ID.Hex(),
				Name:    avatar.Name,
				Assets:  avatar.Assets,
				Preview: avatar.Preview,
				Type:    avatar.Type,
			}
			response = append(avatarRes, avatarResponse)

		}
	} else if itemType == "theme" {
		themeRes := []interfaces.ThemeResponse{}
		for _, inventory := range inventories {

			theme, err := s.themeRepo.GetThemeById(inventory.Item)
			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {
					return nil, errs.NewBadRequestError("Avatar not found.")
				}
				return nil, errs.NewInternalError(err.Error())
			}

			themeResponse := interfaces.ThemeResponse{
				ID:     theme.ID.Hex(),
				Name:   theme.Name,
				Assets: theme.Assets,
				Price:  theme.Price,
				Song:   theme.Song,
			}
			response = append(themeRes, themeResponse)

		}

	}

	return utils.DataResponse{
		Data:    response,
		Message: "Get " + itemType + " inventory list of " + user.Username + " success.",
	}, nil
}

func (s inventoryService) AddItem(token string, id string, itemType string) (interface{}, error) {
	if id == "" {
		return nil, errs.NewBadRequestError("Please attach item_id to parameter path.")
	}
	if itemType == "" {
		return nil, errs.NewBadRequestError("Please attach item_type to parameter path.")
	}
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
	itemId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}
	updateCoin := 0
	if itemType == "avatar" {
		items, err := s.avatarRepo.GetById(itemId)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errs.NewBadRequestError("Avatar not found.")
			}
			return nil, errs.NewInternalError(err.Error())
		}

		isExist, err := s.inventoryRepo.GetByUserIdAndItemId(user.ID, itemId)
		if isExist != nil {
			return nil, errs.NewBadRequestError(items.Name + " is exist in your inventory.")
		}
		if items.Price > user.Coin {
			return nil, errs.NewBadRequestError("Your coin not enough.")
		}
		updateCoin = user.Coin - items.Price
	} else if itemType == "theme" {
		items, err := s.themeRepo.GetThemeById(itemId)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errs.NewBadRequestError("Theme not found.")
			}
			return nil, errs.NewInternalError(err.Error())
		}

		isExist, err := s.inventoryRepo.GetByUserIdAndItemId(user.ID, itemId)
		if isExist != nil {
			return nil, errs.NewBadRequestError(items.Name + " is exist in your inventory.")
		}
		if items.Price > user.Coin {
			return nil, errs.NewBadRequestError("Your coin not enough.")
		}
		updateCoin = user.Coin - items.Price
	} else {
		return nil, errs.NewBadRequestError("Check item type again.")
	}

	result, err := s.inventoryRepo.Create(user.ID, itemId, itemType)
	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}

	_, err = s.userRepo.UpdateCoinById(user.ID, updateCoin)
	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}

	return utils.DataResponse{
		Data:    result,
		Message: "Add item: " + result.Item.Hex() + " to inventory success.",
	}, nil
}
