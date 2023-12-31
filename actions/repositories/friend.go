package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"meetme/be/actions/repositories/interfaces"
)

type FriendRepository struct {
	db *mongo.Database
}

func NewFriendRepositoryDB(db *mongo.Database) FriendRepository {
	return FriendRepository{db: db}
}

func (r FriendRepository) Create(invite interfaces.FriendRequest) (*interfaces.FriendRequest, error) {

	newInvite := interfaces.FriendRequest{
		Receiver: invite.Receiver,
		Sender:   invite.Sender,
		Status:   "PENDING",
	}

	_, err := r.db.Collection("friends").InsertOne(context.TODO(), newInvite)

	if err != nil {
		return nil, err
	}

	return &newInvite, nil
}

func (r FriendRepository) GetByReceiverId(receiverId primitive.ObjectID, status string) ([]interfaces.FriendResponse, error) {
	var invitation []interfaces.FriendResponse

	filter := bson.D{{"receiver_id", receiverId}, {"status", status}}
	coll := r.db.Collection("friends")
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &invitation); err != nil {
		return nil, err
	}

	return invitation, nil
}

func (r FriendRepository) UpdateStatus(receiverId primitive.ObjectID, id primitive.ObjectID) ([]interfaces.FriendResponse, error) {

	filter := bson.D{}
	if id != primitive.NilObjectID {
		filter = bson.D{{"_id", id}}
	} else if receiverId != primitive.NilObjectID {
		filter = bson.D{{"receiver_id", receiverId}}
	}

	update := bson.D{{"$set", bson.D{{"status", "FRIEND"}}}}
	coll := r.db.Collection("friends")
	_, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var friend []interfaces.FriendResponse
	if err = cursor.All(context.TODO(), &friend); err != nil {
		panic(err)
	}
	return friend, nil
}

func (r FriendRepository) Delete(receiverId primitive.ObjectID, id primitive.ObjectID, status string) error {

	filter := bson.D{}
	if id != primitive.NilObjectID {
		filter = bson.D{{"_id", id}, {"status", status}}
	} else if receiverId != primitive.NilObjectID {
		filter = bson.D{{"receiver_id", receiverId}, {"status", status}}
	}

	coll := r.db.Collection("friends")

	_, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (r FriendRepository) GetByReceiverIdAndSenderId(receiverId primitive.ObjectID, senderId primitive.ObjectID) (*interfaces.FriendResponse, error) {
	var invitation interfaces.FriendResponse

	filter := bson.D{{"$or", []interface{}{
		bson.D{{"receiver_id", receiverId}, {"sender_id", senderId}},
		bson.D{{"receiver_id", senderId}, {"sender_id", receiverId}},
	}}}
	coll := r.db.Collection("friends")
	err := coll.FindOne(context.TODO(), filter).Decode(&invitation)
	if err != nil {
		return nil, err
	}

	return &invitation, nil
}

func (r FriendRepository) GetByIdAndReceiverIdAndStatus(id primitive.ObjectID, receiverId primitive.ObjectID, status string) (*interfaces.FriendResponse, error) {
	var invitation interfaces.FriendResponse

	filter := bson.D{{"_id", id}, {"receiver_id", receiverId}, {"status", status}}
	coll := r.db.Collection("friends")
	err := coll.FindOne(context.TODO(), filter).Decode(&invitation)
	if err != nil {
		return nil, err
	}

	return &invitation, nil
}

func (r FriendRepository) GetByUserId(id primitive.ObjectID, status string) ([]interfaces.FriendResponse, error) {
	var invitation []interfaces.FriendResponse

	filter := bson.D{
		//{"receiver_id", id},
		{"$or", bson.A{
			bson.D{{"sender_id", id}},
			bson.D{{"receiver_id", id}},
		}},
		{"status", status},
	}
	//filter := bson.D{{"receiver_id", id}, {"$or", bson.D{"sender_id", id}}, {"status", status}}
	coll := r.db.Collection("friends")
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &invitation); err != nil {
		return nil, err
	}

	return invitation, nil
}
