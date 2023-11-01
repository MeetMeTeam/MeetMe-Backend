package interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

type FriendRequest struct {
	Receiver primitive.ObjectID `bson:"receiver_id"`
	Sender   primitive.ObjectID `bson:"sender_id"`
	Status   string             `bson:"status"`
}

type FriendResponse struct {
	ID      primitive.ObjectID `bson:"_id"`
	Receive primitive.ObjectID `bson:"receiver_id"`
	Sender  primitive.ObjectID `bson:"sender_id"`
	Status  string             `bson:"status"`
}

type FriendRepository interface {
	Create(FriendRequest) (*FriendRequest, error)
	GetInvitationByReceiverId(string) ([]FriendResponse, error)
	UpdateStatus(primitive.ObjectID) (*FriendRequest, error)
	Delete(primitive.ObjectID) error
	GetByReceiverIdAndSenderId(primitive.ObjectID, primitive.ObjectID) (*FriendResponse, error)
	//GetInvitationByIdAndReceiverId(string, string) (*FriendInvitationResponse, error)
}
