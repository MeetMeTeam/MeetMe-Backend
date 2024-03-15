package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"meetme/be/actions/repositories/interfaces"
)

type FavoriteRepository struct {
	db *mongo.Database
}

func NewFavoriteRepositoryDB(db *mongo.Database) FavoriteRepository {
	return FavoriteRepository{db: db}
}
func (r FavoriteRepository) AddFav(giver primitive.ObjectID, receiver primitive.ObjectID) (*interfaces.Favorite, error) {
	newFav := interfaces.Favorite{
		Giver:    giver,
		Receiver: receiver,
	}
	_, err := r.db.Collection("favorites").InsertOne(context.TODO(), newFav)

	if err != nil {
		return nil, err
	}

	return &newFav, nil
}

func (r FavoriteRepository) GetByGiverAndReceiver(giver primitive.ObjectID, receiver primitive.ObjectID) (*interfaces.FavoriteResponse, error) {
	var fav interfaces.FavoriteResponse

	filter := bson.D{{"giver_id", giver}, {"receiver_id", receiver}}
	coll := r.db.Collection("favorites")
	err := coll.FindOne(context.TODO(), filter).Decode(&fav)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		panic(err)
	}

	return &fav, nil
}
