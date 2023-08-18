package handlers

import (
	svInter "meetme/be/actions/services/interfaces"
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
// @Summary      Register user
// @Description  Create user.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param users body interfaces.RegisterRequest true "request body register"
// @Success      200  {object}  utils.DataResponse
// @Router       /register [post]
func (h userHandler) Register(c echo.Context) error {
	request := new(svInter.RegisterRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: "Something wrong.",
		})
	}

	users, err := h.userService.CreateUser(*request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}

// Login godoc
// @Summary      Login
// @Description  Login user.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param users body interfaces.Login true "request body login"
// @Success      200  {object}  utils.DataResponse
// @Router       /login [post]
func (h userHandler) Login(c echo.Context) error {

	request := new(svInter.Login)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: "Something wrong.",
		})
	}

	users, err := h.userService.Login(*request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}
