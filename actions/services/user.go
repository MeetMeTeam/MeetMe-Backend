package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"log"
	"meetme/be/actions/repositories"
	"meetme/be/actions/services/interfaces"
	"meetme/be/config"
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
	userRepo      repositories.UserRepository
	inventoryRepo repositories.InventoryRepository
	avatarRepo    repositories.AvatarRepository
}

type jwtCustomClaims struct {
	Email     string `json:"email"`
	IsRefresh bool   `json:"isRefresh"`
	Admin     bool   `json:"admin"`
	jwt.RegisteredClaims
}

func NewUserService(userRepo repositories.UserRepository, inventoryRepo repositories.InventoryRepository, avatarRepo repositories.AvatarRepository) interfaces.UserService {
	return userService{userRepo: userRepo, inventoryRepo: inventoryRepo, avatarRepo: avatarRepo}
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
		Username:    request.Username,
		IsAdmin:     request.IsAdmin,
	}
	userResult, err := s.userRepo.Create(newUser)
	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}

	itemId, err := primitive.ObjectIDFromHex(request.CharacterId)
	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}

	inventResult, err := s.inventoryRepo.Create(userResult.ID, itemId, "avatar")
	userResult, err = s.userRepo.UpdateAvatarById(userResult.ID, inventResult.ID)
	response := utils.DataResponse{
		Data: &interfaces.RegisterResponse{
			Birthday: userResult.Birthday,
			Email:    userResult.Email,
			Username: userResult.Username,
		},
		Message: "Create user success.",
	}

	////send verify email
	//templateData := interfaces.TemplateEmailData{
	//	Username: userResult.Username,
	//	Email:    userResult.Email,
	//	Title:    "Verify Email",
	//	Button:   "Verify Your Email",
	//	URL:      os.Getenv("APP_URL"),
	//}
	//r := config.NewRequest([]string{userResult.Email}, "[meetmeplay] Verify Your Account", "")
	//err = r.ParseTemplate("verifyFile.html", templateData)
	//
	//if err != nil {
	//	return nil, errs.NewInternalError(err.Error())
	//}
	//
	//ok, err := r.SendEmail()
	//if err != nil || !ok {
	//	return nil, errs.NewInternalError(err.Error())
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
			user.Email,
			false,
			user.IsAdmin,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
			},
		}
		refreshClaims := &jwtCustomClaims{
			user.Email,
			true,
			user.IsAdmin,
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
				Coin:     user.Coin,
				IsAdmin:  user.IsAdmin,
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
		Message: "Get users success.[test automated deploy]",
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
		refreshClaims.IsAdmin,
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

func (s userService) ForgotPassword(mail interfaces.Email) (interface{}, error) {

	user, err := s.userRepo.GetByEmail(mail.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("This email address is not registered yet.")
		}
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	claims := &jwtCustomClaims{
		user.Email,
		false,
		user.IsAdmin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}

	//send reset email
	templateData := interfaces.TemplateEmailData{
		Username: user.Username,
		Email:    user.Email,
		Title:    "Reset Password",
		Button:   "Reset Your Password",
		URL:      os.Getenv("APP_URL") + "reset-password/" + t,
	}
	r := config.NewRequest([]string{user.Email}, "[meetmeplay] Reset Your Password", "")
	err = r.ParseTemplate("verifyFile.html", templateData)

	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}

	ok, err := r.SendEmail()
	if err != nil || !ok {
		return nil, errs.NewInternalError(err.Error())
	}

	return utils.ErrorResponse{Message: "Send mail for reset password success."}, nil
}
func (s userService) ResetPassword(token string, password interfaces.Password) (interface{}, error) {
	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password.Password), 14)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	result, err := s.userRepo.UpdatePasswordByEmail(email.Email, string(bytes))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewUnauthorizedError("User not found")
		}
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}

	response := utils.ErrorResponse{
		Message: "Change password of " + result.Email + " success.",
	}
	return response, nil
}

func (s userService) GetCoin(token string) (interface{}, error) {
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

	type Coin struct {
		Coin int `json:"coin"`
	}
	return utils.DataResponse{
		Data: Coin{
			Coin: user.Coin,
		},
		Message: "Get Coin of " + user.Email + " success.",
	}, nil
}

func (s userService) GetAvatars(token string, id string) (interface{}, error) {
	_, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return nil, errs.NewInternalError(err.Error())
	}
	user, err := s.userRepo.GetById(userId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	inventory, err := s.inventoryRepo.GetById(user.Inventory)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("Inventory not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	avatar, err := s.avatarRepo.GetById(inventory.Item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("Avatar not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	return utils.DataResponse{
		Data: interfaces.AvatarResponse{
			ID:      avatar.ID.Hex(),
			Name:    avatar.Name,
			Assets:  avatar.Assets,
			Preview: avatar.Preview,
		},
		Message: "Get avatar of " + user.Username + " success.",
	}, nil
}

func (s userService) ChangeAvatar(token string, itemId string) (interface{}, error) {
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

	id, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}

	inventory, err := s.inventoryRepo.GetByUserIdAndItemId(user.ID, id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("Inventory not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	if inventory == nil {
		return nil, errs.NewBadRequestError("This item is not exist in inventory.")
	} else if inventory.ID == user.Inventory {
		return nil, errs.NewBadRequestError("This item is current avatar.")
	}

	updateUser, err := s.userRepo.UpdateAvatarById(user.ID, inventory.ID)

	return utils.DataResponse{
		Data:    updateUser,
		Message: "Change avatar from " + user.Inventory.Hex() + " to " + updateUser.Inventory.Hex() + " success.",
	}, nil

}
