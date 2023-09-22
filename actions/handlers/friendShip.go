package handlers

import (
	"github.com/labstack/echo/v4"
	svInter "meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
	"net/http"
)

type friendShipHandler struct {
	friendService svInter.FriendShipService
}

func NewFriendShipHandler(friendService svInter.FriendShipService) friendShipHandler {
	return friendShipHandler{friendService: friendService}
}

func (h friendShipHandler) FriendList(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	result, err := h.friendService.GetFriendList(token)
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
	return c.JSON(http.StatusOK, result)
}
