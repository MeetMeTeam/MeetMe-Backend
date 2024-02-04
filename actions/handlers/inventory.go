package handlers

import svInter "meetme/be/actions/services/interfaces"

type inventoryHandler struct {
	inventoryService svInter.InventoryService
}

func NewInventoryHandler(inventoryService svInter.InventoryService) inventoryHandler {
	return inventoryHandler{inventoryService: inventoryService}
}
