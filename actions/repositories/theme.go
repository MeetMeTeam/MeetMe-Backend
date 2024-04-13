package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"meetme/be/actions/repositories/interfaces"
)

type ThemeRepository struct {
	db *mongo.Database
}

func NewThemeRepositoryDB(db *mongo.Database) ThemeRepository {
	return ThemeRepository{db: db}
}

func (r ThemeRepository) CreateTheme(request interfaces.Theme) (*interfaces.Theme, error) {

	_, err := r.db.Collection("theme_shops").InsertOne(context.TODO(), request)

	if err != nil {
		return nil, err
	}

	return &request, nil
}

func (r ThemeRepository) GetAllTheme() ([]interfaces.ThemeResponse, error) {
	filter := bson.D{}
	coll := r.db.Collection("theme_shops")
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var themes []interfaces.ThemeResponse
	if err = cursor.All(context.TODO(), &themes); err != nil {
		panic(err)
	}

	return themes, nil
}

func (r ThemeRepository) GetThemeById(id primitive.ObjectID) (*interfaces.ThemeResponse, error) {
	var theme interfaces.ThemeResponse
	filter := bson.D{{"_id", id}}
	coll := r.db.Collection("theme_shops")
	err := coll.FindOne(context.TODO(), filter).Decode(&theme)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		panic(err)
	}

	return &theme, nil
}
