package interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

type BgResponse struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	Assets string             `bson:"assets"`
	Price  int                `bson:"price"`
}
type Background struct {
	Name   string `bson:"name"`
	Assets string `bson:"assets"`
	Price  int    `bson:"price"`
}
type BgRepository interface {
	//GetById(primitive.ObjectID) (*AvatarResponse, error)
	GetAll() ([]BgResponse, error)
	Create(Background) (*Background, error)
	//GetByType(string) ([]AvatarResponse, error)
}
