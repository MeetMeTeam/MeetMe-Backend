package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"meetme/be/actions/repositories/interfaces"
)

type Question struct {
	Eng      string             `bson:"eng"`
	Thai     string             `bson:"thai"`
	Category primitive.ObjectID `bson:"category_id"`
}
type QuestionCategories struct {
	Name string `bson:"name"`
}

func InsertQuestionCategoryBase(db *mongo.Database) {
	// Reference the database and collection to use
	collection := db.Collection("question_categories")
	// Create new documents
	question := []interface{}{
		QuestionCategories{
			Name: "General",
		},
		QuestionCategories{
			Name: "Love",
		},
		QuestionCategories{
			Name: "Hobbies",
		},
		QuestionCategories{
			Name: "Deep Talk",
		},
	}
	// Insert the document into the specified collection
	_, err := collection.InsertMany(context.TODO(), question)
	if err != nil {
		return
	}
	// Find the document
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var result []QuestionCategories
	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}

	fmt.Printf("Document Found:\n%+v\n", result)

}
func InsertGeneralQuestionBase(db *mongo.Database) {
	var qc interfaces.CategoryResponse
	filter := bson.D{{"name", "General"}}
	coll := db.Collection("question_categories")
	err := coll.FindOne(context.TODO(), filter).Decode(&qc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			panic(err)
		}
		panic(err)
	}
	// Reference the database and collection to use
	collection := db.Collection("questions")
	// Create new documents
	question := []interface{}{
		Question{
			Eng:      "Who are You",
			Thai:     "คุณเป็นใคร",
			Category: qc.ID,
		},
		Question{
			Eng:      "What is your motto?",
			Thai:     "คติประจำตัวของคุณคืออะไร?",
			Category: qc.ID,
		},
		Question{
			Eng:      "Which season do you like?",
			Thai:     "คุณชอบฤดูกาลอะไร??",
			Category: qc.ID,
		},
	}
	// Insert the document into the specified collection
	collection.InsertMany(context.TODO(), question)
	if err != nil {
		return
	}
	// Find the document
	filter = bson.D{{"category_id", qc.ID}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var result []Question
	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}

	fmt.Printf("Document Found:\n%+v\n", result)
}

func InsertLoveQuestionBase(db *mongo.Database) {
	var qc interfaces.CategoryResponse
	filter := bson.D{{"name", "Love"}}
	coll := db.Collection("question_categories")
	err := coll.FindOne(context.TODO(), filter).Decode(&qc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			panic(err)
		}
		panic(err)
	}
	// Reference the database and collection to use
	collection := db.Collection("questions")
	// Create new documents
	question := []interface{}{
		Question{
			Eng:      "Which person you like?",
			Thai:     "ใครคือคนที่คุณชอบ?",
			Category: qc.ID,
		},
		Question{
			Eng:      "What is your favourite moment together last year",
			Thai:     "ในปีที่แล้ว ช่วงเวลาที่ใช้ร่วมกันแล้วคุณชื่นชอบคืออะไร?",
			Category: qc.ID,
		},
		Question{
			Eng:      "One thing you want to do together?",
			Thai:     "1 สิ่งที่คุณต้องการจะทำร่วมกันคืืออะไร?",
			Category: qc.ID,
		},
	}
	// Insert the document into the specified collection
	collection.InsertMany(context.TODO(), question)
	if err != nil {
		return
	}
	// Find the document
	filter = bson.D{{"category_id", qc.ID}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var result []Question
	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}

	fmt.Printf("Document Found:\n%+v\n", result)
}
func InsertHobbiesQuestionBase(db *mongo.Database) {
	var qc interfaces.CategoryResponse
	filter := bson.D{{"name", "Hobbies"}}
	coll := db.Collection("question_categories")
	err := coll.FindOne(context.TODO(), filter).Decode(&qc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			panic(err)
		}
		panic(err)
	}
	// Reference the database and collection to use
	collection := db.Collection("questions")
	// Create new documents
	question := []interface{}{
		Question{
			Eng:      "What is your favourite hobby?",
			Thai:     "งานอดิเรกที่คุณชื่นชอบคืออะไร?",
			Category: qc.ID,
		},
		Question{
			Eng:      "What is your favorite sport or physical activity?",
			Thai:     "อะไรคือกีฬาหรือการออกกำลังกายที่คุณชื่นชอบ?",
			Category: qc.ID,
		},
		Question{
			Eng:      "What is your idea of fun?",
			Thai:     "ความสนุกของคุณคืออะไร?",
			Category: qc.ID,
		},
	}
	// Insert the document into the specified collection
	collection.InsertMany(context.TODO(), question)
	if err != nil {
		return
	}
	// Find the document
	filter = bson.D{{"category_id", qc.ID}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var result []Question
	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}

	fmt.Printf("Document Found:\n%+v\n", result)
}

func InsertDeepTalkQuestionBase(db *mongo.Database) {
	var qc interfaces.CategoryResponse
	filter := bson.D{{"name", "Deep Talk"}}
	coll := db.Collection("question_categories")
	err := coll.FindOne(context.TODO(), filter).Decode(&qc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			panic(err)
		}
		panic(err)
	}
	// Reference the database and collection to use
	collection := db.Collection("questions")
	// Create new documents
	question := []interface{}{
		Question{
			Eng:      "What is the story of your name?",
			Thai:     "ชื่อของคุณมีที่มาว่าอย่างไร?",
			Category: qc.ID,
		},
		Question{
			Eng:      "What did you like most about the last week?",
			Thai:     "อาทิตย์ที่แล้วชอบอะไรมากที่สุด?",
			Category: qc.ID,
		},
		Question{
			Eng:      "Your own habits that you like/dislike",
			Thai:     "นิสัยตัวเองที่ชอบ/ไม่ชอบ",
			Category: qc.ID,
		},
	}
	// Insert the document into the specified collection
	collection.InsertMany(context.TODO(), question)
	if err != nil {
		return
	}
	// Find the document
	filter = bson.D{{"category_id", qc.ID}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var result []Question
	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}

	fmt.Printf("Document Found:\n%+v\n", result)
}
