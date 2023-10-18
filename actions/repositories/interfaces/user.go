package interfaces

type User struct {
	ID        int    `bson:"id"`
	Firstname string `bson:"firstname"`
	Lastname  string `bson:"lastname"`
	Birthday  string `bson:"birthday"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
	Image     string `bson:"image"`
	Username  string `bson:"username"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByEmail(string) (*User, error)
	GetById(int) (*User, error)
	Create(User) (*User, error)
	AddFriend() (*User, error)
	// UpdateTotalPoint(int, string) (*User, error)
}
