package handlers

import (
	"github.com/labstack/echo/v4"
	svInter "meetme/be/actions/services/interfaces"
	"meetme/be/errs"
	"meetme/be/utils"
	"net/http"
)

type inventoryHandler struct {
	inventoryService svInter.InventoryService
}

func NewInventoryHandler(inventoryService svInter.InventoryService) inventoryHandler {
	return inventoryHandler{inventoryService: inventoryService}
}

// GetInventory godoc
//
//	@Summary		Get user's inventory.
//	@Description	Get inventory by token.
//	@Tags			inventories
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	utils.DataResponse
//	@Router			/users/inventories [get]
//
// @Security BearerAuth
func (h inventoryHandler) GetInventory(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")

	inventory, err := h.inventoryService.GetInventory(token)
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

	return c.JSON(http.StatusOK, inventory)
}

// AddItem godoc
//
//	@Summary		Add item to inventories
//	@Description	Add item to user inventories
//	@Tags			inventories
//	@Accept			json
//	@Produce		json
//	@Param			item_id	query	string	true	"item id that you want to add"
//	@Param			item_type	query	string	true	"item type that you want to add [avatar]"
//
// @Success		200		{object}	utils.DataResponse
// @Router			/users/inventories [post]
func (h inventoryHandler) AddItem(c echo.Context) error {
	//id := c.Param("itemId")
	id := c.QueryParam("item_id")
	itemType := c.QueryParam("item_type")

	token := c.Request().Header.Get("Authorization")

	users, err := h.inventoryService.AddItem(token, id, itemType)
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
