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

	avatars, err := h.avatarService.GetAvatarShops(token)
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
