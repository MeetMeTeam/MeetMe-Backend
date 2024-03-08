package handlers

import (
	svInter "meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService svInter.UserService
}

func NewUserHandler(userService svInter.UserService) userHandler {
	return userHandler{userService: userService}
}

// Register godoc
//
//	@Summary		Register user
//	@Description	Create user.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			users	body		interfaces.RegisterRequest	true	"request body register"
//	@Success		200		{object}	utils.DataResponse
//	@Router			/register [post]
func (h userHandler) Register(c echo.Context) error {
	request := new(svInter.RegisterRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}
	errrr := utils.CustomValidator(*request)

	if errrr != nil {
		return c.JSON(http.StatusBadRequest, utils.ValidateResponse{
			Message: errrr,
		})
	}

	users, err := h.userService.CreateUser(*request)
	if err != nil {

		appErr, ok := err.(errs.AppError)
		if ok {
			return c.JSON(appErr.Code, utils.ErrorResponse{
				Message: appErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login user.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			users	body		interfaces.Login	true	"request body login"
//	@Success		200		{object}	utils.DataResponse
//	@Router			/login [post]
func (h userHandler) Login(c echo.Context) error {

	request := new(svInter.Login)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}
	errrr := utils.CustomValidator(*request)

	if errrr != nil {
		return c.JSON(http.StatusBadRequest, utils.ValidateResponse{
			Message: errrr,
		})
	}

	users, err := h.userService.Login(*request)
	if err != nil {

		appErr, ok := err.(errs.AppError)
		if ok {
			return c.JSON(appErr.Code, utils.ErrorResponse{
				Message: appErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}

// GetAllUser godoc
//
//	@Summary		Get all users
//	@Description	return list users.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.DataResponse
//	@Router			/users [get]
func (h userHandler) GetAllUser(c echo.Context) error {
	users, err := h.userService.GetUsers()
	if err != nil {

		appErr, ok := err.(errs.AppError)
		if ok {
			return c.JSON(appErr.Code, utils.ErrorResponse{
				Message: appErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}

// RefreshToken godoc
//
//	@Summary		Refresh Token
//	@Description	return new access token.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.DataResponse
//	@Router			/refresh [post]
//
// @Security BearerAuth
func (h userHandler) RefreshToken(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")

	users, err := h.userService.RefreshToken(token)
	if err != nil {

		appErr, ok := err.(errs.AppError)
		if ok {
			return c.JSON(appErr.Code, utils.ErrorResponse{
				Message: appErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}

// SendMailForResetPassword godoc
//
//	@Summary		Forgot Password
//	@Description	Send mail to reset password.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			email	body		interfaces.Email	true	"request body to send mail"
//	@Success		200		{object}	utils.ErrorResponse
//	@Router			/users/forgot-password [put]
func (h userHandler) SendMailForResetPassword(c echo.Context) error {

	request := new(svInter.Email)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}
	errrr := utils.CustomValidator(*request)

	if errrr != nil {
		return c.JSON(http.StatusBadRequest, utils.ValidateResponse{
			Message: errrr,
		})
	}
	message, err := h.userService.ForgotPassword(*request)
	if err != nil {

		appErr, ok := err.(errs.AppError)
		if ok {
			return c.JSON(appErr.Code, utils.ErrorResponse{
				Message: appErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, message)
}

// ChangePassword godoc
//
//	@Summary		Reset Password
//	@Description	Change password.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			email	body		interfaces.Password	true	"request body to change password"
//	@Success		200		{object}	utils.ErrorResponse
//	@Router			/users/reset-password [put]
//
// @Security BearerAuth
func (h userHandler) ChangePassword(c echo.Context) error {

	request := new(svInter.Password)
	token := c.Request().Header.Get("Authorization")

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}
	errrr := utils.CustomValidator(*request)

	if errrr != nil {
		return c.JSON(http.StatusBadRequest, utils.ValidateResponse{
			Message: errrr,
		})
	}
	message, err := h.userService.ResetPassword(token, *request)
	if err != nil {

		appErr, ok := err.(errs.AppError)
		if ok {
			return c.JSON(appErr.Code, utils.ErrorResponse{
				Message: appErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, message)
}

// GetCoins godoc
//
//	@Summary		Get Coin.
//	@Description	Get coin by token.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	utils.ErrorResponse
//	@Router			/users/coins [get]
//
// @Security BearerAuth
func (h userHandler) GetCoins(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")

	users, err := h.userService.GetCoin(token)
	if err != nil {

		appErr, ok := err.(errs.AppError)
		if ok {
			return c.JSON(appErr.Code, utils.ErrorResponse{
				Message: appErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}

// GetAvatarsByUserId godoc
//
//	@Summary		Get Avatar.
//	@Description	Get avatar by user Id.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param        	id   path      string  true  "User ID"
//	@Success		200		{object}	utils.DataResponse
//	@Router			/users/avatars/{id} [get]
//
// @Security BearerAuth
func (h userHandler) GetAvatarsByUserId(c echo.Context) error {
	id := c.Param("userId")
	token := c.Request().Header.Get("Authorization")

	users, err := h.userService.GetAvatars(token, id)
	if err != nil {

		appErr, ok := err.(errs.AppError)
		if ok {
			return c.JSON(appErr.Code, utils.ErrorResponse{
				Message: appErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}

// ChangeAvatar godoc
//
//	@Summary		Change Avatar.
//	@Description	User change avatar.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param        	id   path      string  true  "Item ID"
//	@Success		200		{object}	utils.DataResponse
//	@Router			/users/avatars/{id} [put]
//
// @Security BearerAuth
func (h userHandler) ChangeAvatar(c echo.Context) error {
	id := c.Param("itemId")
	token := c.Request().Header.Get("Authorization")

	users, err := h.userService.ChangeAvatar(token, id)
	if err != nil {

		appErr, ok := err.(errs.AppError)
		if ok {
			return c.JSON(appErr.Code, utils.ErrorResponse{
				Message: appErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}

// EditUserInfo godoc
//
//	@Summary		Edit profile.
//	@Description	User change profile information.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			users	body		interfaces.EditUserRequest	true	"request body for editing user's profile"
//
// @Success		200		{object}	utils.DataResponse
// @Router			/users [put]
//
// @Security BearerAuth
func (h userHandler) EditUserInfo(c echo.Context) error {
	request := new(svInter.EditUserRequest)
	token := c.Request().Header.Get("Authorization")

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}
	errrr := utils.CustomValidator(*request)

	if errrr != nil {
		return c.JSON(http.StatusBadRequest, utils.ValidateResponse{
			Message: errrr,
		})
	}

	users, err := h.userService.EditUser(*request, token)
	if err != nil {

		appErr, ok := err.(errs.AppError)
		if ok {
			return c.JSON(appErr.Code, utils.ErrorResponse{
				Message: appErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}
