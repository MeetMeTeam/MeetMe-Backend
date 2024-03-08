package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"log"
	"meetme/be/actions/repositories"
	repoInt "meetme/be/actions/repositories/interfaces"
	"meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
)

type themeService struct {
	themeRepo     repositories.ThemeRepository
	userRepo      repositories.UserRepository
	inventoryRepo repositories.InventoryRepository
}

func NewThemeService(themeRepo repositories.ThemeRepository, userRepo repositories.UserRepository, inventoryRepo repositories.InventoryRepository) interfaces.ThemeService {
	return themeService{themeRepo: themeRepo, userRepo: userRepo, inventoryRepo: inventoryRepo}
}

func (s themeService) AddThemeShop(token string, request interfaces.ThemeCreateRequest) (interface{}, error) {
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
	newTheme := repoInt.Theme{
		Name:   request.Name,
		Price:  request.Price,
		Assets: request.Assets,
		Song:   request.Song,
	}
	resultTheme, err := s.themeRepo.CreateTheme(newTheme)
	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}

	return utils.DataResponse{
		Data: interfaces.ThemeCreateRequest{
			Name:   resultTheme.Name,
			Assets: resultTheme.Assets,
			Price:  resultTheme.Price,
			Song:   resultTheme.Song,
		},
		Message: "Create " + resultTheme.Name + " success.",
	}, nil
}

func (s themeService) GetThemeShops(token string) (interface{}, error) {

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

	themes, err := s.themeRepo.GetAllTheme()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.DataResponse{
				Data:    []int{},
				Message: "Get theme shop success.",
			}, nil
		}
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	themeResponses := []interfaces.ThemeResponse{}
	for _, theme := range themes {
		_, err := s.inventoryRepo.GetByUserIdAndItemId(user.ID, theme.ID)
		isOwner := false
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				isOwner = false
			}
		} else {
			isOwner = true
		}
		themeResponse := interfaces.ThemeResponse{
			ID:      theme.ID.Hex(),
			Name:    theme.Name,
			Assets:  theme.Assets,
			Price:   theme.Price,
			IsOwner: isOwner,
		}
		themeResponses = append(themeResponses, themeResponse)
	}

	response := utils.DataResponse{
		Data:    themeResponses,
		Message: "Get theme shop success.",
	}

	return response, nil
}
