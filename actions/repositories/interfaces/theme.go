package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ThemeResponse struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	Assets string             `bson:"assets"`
	Price  int                `bson:"price"`
	Song   string             `bson:"song"`
}
type Theme struct {
	Name   string `bson:"name"`
	Assets string `bson:"assets"`
	Price  int    `bson:"price"`
	Song   string `bson:"song"`
}
type ThemeRepository interface {
	GetThemeById(primitive.ObjectID) (*ThemeResponse, error)
	GetAllTheme() ([]ThemeResponse, error)
	CreateTheme(Theme) (*Theme, error)
	//GetByType(string) ([]AvatarResponse, error)
}
