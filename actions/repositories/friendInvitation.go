package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meetme/be/actions/repositories/interfaces"
)

type FriendInvitationRepository struct {
	db *mongo.Database
}

func NewFriendInvitationRepositoryDB(db *mongo.Database) FriendInvitationRepository {
	return FriendInvitationRepository{db: db}
}

func (r FriendInvitationRepository) Create(invite interfaces.FriendInvitation) (*interfaces.FriendInvitation, error) {

	newInvite := interfaces.FriendInvitation{
		ReceiverId: invite.ReceiverId,
		SenderId:   invite.SenderId,
	}

	_, err := r.db.Collection("friendInvitation").InsertOne(context.TODO(), newInvite)

	if err != nil {
		return nil, err
	}

	return &newInvite, nil
}

func (r FriendInvitationRepository) GetInvitationByReceiverId(receiverId string) ([]interfaces.FriendInvitationResponse, error) {

	filter := bson.D{{"receiver_id", receiverId}}
	coll := r.db.Collection("friendInvitation")
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var invitation []interfaces.FriendInvitationResponse
	if err = cursor.All(context.TODO(), &invitation); err != nil {
		panic(err)
	}

	return invitation, nil

}

func (r FriendInvitationRepository) Delete(inviteId string) error {
	//var invitation interfaces.FriendInvitation
	//result := r.db.Where("id = ?", inviteId).Delete(&invitation)
	//if result.Error != nil {
	//	return result.Error
	//}
	return nil
}

func (r FriendInvitationRepository) GetByReceiverIdAndSenderId(receiverId string, senderId string) (*interfaces.FriendInvitationResponse, error) {
	//result := r.db.Where("(receiver_id = ? AND sender_id = ?) OR (receiver_id = ? AND sender_id = ?)", receiverId, senderId, senderId, receiverId).First(&invitation)

	var invitation interfaces.FriendInvitationResponse
	filter := bson.D{{"$or", bson.A{
		bson.D{{"$and",
			bson.A{
				bson.D{{"receiver_id", receiverId}},
				bson.D{{"sender_id", senderId}},
			}},
		},
		bson.D{{"$and",
			bson.A{
				bson.D{{"receiver_id", senderId}},
				bson.D{{"sender_id", receiverId}},
			}},
		},
	}}}

	coll := r.db.Collection("friendInvitation")
	err := coll.FindOne(context.TODO(), filter).Decode(&invitation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		panic(err)
	}

	return &invitation, nil
}

func (r FriendInvitationRepository) GetInvitationByIdAndReceiverId(id string, receiverId string) (*interfaces.FriendInvitationResponse, error) {
	//result := r.db.Where("(id = ? AND receiver_id = ?)", id, receiverId).First(&invitation)

	var invitation interfaces.FriendInvitationResponse
	filter := bson.D{{"$and",
		bson.A{
			bson.D{{"id", id}},
			bson.D{{"receiver_id", receiverId}},
		}}}

	coll := r.db.Collection("friendInvitation")
	err := coll.FindOne(context.TODO(), filter).Decode(&invitation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		panic(err)
	}

	return &invitation, nil
	return nil, nil
}
