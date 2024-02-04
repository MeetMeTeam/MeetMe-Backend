package interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

type InventoryResponse struct {
	ID   primitive.ObjectID `bson:"_id"`
	User primitive.ObjectID `bson:"user_id"`
	Item primitive.ObjectID `bson:"item_id"`
	Type string             `bson:"type_item"`
}
type InventoryRepository interface {
	GetById(primitive.ObjectID) (*InventoryResponse, error)
}
