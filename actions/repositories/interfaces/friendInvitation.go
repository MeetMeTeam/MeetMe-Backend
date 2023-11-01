package interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

type FriendInvitation struct {
	ReceiverId string `db:"receiver_id"`
	SenderId   string `db:"sender_id"`
}

type FriendInvitationResponse struct {
	ID         primitive.ObjectID `bson:"_id"`
	ReceiverId primitive.ObjectID `db:"receiver_id"`
	SenderId   primitive.ObjectID `db:"sender_id"`
}

type FriendInvitationRepository interface {
	Create(FriendInvitation) (*FriendInvitation, error)
	GetInvitationByReceiverId(string) ([]FriendInvitationResponse, error)
	Delete(int) error
	GetByReceiverIdAndSenderId(string, string) (*FriendInvitationResponse, error)
	GetInvitationByIdAndReceiverId(string, string) (*FriendInvitationResponse, error)
}
