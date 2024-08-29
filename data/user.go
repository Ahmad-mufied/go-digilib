package data

import "github.com/Ahmad-mufied/go-digilib/constants"

type User struct {
	ID        uint                      `db:"id"`
	FullName  string                    `db:"full_name"`
	Username  string                    `db:"username"`
	Email     string                    `db:"email"`
	Password  string                    `db:"password"`
	Status    constants.UsersStatusEnum `db:"status"`
	Role      constants.UserRoleEnum    `db:"role"`
	BookCount int                       `db:"book_count"`
}

func (u *User) CreateUser(user *User) (uint, error) {

	sqlStatement := `INSERT INTO users (full_name, username, email, password, status, role, book_count) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	lastInsertID := 0
	// Insert user data to database
	err := db.QueryRow(sqlStatement, user.FullName, user.Username, user.Email, user.Password, user.Status, user.Role, user.BookCount).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return uint(lastInsertID), nil
}

func (u *User) GetUserById(userId uint) (*User, error) {

	sqlStatement := `SELECT id, full_name, username, email, password, status, role, book_count FROM users WHERE id = $1`

	// Get user data from database
	var user User
	err := db.Get(&user, sqlStatement, userId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) CheckUserId(userId uint) (bool, error) {

	sqlStatement := `SELECT COUNT(id) FROM users WHERE id = $1`

	// Check if user ID exists
	var count int
	err := db.Get(&count, sqlStatement, userId)
	if err != nil {
		return false, err
	}

	return count > 0, nil

}

func (u *User) CheckEmail(email string) (bool, error) {

	sqlStatement := `SELECT COUNT(email) FROM users WHERE email = $1`

	// Check if email is already registered
	var count int
	err := db.Get(&count, sqlStatement, email)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (u *User) GetUserByEmail(email string) (*User, error) {

	sqlStatement := `SELECT id, full_name, username, email, password, status, role, book_count FROM users WHERE email = $1`

	// Get user data from database
	var user User
	err := db.Get(&user, sqlStatement, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) GetPasswordByEmail(email string) (string, error) {

	sqlStatement := `SELECT password FROM users WHERE email = $1`

	// Get user password from database
	var password string
	err := db.Get(&password, sqlStatement, email)
	if err != nil {
		return "", err
	}

	return password, nil
}
