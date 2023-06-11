package data

type URL struct {
	ID      int
	LongVer string
	SrtVer  string
	UserID  int
}

type User struct {
	ID       int
	Email    string
	password string
}

type URLModel interface {
	GetURLByID(id int) (*URL, error)
	GetAllURLOfUser(id int) ([]*URL, error)
	CreateURL(url string) (*URL, error)
	UpdateURLByID(id int) error
	DeleteURLByID(id int) error
}

type UserModel interface {
	GetUserByID(id int) (*User, error)
	CreateUser(email, password string) (*User, error)
	DeleteUserByID(id int) error
	UpdateUser(id int) error
}
