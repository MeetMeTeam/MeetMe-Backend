package interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	DisplayName string `bson:"displayName"`
	Birthday    string `bson:"birthday"`
	Email       string `bson:"email"`
	Password    string `bson:"password"`
	Image       string `bson:"image"`
	Username    string `bson:"username"`
}

type UserResponse struct {
	ID          primitive.ObjectID `bson:"_id"`
	DisplayName string             `bson:"displayName"`
	Birthday    string             `bson:"birthday"`
	Email       string             `bson:"email"`
	Password    string             `bson:"password"`
	Image       string             `bson:"image"`
	Username    string             `bson:"username"`
}
type UserRepository interface {
	GetAll() ([]UserResponse, error)
	GetByEmail(string) (*UserResponse, error)
	GetById(int) (*UserResponse, error)
	Create(User) (*User, error)
	AddFriend() (*User, error)
	GetByUsername(string) (*UserResponse, error)
	// UpdateTotalPoint(int, string) (*User, error)
}
