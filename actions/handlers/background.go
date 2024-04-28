package handlers

import (
	"github.com/labstack/echo/v4"
	svInter "meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
	"net/http"
)

type bgShopHandler struct {
	bgService svInter.BgService
}

func NewBgShopHandler(bgService svInter.BgService) bgShopHandler {
	return bgShopHandler{bgService: bgService}
}

// AddBgToShop godoc
//
//	@Summary		Add background to shop
//	@Description	Only admin add bg to shop.
//	@Tags			background shop
//	@Accept			json
//	@Produce		json
//	@Param			bg	body		interfaces.BgRequest	true	"request body for adding background to shop"
//
// @Success		200		{object}	utils.DataResponse
// @Router			/backgrounds [post]
// @Security BearerAuth
func (h bgShopHandler) AddBgToShop(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	request := new(svInter.BgRequest)

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

	avatars, err := h.bgService.AddBgShop(token, *request)
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

// GetBgShop godoc
//
//	@Summary		Get background's shop.
//	@Description	Get background's shop.
//	@Tags			background shop
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	utils.DataResponse
//	@Router			/backgrounds [get]
//
// @Security BearerAuth
func (h bgShopHandler) GetBgShop(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")

	themes, err := h.bgService.GetBgShops(token)
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

	return c.JSON(http.StatusOK, themes)
}
