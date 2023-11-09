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
	GetByReceiverId(primitive.ObjectID, string) ([]FriendResponse, error)
	UpdateStatus(primitive.ObjectID, primitive.ObjectID) ([]FriendResponse, error)
	Delete(primitive.ObjectID) error
	GetByReceiverIdAndSenderId(primitive.ObjectID, primitive.ObjectID) (*FriendResponse, error)
	GetByIdAndReceiverIdAndStatus(primitive.ObjectID, primitive.ObjectID, string) (*FriendResponse, error)
	GetByUserId(primitive.ObjectID, string) ([]FriendResponse, error)
	//GetInvitationByIdAndReceiverId(string, string) (*FriendInvitationResponse, error)
}
