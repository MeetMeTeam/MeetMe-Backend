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

func (r UserRepository) GetByEmail(email string) (*interfaces.User, error) {

	var user interfaces.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r UserRepository) GetById(id int) (*interfaces.User, error) {
	var user interfaces.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil

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

	result := r.db.Create(&newUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newUser, nil
}

func (r UserRepository) GetAll() ([]interfaces.User, error) {
	var users []interfaces.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
