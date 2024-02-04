package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"meetme/be/actions/repositories/interfaces"
)

type InventoryRepository struct {
	db *mongo.Database
}

func NewInventoryRepositoryDB(db *mongo.Database) InventoryRepository {
	return InventoryRepository{db: db}
}

func (r InventoryRepository) GetById(id primitive.ObjectID) (*interfaces.InventoryResponse, error) {
	var inventory interfaces.InventoryResponse
	filter := bson.D{{"_id", id}}
	coll := r.db.Collection("inventories")
	err := coll.FindOne(context.TODO(), filter).Decode(&inventory)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		panic(err)
	}

	return &inventory, nil
}