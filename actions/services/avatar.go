package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"log"
	"meetme/be/actions/repositories"
	"meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
)

type avatarService struct {
	avatarRepo    repositories.AvatarRepository
	userRepo      repositories.UserRepository
	inventoryRepo repositories.InventoryRepository
}

func NewAvatarService(avatarRepo repositories.AvatarRepository, userRepo repositories.UserRepository, inventoryRepo repositories.InventoryRepository) interfaces.AvatarService {
	return avatarService{avatarRepo: avatarRepo, userRepo: userRepo, inventoryRepo: inventoryRepo}
}

func (s avatarService) GetAvatarShops(token string) (interface{}, error) {
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

	avatars, err := s.avatarRepo.GetAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.DataResponse{
				Data:    []int{},
				Message: "Get avatar shop success.",
			}, nil
		}
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	avatarResponses := []interfaces.AvatarShopResponse{}
	for _, avatar := range avatars {
		_, err := s.inventoryRepo.GetByUserIdAndItemId(user.ID, avatar.ID)
		isOwner := false
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				isOwner = false
			}
		} else {
			isOwner = true
		}
		avatarResponse := interfaces.AvatarShopResponse{
			ID:      avatar.ID.Hex(),
			Name:    avatar.Name,
			Assets:  avatar.Assets,
			Preview: avatar.Preview,
			Price:   avatar.Price,
			IsOwner: isOwner,
		}
		avatarResponses = append(avatarResponses, avatarResponse)
	}

	response := utils.DataResponse{
		Data:    avatarResponses,
		Message: "Get avatar shop success.",
	}

	return response, nil
}
