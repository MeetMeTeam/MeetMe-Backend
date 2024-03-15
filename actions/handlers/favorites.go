package handlers

import (
	"github.com/labstack/echo/v4"
	svInter "meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
	"net/http"
)

type favoriteHandler struct {
	favoriteService svInter.FavoriteService
}

func NewFavoriteHandler(favoriteService svInter.FavoriteService) favoriteHandler {
	return favoriteHandler{favoriteService: favoriteService}
}

// FavUser godoc
//
//	@Summary		Favorite user
//	@Description	Favorite other user.
//	@Tags			favorites
//	@Accept			json
//	@Produce		json
//	@Param        	receiverId   path      string  true  "user id that you want to like."
//	@Success		200		{object}	utils.DataResponse
//	@Router			/users/favorites/{receiverId} [post]
//
// @Security BearerAuth
func (h favoriteHandler) FavUser(c echo.Context) error {
	id := c.Param("userId")
	token := c.Request().Header.Get("Authorization")

	avatars, err := h.favoriteService.FavUser(token, id)
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

	return c.JSON(http.StatusOK, avatars)
}

// UnFavUser godoc
//
//	@Summary		UnFavorite user
//	@Description	Remove Favorite other user.
//	@Tags			favorites
//	@Accept			json
//	@Produce		json
//	@Param        	receiverId   path      string  true  "user id that you want to like."
//	@Success		200		{object}	utils.DataResponse
//	@Router			/users/favorites/{receiverId} [delete]
//
// @Security BearerAuth
func (h favoriteHandler) UnFavUser(c echo.Context) error {
	id := c.Param("userId")
	token := c.Request().Header.Get("Authorization")

	avatars, err := h.favoriteService.UnFavUser(token, id)
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

	return c.JSON(http.StatusOK, avatars)
}

// CountFavUser godoc
//
//	@Summary		Count Fav
//	@Description	Count Favorite of user.
//	@Tags			favorites
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	utils.DataResponse
//	@Router			/users/favorites [get]
//
// @Security BearerAuth
func (h favoriteHandler) CountFavUser(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")

	avatars, err := h.favoriteService.GetCountFav(token)
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

	return c.JSON(http.StatusOK, avatars)
}
