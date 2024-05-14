package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Gift struct {
	Img   string `bson:"img"`
	Price int    `bson:"price"`
	Name  string `bson:"name"`
}

func InsertGiftBase(db *mongo.Database) {
	// Reference the database and collection to use
	collection := db.Collection("gift_lists")
	// Create new documents
	gift := []interface{}{
		Gift{
			Name:  "Ice Cream",
			Img:   "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/gift%2Fice-cream-candy-cute-dessert-sweet-character-isolated-3d-rendering-png.webp?alt=media&token=e9e86e81-2881-4479-89da-485a49d55031",
			Price: 200,
		},
		Gift{
			Name:  "Sheep",
			Img:   "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/gift%2Fpngtree-small-sheep-3d-funny-png-image_6706536.png?alt=media&token=05b1fe83-f363-4fde-a283-4c188e28f5f6",
			Price: 300,
		},
		Gift{
			Name:  "Rose",
			Img:   "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/gift%2Frose-removebg-preview.png?alt=media&token=ad3950d4-23a3-4b7a-a21d-0024906c1a7b",
			Price: 400,
		},
		Gift{
			Name:  "Halloween Candy",
			Img:   "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/gift%2Fimage-removebg-preview%20(1).png?alt=media&token=731771e0-fa5f-4a6b-9063-a585d1dff815",
			Price: 2000,
		},
	}
	// Insert the document into the specified collection
	_, err := collection.InsertMany(context.TODO(), gift)
	if err != nil {
		return
	}
	// Find the document
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var result []Gift
	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}

	fmt.Printf("Document Found:\n%+v\n", result)

}
