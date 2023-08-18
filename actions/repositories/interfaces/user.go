package interfaces

type User struct {
	ID        int    `db:"id"`
	Firstname string `db:"firstname"`
	Lastname  string `db:"lastname"`
	Birthday  string `db:"birthday"`
	Email     string `db:"email"`
	Password  string `db:"password"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByEmail(string) (*User, error)
	Create(User) (*User, error)
	// UpdateTotalPoint(int, string) (*User, error)
	// Update(User, string) (*User, error)
}
