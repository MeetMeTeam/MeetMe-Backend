package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meetme/be/actions/repositories/interfaces"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepositoryDB(db *mongo.Database) UserRepository {
	return UserRepository{db: db}
}

func (r UserRepository) GetByEmail(email string) (*interfaces.User, error) {

	//var user interfaces.User
	//result := r.db.Where("email = ?", email).First(&user)
	//if result.Error != nil {
	//	return nil, result.Error
	//}
	//return &user, nil
	return nil, nil
}

func (r UserRepository) GetById(id int) (*interfaces.User, error) {
	//var user interfaces.User
	//result := r.db.First(&user, id)
	//if result.Error != nil {
	//	return nil, result.Error
	//}
	//return &user, nil

	return nil, nil
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
	_, err := r.db.Collection("users").InsertOne(context.TODO(), newUser)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (r UserRepository) GetAll() ([]interfaces.User, error) {
	//client, _ := mongo.Connect(context.TODO())
	filter := bson.D{}
	coll := r.db.Collection("users")
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var users []interfaces.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		panic(err)
	}

	return users, nil
}
