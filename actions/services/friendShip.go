package services

import (
	"meetme/be/actions/repositories"
)

type friendShipService struct {
	friendRepo repositories.FriendshipRepository
	userRepo   repositories.UserRepository
}

func NewFriendShipService(friendRepo repositories.FriendshipRepository, userRepo repositories.UserRepository) friendShipService {
	return friendShipService{friendRepo: friendRepo, userRepo: userRepo}
}

//func (s friendShipService) GetFriendList(token string) (interface{}, error) {
//	email, err := utils.IsTokenValid(token)
//	if err != nil {
//		return nil, err
//	}
//	user, err := s.userRepo.GetByEmail(email)
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, errs.NewNotFoundError("User not found.")
//		}
//		return nil, errs.NewInternalError(err.Error())
//	}
//	result, err := s.friendRepo.GetFriendById(user.ID)
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, errs.NewNotFoundError("Friend list is empty.")
//		}
//		return nil, errs.NewInternalError(err.Error())
//	}
//
//	userResponses := []interfaces.RegisterResponse{}
//	for _, friend := range result {
//		id := 0
//		if user.ID == friend.UserId1 {
//			id = friend.UserId2
//		} else if user.ID == friend.UserId2 {
//			id = friend.UserId1
//		}
//		friendKub, err := s.userRepo.GetById(id)
//		if err != nil {
//			log.Println(err)
//			return nil, errs.NewInternalError(err.Error())
//		}
//		userResponse := interfaces.RegisterResponse{
//			ID:        friendKub.ID,
//			Username:  friendKub.Username,
//			Email:     friendKub.Email,
//			Firstname: friendKub.Firstname,
//			Lastname:  friendKub.Lastname,
//			Birthday:  friendKub.Birthday,
//		}
//		userResponses = append(userResponses, userResponse)
//	}
//
//	response := utils.DataResponse{
//		Data:    userResponses,
//		Message: "Get friend list success.",
//	}
//	return response, nil
//}
