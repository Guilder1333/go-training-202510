package dal

type UserRepository interface {
	CheckUserById(id int) (bool, error)
	DeleteUser(id int) error
}
