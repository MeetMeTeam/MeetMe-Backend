package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
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
