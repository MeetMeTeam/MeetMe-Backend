package handlers

import (
	"github.com/labstack/echo/v4"
	svInter "meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
	"net/http"
)

type friendInvitationHandler struct {
	userService svInter.InviteService
}

func NewFriendInvitationHandler(userService svInter.InviteService) friendInvitationHandler {
	return friendInvitationHandler{userService: userService}
}

func (h friendInvitationHandler) InviteFriend(c echo.Context) error {
	request := new(svInter.InviteRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: "Something wrong.",
		})
	}

	users, err := h.userService.InviteFriend(*request)
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
