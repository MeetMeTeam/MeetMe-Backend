package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"log"
	"meetme/be/actions/repositories"
	"meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"os"
	"strings"
	"time"

	repoInt "meetme/be/actions/repositories/interfaces"

	"meetme/be/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repositories.UserRepository
}

type jwtCustomClaims struct {
	Email     string `json:"email"`
	IsRefresh bool   `json:"isRefresh"`
	// Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func NewUserService(userRepo repositories.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) CreateUser(request interfaces.RegisterRequest) (interface{}, error) {
	isUsernameExist, _ := s.userRepo.GetByUsername(request.Username)

	if isUsernameExist != nil {
		return nil, errs.NewBadRequestError(request.Username + " is already exist.")
	}
	isEmailExist, _ := s.userRepo.GetByEmail(request.Email)
	if isEmailExist != nil {
		return nil, errs.NewBadRequestError(request.Email + " is already exist.")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}
	newUser := repoInt.User{
		DisplayName: request.DisplayName,
		Birthday:    request.Birthday,
		Email:       request.Email,
		Password:    string(bytes),
		Image:       request.Image,
		Username:    request.Username,
	}
	result, err := s.userRepo.Create(newUser)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	response := utils.DataResponse{
		Data: &interfaces.RegisterResponse{
			Birthday: result.Birthday,
			Email:    result.Email,
			Username: result.Username,
		},
		Message: "Create user success.",
	}

	//send verify email
	//templateData := struct {
	//	DisplayName string
	//	URL         string
	//}{
	//	DisplayName: request.DisplayName,
	//	URL:         "www.google.com",
	//}
	//r := config.NewRequest([]string{result.Email}, "Hello Junk!", "Hello, World!")
	//err = r.ParseTemplate("verifyFile.html", templateData)
	//if err == nil {
	//	ok, _ := r.SendEmail()
	//	fmt.Println(ok)
	//}
	return response, nil
}

func (s userService) Login(request interfaces.Login) (interface{}, error) {
	user, err := s.userRepo.GetByEmail(request.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewUnauthorizedError("Email or password incorrect.")
		}
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err == nil {

		claims := &jwtCustomClaims{
			request.Email,
			false,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
			},
		}
		refreshClaims := &jwtCustomClaims{
			request.Email,
			true,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

		t, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
		if err != nil {
			log.Println(err)
			return nil, errs.NewInternalError(err.Error())
		}
		r, err := refreshToken.SignedString([]byte(os.Getenv("APP_SECRET")))

		if err != nil {
			log.Println(err)
			return nil, errs.NewInternalError(err.Error())
		}
		response := interfaces.LoginResponse{

			UserDetails: interfaces.UserDetails{
				Token:    t,
				Refresh:  r,
				Mail:     user.Email,
				Username: user.Username,
				Id:       user.ID.Hex(),
				Image:    user.Image,
			},
		}
		return response, nil
	} else {
		return nil, errs.NewUnauthorizedError("Email or password incorrect.")
	}

}

func (s userService) GetUsers() (interface{}, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.DataResponse{
				Data:    []int{},
				Message: "Get users success.",
			}, nil
		}
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	userResponses := []interfaces.ListUserResponse{}
	for _, user := range users {
		userResponse := interfaces.ListUserResponse{
			ID:          user.ID.Hex(),
			DisplayName: user.DisplayName,
			Email:       user.Email,
			Birthday:    user.Birthday,
			Username:    user.Username,
		}
		userResponses = append(userResponses, userResponse)
	}

	response := utils.DataResponse{
		Data:    userResponses,
		Message: "Get users success.",
	}

	return response, nil
}

func (s userService) RefreshToken(refreshToken string) (interface{}, error) {
	refreshClaims, err := utils.IsTokenValid(refreshToken)
	if err != nil {
		return nil, err
	}
	if refreshClaims.IsRefresh == false {
		return nil, errs.NewInternalError("It's not refresh token.")
	}
	claims := &jwtCustomClaims{
		refreshClaims.Email,
		false,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	response := interfaces.TokenResponse{
		AccessToken:  t,
		RefreshToken: strings.Trim(refreshToken, "Bearer "),
	}

	return response, nil
}
