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
//
//	@Summary		Invite Friend
//	@Description	Invite friend by email.
//	@Tags			invitations
//	@Accept			json
//	@Produce		json
//	@Param			users	body		interfaces.InviteRequest	true	"request body invite friend"
//	@Success		200		{object}	utils.DataResponse
//	@Router			/invitations [post]
//
// @Security BearerAuth
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

// CheckFriendInvite godoc
//
//	@Summary		List Invitation.
//	@Description	Check Invitation List.
//	@Tags			invitations
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	utils.DataResponse
//	@Router			/invitations [get]
//
// @Security BearerAuth
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

// RejectFriend godoc
//
//	@Summary		Reject Invitation
//	@Description	Reject Invitation by Id.
//	@Tags			invitations
//	@Accept			json
//	@Produce		json
//	@Param        	id   path      string  true  "Invitation ID"
//	@Success		200		{object}	utils.DataResponse
//	@Router			/invitations/{id} [delete]
//
// @Security BearerAuth
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

// RejectAllFriend godoc
//
//	@Summary		Reject All Invitations
//	@Description	Reject All Invitations.
//	@Tags			invitations
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	utils.DataResponse
//	@Router			/invitations [delete]
//
// @Security BearerAuth
func (h friendHandler) RejectAllFriend(c echo.Context) error {

	token := c.Request().Header.Get("Authorization")

	users, err := h.friendService.RejectAllInvitation(token)
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

// AcceptFriend godoc
//
//	@Summary		Accept Invitation
//	@Description	Accept Invitation by Id.
//	@Tags			invitations
//	@Accept			json
//	@Produce		json
//	@Param        	id   path      string  true  "Invitation ID"
//	@Success		200		{object}	utils.DataResponse
//	@Router			/invitations/{id} [put]
//
// @Security BearerAuth
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

// AcceptAllFriend godoc
//
//	@Summary		Accept All Invitations
//	@Description	Accept All Invitations.
//	@Tags			invitations
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	utils.DataResponse
//	@Router			/invitations [put]
//
// @Security BearerAuth
func (h friendHandler) AcceptAllFriend(c echo.Context) error {

	token := c.Request().Header.Get("Authorization")

	users, err := h.friendService.AcceptAllInvitations(token)
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

// FriendList godoc
//
//	@Summary		List Friends.
//	@Description	Get Friends List.
//	@Tags			friends
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	utils.DataResponse
//	@Router			/friends [get]
//
// @Security BearerAuth
func (h friendHandler) FriendList(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")

	users, err := h.friendService.GetFriend(token)
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

// RemoveFriend godoc
//
//	@Summary		Remove Friend
//	@Description	Remove Friend by Id.
//	@Tags			friends
//	@Accept			json
//	@Produce		json
//	@Param        	id   path      string  true  "Friend ID"
//	@Success		200		{object}	utils.DataResponse
//	@Router			/friends/{id} [delete]
//
// @Security BearerAuth
func (h friendHandler) RemoveFriend(c echo.Context) error {
	id := c.Param("friendId")
	token := c.Request().Header.Get("Authorization")

	users, err := h.friendService.DeleteFriend(token, id)
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
