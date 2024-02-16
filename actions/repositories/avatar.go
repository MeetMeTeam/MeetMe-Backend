package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"meetme/be/actions/repositories/interfaces"
)

type AvatarRepository struct {
	db *mongo.Database
}

func NewAvatarRepositoryDB(db *mongo.Database) AvatarRepository {
	return AvatarRepository{db: db}
}

func (r AvatarRepository) GetById(id primitive.ObjectID) (*interfaces.AvatarResponse, error) {
	var avatar interfaces.AvatarResponse
	filter := bson.D{{"_id", id}}
	coll := r.db.Collection("avatar_shops")
	err := coll.FindOne(context.TODO(), filter).Decode(&avatar)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		panic(err)
	}

	return &avatar, nil
}

func (r AvatarRepository) GetByType(itemType string) ([]interfaces.AvatarResponse, error) {

	filter := bson.M{"type": bson.M{"$regex": "^" + itemType}}
	coll := r.db.Collection("avatar_shops")
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		panic(err)
	}
	var avatars []interfaces.AvatarResponse
	if err = cursor.All(context.TODO(), &avatars); err != nil {
		panic(err)
	}

	return avatars, nil
}

func (r AvatarRepository) GetAll() ([]interfaces.AvatarResponse, error) {
	filter := bson.D{}
	coll := r.db.Collection("avatar_shops")
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var avatars []interfaces.AvatarResponse
	if err = cursor.All(context.TODO(), &avatars); err != nil {
		panic(err)
	}

	return avatars, nil
}

func (r AvatarRepository) Create(request interfaces.Avatar) (*interfaces.Avatar, error) {

	_, err := r.db.Collection("avatar_shops").InsertOne(context.TODO(), request)

	if err != nil {
		return nil, err
	}

	return &request, nil
}
