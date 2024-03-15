package interfaces

type FavoriteService interface {
	FavUser(string, string) (interface{}, error)
	UnFavUser(string, string) (interface{}, error)
}
