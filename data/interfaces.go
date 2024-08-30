package data

import "github.com/Ahmad-mufied/go-digilib/constants"

type UserInterfaces interface {
	CreateUser(user *User) (uint, uint, error)
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
	UpdateBookStock(bookID uint, stock int) error
	CheckBookById(bookID uint) error
}

type WalletInterfaces interface {
	GetWalletByUserID(userID uint) (*Wallet, error)
	GetWalletIdByUserID(userID uint) (uint, error)
}

type DepositInterfaces interface {
	CreateDeposit(deposit *Deposit) (uint, error)
	GetAllDepositsByWalletID(walletID uint) ([]Deposit, error)
	GetDepositById(depositID uint) (*Deposit, error)
	UpdateDepositStatus(invoiceUrl string, status constants.PaymentStatusEnum) error
	UpdateDepositInvoiceURL(depositID uint, invoiceURL string) error
}
