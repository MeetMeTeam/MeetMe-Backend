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

func (r FriendshipRepository) GetFriendByReceiverAndSender(receiverId int, senderId int) (*interfaces.Friendship, error) {
	var invitation interfaces.Friendship
	result := r.db.Where("(user_id1 = ? AND user_id2 = ?) OR (user_id1 = ? AND user_id2 = ?)", receiverId, senderId, senderId, receiverId).First(&invitation)
	if result.Error != nil {
		return nil, result.Error
	}
	return &invitation, nil
}

func (r FriendshipRepository) GetFriendById(userId int) ([]interfaces.Friendship, error) {
	var friend []interfaces.Friendship
	result := r.db.Where("user_id1 = ? OR user_id2 = ?", userId, userId).Find(&friend)
	if result.Error != nil {
		return nil, result.Error
	}
	return friend, nil
}
