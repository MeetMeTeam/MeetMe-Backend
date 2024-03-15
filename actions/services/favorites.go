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

type favoriteService struct {
	userRepo repositories.UserRepository
	favRepo  repositories.FavoriteRepository
}

func NewFavoriteService(userRepo repositories.UserRepository, favRepo repositories.FavoriteRepository) interfaces.FavoriteService {
	return favoriteService{userRepo: userRepo, favRepo: favRepo}
}

func (s favoriteService) FavUser(token string, userId string) (interface{}, error) {
	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	giver, err := s.userRepo.GetByEmail(email.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	receiverId, err := primitive.ObjectIDFromHex(userId)

	if giver.ID == receiverId {
		return nil, errs.NewBadRequestError("You cannot like yourself.")
	} else {
		receiver, err := s.userRepo.GetById(receiverId)

		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errs.NewBadRequestError("Receiver ID: " + userId + " not found.")
			}
			return nil, errs.NewInternalError(err.Error())
		} else {
			result, _ := s.favRepo.GetByGiverAndReceiver(giver.ID, receiver.ID)
			if result != nil {
				return nil, errs.NewBadRequestError("You already like " + receiver.Username)
			} else {
				_, err := s.favRepo.AddFav(giver.ID, receiverId)
				if err != nil {
					return nil, errs.NewInternalError(err.Error())
				}
			}
		}
	}

	return utils.ErrorResponse{
		Message: "You press like " + receiverId.Hex() + " success.",
	}, nil
}

func (s favoriteService) UnFavUser(token string, userId string) (interface{}, error) {
	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	giver, err := s.userRepo.GetByEmail(email.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	receiverId, err := primitive.ObjectIDFromHex(userId)

	if giver.ID == receiverId {
		return nil, errs.NewBadRequestError("You cannot unlike yourself.")
	} else {

		_, err = s.favRepo.GetByGiverAndReceiver(giver.ID, receiverId)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errs.NewBadRequestError("You have not like yet!")
			}
			return nil, errs.NewInternalError(err.Error())
		} else {
			err = s.favRepo.DeleteFav(giver.ID, receiverId)
			if err != nil {
				return nil, errs.NewInternalError(err.Error())
			}
		}

	}

	return utils.ErrorResponse{
		Message: "You remove like " + receiverId.Hex() + " success.",
	}, nil
}

func (s favoriteService) GetCountFav(token string) (interface{}, error) {
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

	num, err := s.favRepo.CountFav(user.ID)
	if err != nil {
		return nil, err
	}

	type Fav struct {
		CountFav int `json:"countFav"`
	}
	return utils.DataResponse{
		Data: Fav{
			CountFav: num,
		},
		Message: "Count favorite of " + user.Username + " success.",
	}, nil
}
