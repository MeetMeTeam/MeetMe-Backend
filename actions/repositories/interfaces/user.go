package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"meetme/be/actions/services/interfaces"
	"time"
)

type User struct {
	DisplayName string             `bson:"displayName"`
	Birthday    string             `bson:"birthday"`
	Email       string             `bson:"email"`
	Password    string             `bson:"password"`
	Username    string             `bson:"username"`
	Coin        int                `bson:"coin"`
	Inventory   primitive.ObjectID `bson:"inventory_id"`
	IsAdmin     bool               `bson:"isAdmin"`
	Bio         string             `bson:"bio"`
	IsVerify    bool               `bson:"isVerify"`
}

type UserResponse struct {
	ID          primitive.ObjectID      `bson:"_id"`
	DisplayName string                  `bson:"displayName"`
	Birthday    string                  `bson:"birthday"`
	Email       string                  `bson:"email"`
	Password    string                  `bson:"password"`
	Username    string                  `bson:"username"`
	Coin        int                     `bson:"coin"`
	Inventory   primitive.ObjectID      `bson:"inventory_id"`
	IsAdmin     bool                    `bson:"isAdmin"`
	Bio         string                  `bson:"bio"`
	Social      []interfaces.EditSocial `bson:"social"`
	Code        string                  `bson:"code"`
	RefCode     string                  `bson:"refCode"`
	ExpiredAt   time.Time               `bson:"expiredAt"`
	IsVerify    bool                    `bson:"isVerify"`
}

type Mail struct {
	Email     string    `bson:"email"`
	Code      string    `bson:"code"`
	RefCode   string    `bson:"refCode"`
	IsVerify  bool      `bson:"isVerify"`
	ExpiredAt time.Time `bson:"expiredAt"`
}

type UserRepository interface {
	GetAll() ([]UserResponse, error)
	GetByEmail(string) (*UserResponse, error)
	//GetByEmailAndIsVerify(string, bool) (*UserResponse, error)
	GetById(int) (*UserResponse, error)
	Create(User) (*User, error)
	AddFriend() (*User, error)
	GetByUsername(string) (*UserResponse, error)
	AddSocial(string, interfaces.EditSocial) (*UserResponse, error)
	UpdatePasswordByEmail(string, string) (*User, error)
	UpdateCoinById(primitive.ObjectID, int) (*User, error)
	UpdateAvatarById(primitive.ObjectID, primitive.ObjectID) (*UserResponse, error)
	UpdateUsernameByEmail(string, string) (*UserResponse, error)
	UpdateDisplayNameByEmail(string, string) (*UserResponse, error)
	UpdateBioByEmail(string, string) (*UserResponse, error)
	UpdateSocialByEmail(string, interfaces.EditSocial) (*UserResponse, error)
	UpdateBioUsernameDisplayNameByEmail(string, string, string, string) (*UserResponse, error)
	CreateVerifyMail(string, string, string, time.Time) (*Mail, error)
	UpdateVerifyMailCode(string, string, string, time.Time) (*Mail, error)
}
