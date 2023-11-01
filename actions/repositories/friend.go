package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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

func (r FriendRepository) UpdateStatus(id primitive.ObjectID) (*interfaces.FriendResponse, error) {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"status", "FRIEND"}}}}
	coll := r.db.Collection("friends")
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
		return nil, err
	}

	var friend interfaces.FriendResponse
	err = coll.FindOne(context.TODO(), filter).Decode(&friend)
	if err != nil {
		log.Print("mongo no doc")
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		return nil, err
	}
	return &friend, nil
}

func (r FriendRepository) Delete(inviteId primitive.ObjectID) error {

	filter := bson.D{{"_id", inviteId}, {"status", "PENDING"}}
	coll := r.db.Collection("friends")

	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (r FriendRepository) GetByReceiverIdAndSenderId(receiverId primitive.ObjectID, senderId primitive.ObjectID) (*interfaces.FriendResponse, error) {
	var invitation interfaces.FriendResponse

	filter := bson.D{{"receiver_id", receiverId}, {"sender_id", senderId}}
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
