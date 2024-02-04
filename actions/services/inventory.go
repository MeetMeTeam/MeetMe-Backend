package services

import (
	"meetme/be/actions/repositories"
	"meetme/be/actions/services/interfaces"
)

type inventoryService struct {
	inventoryRepo repositories.InventoryRepository
}

func NewInventoryService(inventoryRepo repositories.InventoryRepository) interfaces.InventoryService {
	return inventoryService{inventoryRepo: inventoryRepo}
}
