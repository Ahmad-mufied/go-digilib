package data

type UserInterfaces interface {
	CreateUser(user *User) (uint, error)
	GetUserById(userId uint) (*User, error)
	CheckUserId(userId uint) (bool, error)
	CheckEmail(email string) (bool, error)
	GetUserByEmail(email string) (*User, error)
	GetPasswordByEmail(email string) (string, error)
}

type BookInterfaces interface {
	CreateBook(book *Book, physicalDetails *BookPhysicalDetails) (uint, error)
	GetStockById(bookID uint) (int, error)
	CheckBookBySKU(sku string) (bool, error)
	GetBookById(bookID uint) (*Book, error)
	GetDetailBookById(bookID uint) (*Book, *BookPhysicalDetails, error)
	GetAllBooks() ([]Book, error)
}
