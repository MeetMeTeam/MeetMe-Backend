package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"meetme/be/actions/repositories/interfaces"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepositoryDB(db *mongo.Database) UserRepository {
	return UserRepository{db: db}
}

func (r UserRepository) GetByEmail(email string) (*interfaces.UserResponse, error) {
	var users interfaces.UserResponse
	filter := bson.D{{"email", email}}
	coll := r.db.Collection("users")
	err := coll.FindOne(context.TODO(), filter).Decode(&users)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		panic(err)
	}

	return &users, nil
}

func (r UserRepository) GetByUsername(username string) (*interfaces.UserResponse, error) {
	var users interfaces.UserResponse
	filter := bson.D{{"username", username}}
	coll := r.db.Collection("users")
	err := coll.FindOne(context.TODO(), filter).Decode(&users)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		panic(err)
	}

	return &users, nil
}

func (r UserRepository) GetById(id primitive.ObjectID) (*interfaces.UserResponse, error) {

	var users interfaces.UserResponse
	filter := bson.D{{"_id", id}}
	coll := r.db.Collection("users")
	err := coll.FindOne(context.TODO(), filter).Decode(&users)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		panic(err)
	}

	return &users, nil
}
func (r UserRepository) Create(user interfaces.User) (*interfaces.User, error) {

	newUser := interfaces.User{
		DisplayName: user.DisplayName,
		Email:       user.Email,
		Birthday:    user.Birthday,
		Password:    user.Password,
		Image:       user.Image,
		Username:    user.Username,
	}
	_, err := r.db.Collection("users").InsertOne(context.TODO(), newUser)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (r UserRepository) GetAll() ([]interfaces.UserResponse, error) {

	filter := bson.D{}
	coll := r.db.Collection("users")
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var users []interfaces.UserResponse
	if err = cursor.All(context.TODO(), &users); err != nil {
		panic(err)
	}

	return users, nil
}

func (r UserRepository) UpdatePasswordByEmail(email string, password string) (*interfaces.User, error) {
	filter := bson.D{{"email", email}}

	update := bson.D{{"$set", bson.D{{"password", password}}}}
	coll := r.db.Collection("users")
	_, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	var users *interfaces.User
	err = coll.FindOne(context.TODO(), filter).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
