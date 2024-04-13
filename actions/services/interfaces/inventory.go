package interfaces

type InventoryResponse struct {
}
type InventoryService interface {
	GetInventory(string, string) (interface{}, error)
	AddItem(string, string, string) (interface{}, error)
}
