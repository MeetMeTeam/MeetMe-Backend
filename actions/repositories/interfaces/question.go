package interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

type QuestionResponse struct {
	ID       primitive.ObjectID `bson:"_id"`
	Eng      string             `bson:"eng"`
	Thai     string             `bson:"thai"`
	Category primitive.ObjectID `bson:"category_id"`
}
type Question struct {
	Eng      string             `bson:"eng"`
	Thai     string             `bson:"thai"`
	Category primitive.ObjectID `bson:"category_id"`
}
type CategoryResponse struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}
type QuestionRepository interface {
	GetAll() ([]QuestionResponse, error)
	GetCategoryAll() ([]CategoryResponse, error)
	GetByCategory(primitive.ObjectID) ([]QuestionResponse, error)
	GetCateById(primitive.ObjectID) (*CategoryResponse, error)
	CreateQuestions(string, Question) (*Question, error)
}
