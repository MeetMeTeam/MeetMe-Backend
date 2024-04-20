package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"meetme/be/actions/repositories/interfaces"
)

type BgRepository struct {
	db *mongo.Database
}

func NewBgRepositoryDB(db *mongo.Database) BgRepository {
	return BgRepository{db: db}
}

func (r BgRepository) Create(request interfaces.Background) (*interfaces.Background, error) {

	_, err := r.db.Collection("bg_shops").InsertOne(context.TODO(), request)

	if err != nil {
		return nil, err
	}

	return &request, nil
}

func (r BgRepository) GetAll() ([]interfaces.BgResponse, error) {
	filter := bson.D{}
	coll := r.db.Collection("bg_shops")
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var bg []interfaces.BgResponse
	if err = cursor.All(context.TODO(), &bg); err != nil {
		panic(err)
	}

	return bg, nil
}

func (r BgRepository) GetById(id primitive.ObjectID) (*interfaces.BgResponse, error) {
	var bg interfaces.BgResponse
	filter := bson.D{{"_id", id}}
	coll := r.db.Collection("bg_shops")
	err := coll.FindOne(context.TODO(), filter).Decode(&bg)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		panic(err)
	}

	return &bg, nil
}
