package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Avatar struct {
	Name    string   `bson:"name"`
	Assets  []string `bson:"assets"`
	Preview string   `bson:"preview"`
	Price   int      `bson:"price"`
	Type    string   `bson:"type"`
}

type Theme struct {
	Name   string `bson:"name"`
	Assets string `bson:"assets"`
	Price  int    `bson:"price"`
	Song   string `bson:"song"`
}

type Background struct {
	Name   string `bson:"name"`
	Assets string `bson:"assets"`
	Price  int    `bson:"price"`
}

func InsertAvatarBase(db *mongo.Database) {
	// Reference the database and collection to use
	collection := db.Collection("avatar_shops")
	// Create new documents
	avatar := []interface{}{
		Avatar{
			Name:    "Steak",
			Assets:  []string{"https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/avatar%2Ftest%2Fgifmaker_me.gif?alt=media&token=c4f1e971-2e92-42e0-b56e-30941347e199"},
			Preview: "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/avatar%2Ftest%2Fgifmaker_me.gif?alt=media&token=c4f1e971-2e92-42e0-b56e-30941347e199",
			Price:   200,
		},
		Avatar{
			Name:    "Happy Banana",
			Assets:  []string{"https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/avatar%2Ftest%2Foutput-onlinegiftools%20(2).gif?alt=media&token=314c1f28-502f-4668-9c7b-e632d078723b"},
			Preview: "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/avatar%2Ftest%2Foutput-onlinegiftools%20(2).gif?alt=media&token=314c1f28-502f-4668-9c7b-e632d078723b",
			Price:   100,
		},
		Avatar{
			Name:    "Tom Yum",
			Assets:  []string{"https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/avatar%2Ftest%2Frectangle.gif?alt=media&token=2e088ea5-13b6-4c95-8b76-736732615f37"},
			Preview: "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/avatar%2Ftest%2Frectangle.gif?alt=media&token=2e088ea5-13b6-4c95-8b76-736732615f37",
			Price:   5,
			Type:    "C_1",
		},
		Avatar{
			Name:    "Kao Kreab",
			Assets:  []string{"https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/avatar%2Ftest%2Ftriangle.gif?alt=media&token=13eae37a-09c7-40e2-af31-d071a1df2582"},
			Preview: "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/avatar%2Ftest%2Ftriangle.gif?alt=media&token=13eae37a-09c7-40e2-af31-d071a1df2582",
			Price:   5,
			Type:    "C_2",
		},
		Avatar{
			Name:    "Hong Thai",
			Assets:  []string{"https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/avatar%2Ftest%2FgreenCircle.gif?alt=media&token=d2f01a05-eb06-4bfc-8227-fd93a6f85044"},
			Preview: "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/avatar%2Ftest%2FgreenCircle.gif?alt=media&token=d2f01a05-eb06-4bfc-8227-fd93a6f85044",
			Price:   5,
			Type:    "C_3",
		},
	}
	// Insert the document into the specified collection
	_, err := collection.InsertMany(context.TODO(), avatar)
	if err != nil {
		return
	}
	// Find the document
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var result []Avatar
	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}

	fmt.Printf("Document Found:\n%+v\n", result)

}
func InsertThemeBase(db *mongo.Database) {
	// Reference the database and collection to use
	collection := db.Collection("theme_shops")
	// Create new documents
	theme := []interface{}{
		Theme{
			Name:   "Suspicious Town",
			Assets: "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/theme%2Ftown.jpg?alt=media&token=2ed7270e-80a9-4cf0-9993-5b040f8018b2",
			Price:  123,
		},
		Theme{
			Name:   "Bed Room",
			Assets: "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/theme%2Fbedroom.jpg?alt=media&token=9c20af8b-6fa0-4c23-845a-342af463f264",
			Price:  123,
		},
		Theme{
			Name:   "Warm House",
			Assets: "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/theme%2Fponyo006.jpg?alt=media&token=3b8be510-e6e3-4c87-9371-5ff6ceb6a040",
			Price:  300,
		},
		Theme{
			Name:   "52hzWhale",
			Assets: "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/theme%2F4.png?alt=media&token=3af8b589-8ef8-46a9-83a0-a7299860ee97",
			Price:  300,
		},
	}
	// Insert the document into the specified collection
	_, err := collection.InsertMany(context.TODO(), theme)
	if err != nil {
		return
	}
	// Find the document
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var result []Avatar
	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}

	fmt.Printf("Document Found:\n%+v\n", result)

}
func InsertBackgroundBase(db *mongo.Database) {
	// Reference the database and collection to use
	collection := db.Collection("bg_shops")
	// Create new documents
	bg := []interface{}{
		Background{
			Name:   "Default",
			Assets: "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/background%2FbgShop.png?alt=media&token=2811d8e3-6ceb-4a41-ad2d-94a511ab9cb9",
			Price:  0,
		},
		Background{
			Name:   "Window Flower",
			Assets: "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/background%2F11.jpg?alt=media&token=c07205f2-e787-466d-b868-7fb7c8cc56e5",
			Price:  200,
		},
		Background{
			Name:   "Flower Land",
			Assets: "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/background%2F12.png?alt=media&token=3cb58838-5290-4db7-b820-cd51aa6d3466",
			Price:  200,
		},
		Background{
			Name:   "WindowXP",
			Assets: "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/background%2F13.png?alt=media&token=1038d5d5-1456-4c08-8580-a6a7f2abbd14",
			Price:  200,
		},
		Background{
			Name:   "Wink Wink",
			Assets: "https://firebasestorage.googleapis.com/v0/b/meetme-1815f.appspot.com/o/background%2F8.png?alt=media&token=71bd4617-a3d4-465f-8a34-1a0b52fe44cc",
			Price:  200,
		},
	}
	// Insert the document into the specified collection
	_, err := collection.InsertMany(context.TODO(), bg)
	if err != nil {
		return
	}
	// Find the document
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var result []Avatar
	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}

	fmt.Printf("Document Found:\n%+v\n", result)

}
