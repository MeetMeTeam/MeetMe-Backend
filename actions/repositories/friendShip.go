package repositories

import (
	"gorm.io/gorm"
	"meetme/be/actions/repositories/interfaces"
)

type FriendshipRepository struct {
	db *gorm.DB
}

func NewFriendshipRepositoryDB(db *gorm.DB) FriendshipRepository {
	return FriendshipRepository{db: db}
}

func (r FriendshipRepository) Create(receiverId int, senderId int) (*interfaces.Friendship, error) {

	newFriend := interfaces.Friendship{
		UserId1: receiverId,
		UserId2: senderId,
	}

	result := r.db.Create(&newFriend)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newFriend, nil
}
