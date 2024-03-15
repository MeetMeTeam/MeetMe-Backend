package interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

type Favorite struct {
	Giver    primitive.ObjectID `bson:"giver_id"`
	Receiver primitive.ObjectID `bson:"receiver_id"`
}

type FavoriteResponse struct {
	ID       primitive.ObjectID `bson:"id"`
	Giver    primitive.ObjectID `bson:"giver_id"`
	Receiver primitive.ObjectID `bson:"receiver_id"`
}

type FavoriteRepository interface {
	AddFav(primitive.ObjectID, primitive.ObjectID) (*Favorite, error)
	DeleteFav(primitive.ObjectID, primitive.ObjectID) error
	GetByGiverAndReceiver(primitive.ObjectID, primitive.ObjectID) (*FavoriteResponse, error)
	CountFav(primitive.ObjectID) (int, error)
}
