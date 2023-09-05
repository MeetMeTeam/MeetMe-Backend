package interfaces

type User struct {
	ID        int    `db:"id"`
	Firstname string `db:"firstname"`
	Lastname  string `db:"lastname"`
	Birthday  string `db:"birthday"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Image     string `db:"image"`
	Username  string `db:"username"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByEmail(string) (*User, error)
	GetById(int) (*User, error)
	Create(User) (*User, error)
	AddFriend() (*User, error)
	// UpdateTotalPoint(int, string) (*User, error)
}
