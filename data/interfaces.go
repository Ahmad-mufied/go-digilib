package data

type UserInterfaces interface {
	CreateUser(user *User) (uint, error)
	GetUserById(userId uint) (*User, error)
	CheckUserId(userId uint) (bool, error)
	CheckEmail(email string) (bool, error)
	GetUserByEmail(email string) (*User, error)
	GetPasswordByEmail(email string) (string, error)
}
