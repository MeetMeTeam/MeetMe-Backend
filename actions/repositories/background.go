package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"meetme/be/actions/repositories/interfaces"
)

type BgRepository struct {
	db *mongo.Database
}

func NewBgRepositoryDB(db *mongo.Database) BgRepository {
	return BgRepository{db: db}
}

func (r BgRepository) Create(request interfaces.Background) (*interfaces.Background, error) {

	_, err := r.db.Collection("bg_shops").InsertOne(context.TODO(), request)

	if err != nil {
		return nil, err
	}

	return &request, nil
}
