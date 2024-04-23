package interfaces

type ChangeBgResponse struct {
	ItemId string `json:"item_id"`
}
type InventoryService interface {
	GetInventory(string, string) (interface{}, error)
	AddItem(string, string, string) (interface{}, error)
}
