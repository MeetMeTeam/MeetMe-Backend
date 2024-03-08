package handlers

import (
	"github.com/labstack/echo/v4"
	svInter "meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
	"net/http"
)

type avatarShopHandler struct {
	avatarService svInter.AvatarService
}

func NewAvatarShopHandler(avatarService svInter.AvatarService) avatarShopHandler {
	return avatarShopHandler{avatarService: avatarService}
}

// GetAvatarShop godoc
//
//	@Summary		Get avatar's shop.
//	@Description	Get avatar's shop.
//	@Tags			avatar shop
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	utils.DataResponse
//	@Router			/avatars [get]
func (h avatarShopHandler) GetAvatarShop(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	itemType := c.QueryParam("type")
	avatars, err := h.avatarService.GetAvatarShops(token, itemType)
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

// AddAvatarToShop godoc
//
//	@Summary		Add avatar to shop
//	@Description	Only admin qdd avatar to shop.
//	@Tags			avatar shop
//	@Accept			json
//	@Produce		json
//	@Param			avatars	body		interfaces.AvatarRequest	true	"request body for adding avatar to shop"
//
// @Success		200		{object}	utils.DataResponse
// @Router			/avatars [post]
// @Security BearerAuth
func (h avatarShopHandler) AddAvatarToShop(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	request := new(svInter.AvatarRequest)

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

	avatars, err := h.avatarService.AddAvatarShop(token, *request)
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
