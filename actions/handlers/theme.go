package handlers

import (
	"github.com/labstack/echo/v4"
	svInter "meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
	"net/http"
)

type themeShopHandler struct {
	themeService svInter.ThemeService
}

func NewThemeShopHandler(themeService svInter.ThemeService) themeShopHandler {
	return themeShopHandler{themeService: themeService}
}

// AddThemeToShop godoc
//
//	@Summary		Add theme to shop
//	@Description	Only admin add theme to shop.
//	@Tags			theme shop
//	@Accept			json
//	@Produce		json
//	@Param			avatars	body		interfaces.ThemeCreateRequest	true	"request body for adding theme to shop"
//
// @Success		200		{object}	utils.DataResponse
// @Router			/themes [post]
// @Security BearerAuth
func (h themeShopHandler) AddThemeToShop(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	request := new(svInter.ThemeCreateRequest)

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

	avatars, err := h.themeService.AddThemeShop(token, *request)
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
