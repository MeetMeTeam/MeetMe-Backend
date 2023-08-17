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

func (h userHandler) Register(c echo.Context) error {

	request := new(svInter.RegisterRequest)

	//if err := c.Validate(request); err != nil {
	//	return c.JSON(http.StatusBadRequest, utils.ErrorResponse{
	//		Message: err.Error(),
	//	})
	//}

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
