package services

import (
	"errors"
	"meetme/be/actions/repositories"
	"meetme/be/actions/services/interfaces"
	"meetme/be/config"
	"meetme/be/errs"
	"os"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"

	repoInt "meetme/be/actions/repositories/interfaces"

	"meetme/be/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo      repositories.UserRepository
	inventoryRepo repositories.InventoryRepository
	avatarRepo    repositories.AvatarRepository
	favRepo       repositories.FavoriteRepository
	bgRepo        repositories.BgRepository
}

type jwtCustomClaims struct {
	Email     string `json:"email"`
	IsRefresh bool   `json:"isRefresh"`
	Admin     bool   `json:"admin"`
	jwt.RegisteredClaims
}

func NewUserService(userRepo repositories.UserRepository, inventoryRepo repositories.InventoryRepository, avatarRepo repositories.AvatarRepository, favRepo repositories.FavoriteRepository, bgRepo repositories.BgRepository) interfaces.UserService {
	return userService{userRepo: userRepo, inventoryRepo: inventoryRepo, avatarRepo: avatarRepo, favRepo: favRepo, bgRepo: bgRepo}
}

func (s userService) CreateUser(request interfaces.RegisterRequest) (interface{}, error) {

	isEmailExist, err := s.userRepo.GetByEmail(request.Email)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewUnauthorizedError("Please verify email.")
		}

		return nil, errs.NewInternalError(err.Error())
	}

	if request.OTP == isEmailExist.Code && request.RefCode == isEmailExist.RefCode {
		if isEmailExist.IsVerify == true {
			return nil, errs.NewBadRequestError("This email is already verified.")
		} else {
			if isEmailExist.ExpiredAt.Before(time.Now()) {
				return nil, errs.NewBadRequestError("OTP expired, request it again")
			}
		}

		isUsernameExist, _ := s.userRepo.GetByUsername(request.Username)
		if isUsernameExist != nil {
			return nil, errs.NewBadRequestError(request.Username + " is already exist.")
		}

		bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
		if err != nil {

			return nil, errs.NewInternalError(err.Error())
		}

		newUser := repoInt.User{
			DisplayName: request.DisplayName,
			Birthday:    request.Birthday,
			Email:       request.Email,
			Password:    string(bytes),
			Username:    request.Username,
			IsAdmin:     request.IsAdmin,
			IsVerify:    true,
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

		return response, nil
	} else {
		return nil, errs.NewBadRequestError("OTP is incorrect!")
	}

}

func (s userService) Login(request interfaces.Login) (interface{}, error) {
	var user *repoInt.UserResponse
	var err error
	if strings.Contains(request.Email, "@") {
		user, err = s.userRepo.GetByEmail(request.Email)
	} else {
		user, err = s.userRepo.GetByUsername(request.Email)
	}

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewUnauthorizedError("Email or password incorrect.")
		}

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

			return nil, errs.NewInternalError(err.Error())
		}
		r, err := refreshToken.SignedString([]byte(os.Getenv("APP_SECRET")))

		if err != nil {

			return nil, errs.NewInternalError(err.Error())
		}

		num, err := s.favRepo.CountFav(user.ID)
		if err != nil {
			return nil, errs.NewInternalError(err.Error())
		}
		response := interfaces.LoginResponse{

			UserDetails: interfaces.UserDetails{
				Token:       t,
				Refresh:     r,
				Mail:        user.Email,
				Username:    user.Username,
				DisplayName: user.DisplayName,
				Id:          user.ID.Hex(),
				Coin:        user.Coin,
				IsAdmin:     user.IsAdmin,
				CountFav:    num,
				Bio:         user.Bio,
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
		refreshClaims.IsAdmin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil {

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
		URL:      os.Getenv("APP_URL") + "/reset-password/" + t,
		Web:      os.Getenv("APP_URL"),
	}
	r := config.NewRequest([]string{user.Email}, "[meetmefun] Reset Your Password", "")
	err = r.ParseTemplate("reset-password.html", templateData)

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

		return nil, errs.NewInternalError(err.Error())
	}

	result, err := s.userRepo.UpdatePasswordByEmail(email.Email, string(bytes))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewUnauthorizedError("User not found")
		}

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

func (s userService) EditUser(request interfaces.EditUserRequest, token string) (interface{}, error) {
	email, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}

	if reflect.DeepEqual(request, interfaces.EditUserRequest{}) {
		return utils.ErrorResponse{Message: "Not enough data to update information."}, nil
	}

	user, err := s.userRepo.GetByEmail(email.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	//if *request.DisplayName == user.DisplayName && *request.Username == user.Username && *request.Bio == user.Bio {
	//	return utils.ErrorResponse{Message: "Your request data is not change."}, nil
	//}
	var updateUser *repoInt.UserResponse

	//if (request.Username != nil && user.Username != "") || (request.Bio != nil && user.Bio != "") || (request.DisplayName != nil && user.DisplayName != "") {
	//	fmt.Println(*request.Bio == user.Bio)
	//	if *request.DisplayName == user.DisplayName && *request.Username == user.Username && *request.Bio == user.Bio {
	//		return utils.ErrorResponse{Message: "Your request data is not change."}, nil
	//	} else {
	//		if request.Bio != nil {
	//			updateUser, err = s.userRepo.UpdateBioByEmail(user.Email, *request.Bio)
	//			if err != nil {
	//				return nil, errs.NewInternalError(err.Error())
	//			}
	//		}
	//
	//		if request.DisplayName != nil {
	//			updateUser, err = s.userRepo.UpdateDisplayNameByEmail(user.Email, *request.DisplayName)
	//			if err != nil {
	//				return nil, errs.NewInternalError(err.Error())
	//			}
	//		}
	//
	//		if request.Username != nil {
	//			updateUser, err = s.userRepo.UpdateUsernameByEmail(user.Email, *request.Username)
	//			if err != nil {
	//				return nil, errs.NewInternalError(err.Error())
	//			}
	//		}
	//
	//	}
	//}

	if request.Bio != nil {
		updateUser, err = s.userRepo.UpdateBioByEmail(user.Email, *request.Bio)
		if err != nil {
			return nil, errs.NewInternalError(err.Error())
		}
	}

	if request.DisplayName != nil {
		updateUser, err = s.userRepo.UpdateDisplayNameByEmail(user.Email, *request.DisplayName)
		if err != nil {
			return nil, errs.NewInternalError(err.Error())
		}
	}

	if request.Username != nil {
		updateUser, err = s.userRepo.UpdateUsernameByEmail(user.Email, *request.Username)
		if err != nil {
			return nil, errs.NewInternalError(err.Error())
		}
	}

	if request.Social != nil {
		//if user.Social == nil {
		updateUser, err = s.userRepo.AddSocial(user.Email, request.Social)
		//} else {
		//	updateUser, err = s.userRepo.UpdateSocialByEmail(user.Email, request.Social)
		//}

		if err != nil {
			return nil, errs.NewInternalError(err.Error())
		}

	}

	//
	//if user.Bio == "" {
	//	if request.Bio != nil {
	//		updateUser, err = s.userRepo.UpdateBioByEmail(user.Email, *request.Bio)
	//		if err != nil {
	//			return nil, errs.NewInternalError(err.Error())
	//		}
	//	}
	//} else {
	//	if *request.Bio == user.Bio {
	//		return utils.ErrorResponse{Message: "Your request data is not change."}, nil
	//	} else {
	//		updateUser, err = s.userRepo.UpdateBioByEmail(user.Email, *request.Bio)
	//		if err != nil {
	//			return nil, errs.NewInternalError(err.Error())
	//		}
	//	}
	//}
	//
	//if request.DisplayName != nil {
	//	if *request.DisplayName != user.DisplayName {
	//		updateUser, err = s.userRepo.UpdateDisplayNameByEmail(user.Email, *request.DisplayName)
	//		if err != nil {
	//			return nil, errs.NewInternalError(err.Error())
	//		}
	//	} else {
	//		return utils.ErrorResponse{Message: "Your request data is not change."}, nil
	//	}
	//
	//}
	//
	//if request.Username != nil {
	//	if *request.Username != user.Username {
	//		updateUser, err = s.userRepo.UpdateUsernameByEmail(user.Email, *request.Username)
	//		if err != nil {
	//			return nil, errs.NewInternalError(err.Error())
	//		}
	//	} else {
	//		return utils.ErrorResponse{Message: "Your request data is not change."}, nil
	//	}
	//
	//}

	result := interfaces.ListUserResponse{
		ID:          updateUser.ID.Hex(),
		Username:    updateUser.Username,
		DisplayName: updateUser.DisplayName,
		Email:       updateUser.Email,
		Bio:         updateUser.Bio,
		Birthday:    updateUser.Birthday,
		Social:      updateUser.Social,
	}
	return utils.DataResponse{
		Data:    result,
		Message: "Edit user information success.",
	}, nil
}

func (s userService) VerifyEmail(email interfaces.Email) (interface{}, error) {
	var verifyData *repoInt.Mail
	var err error
	isEmailExist, _ := s.userRepo.GetByEmail(email.Email)

	if isEmailExist == nil {
		verifyData, err = s.userRepo.CreateVerifyMail(email.Email, utils.EncodeToString(6, "int"), utils.EncodeToString(6, "string"), time.Now().Add(time.Minute*10))
		if err != nil {
			return nil, errs.NewInternalError(err.Error())
		}

	} else {
		if isEmailExist.IsVerify == true {
			return nil, errs.NewBadRequestError(email.Email + " is already exist.")
		} else {
			verifyData, err = s.userRepo.UpdateVerifyMailCode(email.Email, utils.EncodeToString(6, "int"), utils.EncodeToString(6, "string"), time.Now().Add(time.Minute*10))
			if err != nil {
				return nil, errs.NewInternalError(err.Error())
			}
		}
	}

	//send verify email
	templateData := interfaces.TemplateEmailData{
		Username: email.Email,
		Email:    email.Email,
		Title:    "Verify and SignIn",
		Button:   "Verify Your Email",
		OTP:      verifyData.Code,
		RefCode:  verifyData.RefCode,
		Web:      os.Getenv("APP_URL"),
	}
	r := config.NewRequest([]string{email.Email}, "[meetmefun] Verify Your Account", "")
	err = r.ParseTemplate("verify-mail.html", templateData)

	if err != nil {
		return nil, errs.NewInternalError(err.Error())
	}

	ok, err := r.SendEmail()
	if err != nil || !ok {
		return nil, errs.NewInternalError(err.Error())
	}

	return utils.DataResponse{
		Data: interfaces.OTPResponse{
			Email:     verifyData.Email,
			RefCode:   verifyData.RefCode,
			ExpiredAt: verifyData.ExpiredAt,
		},
		Message: "Send mail to verify success.",
	}, nil

}

func (s userService) ChangeBackground(token string, itemId string) (interface{}, error) {
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

	var updateDefault *repoInt.InventoryResponse
	inventory, err := s.inventoryRepo.GetByTypeAndUserIdAndDefault("bg", user.ID, true)
	if inventory.Item == id {
		return nil, errs.NewBadRequestError("This item id is already default.")
	}
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {

			updateDefault, err = s.inventoryRepo.UpdateDefaultByUserIdAndItemIdAndType(user.ID, id, "bg", true)
			if err != nil {
				return nil, errs.NewInternalError(err.Error())
			}

		}
		return nil, errs.NewInternalError(err.Error())
	} else {
		updateDefault, err = s.inventoryRepo.UpdateDefaultByUserIdAndItemIdAndType(user.ID, inventory.Item, "bg", false)
		if err != nil {
			return nil, errs.NewInternalError(err.Error())
		}

		updateDefault, err = s.inventoryRepo.UpdateDefaultByUserIdAndItemIdAndType(user.ID, id, "bg", true)
		if err != nil {
			return nil, errs.NewInternalError(err.Error())
		}

	}

	return utils.DataResponse{
		Data: interfaces.ChangeBgResponse{
			ItemId: updateDefault.Item.Hex(),
		},
		Message: "Change background from " + inventory.Item.Hex() + " to " + updateDefault.Item.Hex() + " success.",
	}, nil

}

func (s userService) GetBg(token string, id string) (interface{}, error) {
	_, err := utils.IsTokenValid(token)
	if err != nil {
		return nil, err
	}
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		return nil, errs.NewInternalError(err.Error())
	}
	user, err := s.userRepo.GetById(userId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("User not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}

	bg, err := s.inventoryRepo.GetByTypeAndUserIdAndDefault("bg", user.ID, true)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return utils.DataResponse{
				Data: interfaces.DefaultLink{
					Link: "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/background%2FbgShop.png?alt=media&token=2811d8e3-6ceb-4a41-ad2d-94a511ab9cb9",
				},
				Message: "Send default background"}, nil
		}
		return nil, errs.NewInternalError(err.Error())
	}

	bgResult, err := s.bgRepo.GetById(bg.Item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.NewBadRequestError("Background not found.")
		}
		return nil, errs.NewInternalError(err.Error())
	}
	return utils.DataResponse{
		Data: interfaces.BgResponse{
			ID:     bgResult.ID.Hex(),
			Name:   bgResult.Name,
			Assets: bgResult.Assets,
			Price:  bgResult.Price,
		},
		Message: "Get background of " + user.Username + " success.",
	}, nil
}
