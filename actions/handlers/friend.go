package handlers

import (
	"github.com/labstack/echo/v4"
	svInter "meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
	"net/http"
)

type friendHandler struct {
	friendService svInter.FriendService
}

func NewFriendHandler(friendService svInter.FriendService) friendHandler {
	return friendHandler{friendService: friendService}
}

// InviteFriend godoc
// @Summary      Invite Friend
// @Description  Invite friend by email.
// @Tags         friend invitation
// @Accept       json
// @Produce      json
// @Param users body interfaces.InviteRequest true "request body invite friend"
// @Success      200  {object}  utils.DataResponse
// @Router       /invitation/add [post]
func (h friendHandler) InviteFriend(c echo.Context) error {
	request := new(svInter.InviteRequest)
	token := c.Request().Header.Get("Authorization")

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: "Something wrong.",
		})
	}

	users, err := h.friendService.InviteFriend(token, *request)
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

func (h friendHandler) CheckFriendInvite(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")

	users, err := h.friendService.CheckFriendInvite(token)
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

func (h friendHandler) RejectFriend(c echo.Context) error {
	id := c.Param("inviteId")
	token := c.Request().Header.Get("Authorization")

	users, err := h.friendService.RejectInvitation(token, id)
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
func (h friendHandler) AcceptFriend(c echo.Context) error {
	id := c.Param("inviteId")
	token := c.Request().Header.Get("Authorization")

	users, err := h.friendService.AcceptInvitation(token, id)
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
