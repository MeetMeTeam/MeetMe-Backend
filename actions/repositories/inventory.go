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
func (r InventoryRepository) GetByUserId(userId primitive.ObjectID) ([]interfaces.InventoryResponse, error) {
	var inventory []interfaces.InventoryResponse
	filter := bson.D{{"user_id", userId}}
	coll := r.db.Collection("inventories")
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &inventory); err != nil {
		return nil, err
	}

	return inventory, nil
}

func (r InventoryRepository) Create(userId primitive.ObjectID, itemId primitive.ObjectID, itemType string) (*interfaces.InventoryResponse, error) {
	newInventory := interfaces.Inventory{
		User: userId,
		Item: itemId,
		Type: itemType,
	}
	_, err := r.db.Collection("inventories").InsertOne(context.TODO(), newInventory)

	if err != nil {
		return nil, err
	}

	resultInvent, err := r.GetByUserIdAndItemId(userId, itemId)
	if err != nil {
		return nil, err
	}
	return resultInvent, nil
}

func (r InventoryRepository) GetByUserIdAndItemId(userId primitive.ObjectID, itemId primitive.ObjectID) (*interfaces.InventoryResponse, error) {
	var inventory interfaces.InventoryResponse

	filter := bson.D{{"user_id", userId}, {"item_id", itemId}}
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

func (r InventoryRepository) GetByUserIdAndItemType(userId primitive.ObjectID, itemType string) ([]interfaces.InventoryResponse, error) {
	var inventory []interfaces.InventoryResponse
	filter := bson.D{{"user_id", userId}, {"type_item", itemType}}
	coll := r.db.Collection("inventories")
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &inventory); err != nil {
		return nil, err
	}

	return inventory, nil
}

func (r InventoryRepository) GetByTypeAndUserIdAndDefault(itemType string, userId primitive.ObjectID, isDefault bool) (*interfaces.InventoryResponse, error) {
	var inventory interfaces.InventoryResponse
	filter := bson.D{{"user_id", userId}, {"type_item", itemType}, {"is_default", isDefault}}
	coll := r.db.Collection("inventories")

	err := coll.FindOne(context.TODO(), filter).Decode(&inventory)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		panic(err)
	}

	return &inventory, nil
}

func (r InventoryRepository) UpdateDefaultByUserIdAndItemIdAndType(userId primitive.ObjectID, itemId primitive.ObjectID, itemType string, setDefault bool) (*interfaces.InventoryResponse, error) {
	filter := bson.D{{"user_id", userId}, {"item_id", itemId}, {"type_item", itemType}}

	update := bson.D{{"$set", bson.D{{"is_default", setDefault}}}}
	coll := r.db.Collection("inventories")
	_, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	var inventory interfaces.InventoryResponse

	err = coll.FindOne(context.TODO(), filter).Decode(&inventory)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		panic(err)
	}

	return &inventory, nil
}
