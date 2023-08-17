package services

import (
	"log"
	"meetme/be/actions/repositories"
	"meetme/be/actions/services/interfaces"

	repoInt "meetme/be/actions/repositories/interfaces"

	"meetme/be/utils"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) CreateUser(request interfaces.RegisterRequest) (interface{}, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	newUser := repoInt.User{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Birthday:  request.Birthday,
		Email:     request.Email,
		Password:  string(bytes),
	}
	result, err := s.userRepo.Create(newUser)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	response := utils.DataResponse{
		Data: &interfaces.RegisterResponse{
			ID:        result.ID,
			Firstname: result.Firstname,
			Lastname:  result.Lastname,
			Birthday:  result.Birthday,
			Email:     result.Email,
		},
		Message: "Create user success.",
	}

	return response, nil
}
