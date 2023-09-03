package repositories

import (
	"gorm.io/gorm"
	"meetme/be/actions/repositories/interfaces"
)

type FriendInvitationRepository struct {
	db *gorm.DB
}

func NewFriendInvitationRepositoryDB(db *gorm.DB) FriendInvitationRepository {
	return FriendInvitationRepository{db: db}
}

func (r FriendInvitationRepository) Create(invite interfaces.FriendInvitation) (*interfaces.FriendInvitation, error) {

	newInvite := interfaces.FriendInvitation{
		ReceiverId: invite.ReceiverId,
		SenderId:   invite.SenderId,
	}

	result := r.db.Create(&newInvite)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newInvite, nil
}

func (r FriendInvitationRepository) GetInvitationByReceiverId(receiverId int) ([]interfaces.FriendInvitation, error) {
	var invitation []interfaces.FriendInvitation
	result := r.db.Where("receiver_id = ?", receiverId).Find(&invitation)
	if result.Error != nil {
		return nil, result.Error
	}
	return invitation, nil
}