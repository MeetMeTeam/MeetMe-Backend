package repositories

import (
	"meetme/be/actions/repositories/interfaces"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (r UserRepository) Create(user interfaces.User) (*interfaces.User, error) {

	newUser := interfaces.User{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Birthday:  user.Birthday,
		Password:  user.Password,
	}

	result := r.db.Create(&newUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newUser, nil
}
