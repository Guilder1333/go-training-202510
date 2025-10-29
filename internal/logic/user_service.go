package logic

type User struct {
	Id              int
	FirstName       string
	LastName        string
	Age             int
	PhoneNumber     string
	IsPhoneVerified bool
}

type UserService interface {
	GetUserById(id int) (*User, error)
	CreateUser(user *User) (int, error)
	DeleteUser(id int) error
}
