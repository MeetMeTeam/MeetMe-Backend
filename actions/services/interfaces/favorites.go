package interfaces

type FavoriteService interface {
	FavUser(string, string) (interface{}, error)
}
