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
	coll := r.db.Collection("user")
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
	coll := r.db.Collection("user")
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
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Birthday:  user.Birthday,
		Password:  user.Password,
		Image:     user.Image,
		Username:  user.Username,
	}
	_, err := r.db.Collection("user").InsertOne(context.TODO(), newUser)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (r UserRepository) GetAll() ([]interfaces.UserResponse, error) {

	filter := bson.D{}
	coll := r.db.Collection("user")
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
