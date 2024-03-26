package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"meetme/be/actions/repositories/interfaces"
)

type QuestionRepository struct {
	db *mongo.Database
}

func NewQuestionRepositoryDB(db *mongo.Database) QuestionRepository {
	return QuestionRepository{db: db}
}

func (r QuestionRepository) GetAll() ([]interfaces.QuestionResponse, error) {
	filter := bson.D{}
	coll := r.db.Collection("questions")
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var questions []interfaces.QuestionResponse
	if err = cursor.All(context.TODO(), &questions); err != nil {
		panic(err)
	}

	return questions, nil
}

func (r QuestionRepository) GetCategoryAll() ([]interfaces.CategoryResponse, error) {
	filter := bson.D{}
	coll := r.db.Collection("question_categories")
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var cate []interfaces.CategoryResponse
	if err = cursor.All(context.TODO(), &cate); err != nil {
		panic(err)
	}

	return cate, nil
}

func (r QuestionRepository) GetByCategory(cate primitive.ObjectID) ([]interfaces.QuestionResponse, error) {
	filter := bson.D{{"category_id", cate}}
	coll := r.db.Collection("questions")
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var questions []interfaces.QuestionResponse
	if err = cursor.All(context.TODO(), &questions); err != nil {
		panic(err)
	}

	return questions, nil
}
func (r QuestionRepository) GetCateById(id primitive.ObjectID) (*interfaces.CategoryResponse, error) {
	var cate interfaces.CategoryResponse
	filter := bson.D{{"_id", id}}
	coll := r.db.Collection("question_categories")
	err := coll.FindOne(context.TODO(), filter).Decode(&cate)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		panic(err)
	}

	return &cate, nil
}

func (r QuestionRepository) CreateQuestions(question interfaces.Question) (*interfaces.Question, error) {

	_, err := r.db.Collection("questions").InsertOne(context.TODO(), question)

	if err != nil {
		return nil, err
	}

	return &question, nil
}
