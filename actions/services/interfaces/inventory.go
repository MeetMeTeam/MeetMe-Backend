package interfaces

type InventoryResponse struct {
}
type InventoryService interface {
	GetInventory(string) (interface{}, error)
}
