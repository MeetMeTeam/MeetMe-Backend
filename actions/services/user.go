package services

import (
	"errors"
	"log"
	"meetme/be/actions/repositories"
	"meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"time"

	repoInt "meetme/be/actions/repositories/interfaces"

	"meetme/be/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	userRepo repositories.UserRepository
}

type jwtCustomClaims struct {
	Name string `json:"name"`
	// Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func NewUserService(userRepo repositories.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) CreateUser(request interfaces.RegisterRequest) (interface{}, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}
	newUser := repoInt.User{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Birthday:  request.Birthday,
		Email:     request.Email,
		Password:  string(bytes),
		Image:     request.Image,
	}
	result, err := s.userRepo.Create(newUser)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
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

func (s userService) Login(request interfaces.Login) (interface{}, error) {
	user, err := s.userRepo.GetByEmail(request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err == nil {

		claims := &jwtCustomClaims{
			request.Email,
			// true,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("Hk89LSUPn3r4JDL@#@#$LJJKJDP00-.KJOS"))
		if err != nil {
			log.Println(err)
			return nil, errs.NewInternalError(err.Error())
		}
		response := interfaces.LoginResponse{
			Token:    t,
			Mail:     user.Email,
			Username: user.Firstname,
			Id:       user.ID,
		}
		return response, nil
	} else {
		response := utils.ErrorResponse{
			Message: "Email or password incorrect.",
		}
		return response, nil
	}

}

func (s userService) GetUsers() (interface{}, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("User not found.")
		}
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	userResponses := []interfaces.RegisterResponse{}
	for _, user := range users {
		userResponse := interfaces.RegisterResponse{
			ID:        user.ID,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			Birthday:  user.Birthday,
		}
		userResponses = append(userResponses, userResponse)
	}

	response := utils.DataResponse{
		Data:    userResponses,
		Message: "Get users success.",
	}
	return response, nil
}
