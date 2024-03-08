package repositories

import (
	"context"
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
