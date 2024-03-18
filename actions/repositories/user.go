package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"meetme/be/actions/repositories/interfaces"
	interfaces2 "meetme/be/actions/services/interfaces"
	"time"
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
func (r UserRepository) Create(user interfaces.User) (*interfaces.UserResponse, error) {

	filter := bson.D{{"email", user.Email}}

	update := bson.D{{"$set",
		bson.D{
			{"displayName", user.DisplayName},
			{"birthday", user.Birthday},
			{"password", user.Password},
			{"username", user.Username},
			{"isAdmin", user.IsAdmin},
			{"isVerify", user.IsVerify},
		}}}
	coll := r.db.Collection("users")
	_, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	var users *interfaces.UserResponse
	err = coll.FindOne(context.TODO(), filter).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil

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

func (r UserRepository) UpdateCoinById(id primitive.ObjectID, coin int) (*interfaces.User, error) {
	filter := bson.D{{"_id", id}}

	update := bson.D{{"$set", bson.D{{"coin", coin}}}}
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

func (r UserRepository) UpdateAvatarById(userId primitive.ObjectID, inventoryId primitive.ObjectID) (*interfaces.UserResponse, error) {
	filter := bson.D{{"_id", userId}}

	update := bson.D{{"$set", bson.D{{"inventory_id", inventoryId}}}}
	coll := r.db.Collection("users")
	_, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	var users *interfaces.UserResponse
	err = coll.FindOne(context.TODO(), filter).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r UserRepository) UpdateUsernameByEmail(email string, username string) (*interfaces.UserResponse, error) {
	filter := bson.D{{"email", email}}

	update := bson.D{{"$set", bson.D{{"username", username}}}}
	coll := r.db.Collection("users")
	_, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	var users *interfaces.UserResponse
	err = coll.FindOne(context.TODO(), filter).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (r UserRepository) UpdateDisplayNameByEmail(email string, displayName string) (*interfaces.UserResponse, error) {
	filter := bson.D{{"email", email}}

	update := bson.D{{"$set", bson.D{{"displayName", displayName}}}}
	coll := r.db.Collection("users")
	_, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	var users *interfaces.UserResponse
	err = coll.FindOne(context.TODO(), filter).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r UserRepository) UpdateBioByEmail(email string, bio string) (*interfaces.UserResponse, error) {
	filter := bson.D{{"email", email}}

	update := bson.D{{"$set", bson.D{{"bio", bio}}}}
	coll := r.db.Collection("users")
	_, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	var users *interfaces.UserResponse
	err = coll.FindOne(context.TODO(), filter).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r UserRepository) AddSocial(email string, social []interfaces2.EditSocial) (*interfaces.UserResponse, error) {
	filter := bson.D{{"email", email}}

	update := bson.D{{"$set",
		bson.D{{"social",
			social}}}}

	coll := r.db.Collection("users")
	_, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	var users *interfaces.UserResponse
	err = coll.FindOne(context.TODO(), filter).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r UserRepository) UpdateSocialByEmail(email string, social []interfaces2.EditSocial) (*interfaces.UserResponse, error) {
	return nil, nil
}
func (r UserRepository) UpdateBioUsernameDisplayNameByEmail(email string, bio string, username string, display string) (*interfaces.UserResponse, error) {
	filter := bson.D{{"email", email}}

	update := bson.D{{"$set", bson.D{{"bio", bio}, {"username", username}, {"displayName", display}}}}
	coll := r.db.Collection("users")
	_, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	var users *interfaces.UserResponse
	err = coll.FindOne(context.TODO(), filter).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r UserRepository) CreateVerifyMail(email string, code string, refCode string, expiredAt time.Time) (*interfaces.Mail, error) {
	newMail := interfaces.Mail{
		Email:     email,
		Code:      code,
		ExpiredAt: expiredAt,
		IsVerify:  false,
		RefCode:   refCode,
	}
	_, err := r.db.Collection("users").InsertOne(context.TODO(), newMail)

	if err != nil {
		return nil, err
	}

	return &newMail, nil
}

func (r UserRepository) UpdateVerifyMailCode(email string, code string, refCode string, expiredAt time.Time) (*interfaces.Mail, error) {
	filter := bson.D{{"email", email}}

	update := bson.D{{"$set", bson.D{{"code", code}, {"refCode", refCode}, {"expiredAt", expiredAt}}}}
	coll := r.db.Collection("users")
	_, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	var users *interfaces.Mail
	err = coll.FindOne(context.TODO(), filter).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
