package handlers

import (
	"github.com/labstack/echo/v4"
	svInter "meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
	"net/http"
	"strconv"
)

type friendInvitationHandler struct {
	userService svInter.InviteService
}

func NewFriendInvitationHandler(userService svInter.InviteService) friendInvitationHandler {
	return friendInvitationHandler{userService: userService}
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
func (h friendInvitationHandler) InviteFriend(c echo.Context) error {
	request := new(svInter.InviteRequest)
	token := c.Request().Header.Get("Authorization")

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: "Something wrong.",
		})
	}

	users, err := h.userService.InviteFriend(token, *request)
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

func (h friendInvitationHandler) CheckFriendInvite(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("receiverId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
	}

	users, err := h.userService.CheckFriendInvite(id)
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

func (h friendInvitationHandler) RejectFriend(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("inviteId"))
	token := c.Request().Header.Get("Authorization")

	users, err := h.userService.RejectInvitation(token, id)
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

func (h friendInvitationHandler) AcceptFriend(c echo.Context) error {
	request := new(svInter.InviteRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: "Something wrong.",
		})
	}

	users, err := h.userService.AcceptInvitation(*request)
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
