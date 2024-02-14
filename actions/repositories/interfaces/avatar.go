package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AvatarResponse struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name"`
	Assets  []string           `bson:"assets"`
	Preview string             `bson:"preview"`
	Price   int                `bson:"price"`
}
type Avatar struct {
	Name    string   `bson:"name"`
	Assets  []string `bson:"assets"`
	Preview string   `bson:"preview"`
	Price   int      `bson:"price"`
}
type AvatarRepository interface {
	GetById(primitive.ObjectID) (*AvatarResponse, error)
	GetAll() ([]AvatarResponse, error)
	Create(Avatar) (*Avatar, error)
}
