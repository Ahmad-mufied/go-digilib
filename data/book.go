package data

import (
	"github.com/Ahmad-mufied/go-digilib/constants"
	"github.com/jmoiron/sqlx/types"
	"github.com/pkg/errors"
	"time"
)

// Book implement Books table
type Book struct {
	ID            uint                     `db:"id"`
	CategoryID    uint                     `db:"category_id"`
	ISBN          string                   `db:"isbn"`
	SKU           string                   `db:"sku"`
	Author        types.JSONText           `db:"author"`
	Title         string                   `db:"title"`
	Image         string                   `db:"image"`
	Pages         int                      `db:"pages"`
	Language      string                   `db:"language"`
	Description   string                   `db:"description"`
	Stock         int                      `db:"stock"`
	Status        constants.BookStatusEnum `db:"status"`
	BorrowedCount int                      `db:"borrowed_count"`
	PublishedAt   *time.Time               `db:"published_at"`
	BasePrice     float64                  `db:"base_price"`
	CreatedAt     time.Time                `db:"created_at"`
	UpdatedAt     time.Time                `db:"updated_at"`
}

type BookPhysicalDetails struct {
	BookID int     `db:"book_id"`
	Weight float64 `db:"weight"`
	Height int     `db:"height"`
	Width  int     `db:"width"`
}

func (b *Book) CreateBook(book *Book, physicalDetails *BookPhysicalDetails) (uint, error) {

	tx := db.MustBegin()

	sqlStatement := `INSERT INTO books (category_id, isbn, sku, author, title, image, pages, language, description, stock, status, borrowed_count, published_at, base_price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, NOW(), NOW()) RETURNING id`

	lastInsertID := 0
	// Insert book data to database
	err := tx.QueryRow(sqlStatement, book.CategoryID, book.ISBN, book.SKU, book.Author, book.Title, book.Image, book.Pages, book.Language, book.Description, book.Stock, book.Status, book.BorrowedCount, book.PublishedAt, book.BasePrice).Scan(&lastInsertID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Insert book physical details to database
	_, err = tx.Exec(`INSERT INTO book_physical_details (book_id, weight, height, width) VALUES ($1, $2, $3, $4)`, lastInsertID, physicalDetails.Weight, physicalDetails.Height, physicalDetails.Width)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return uint(lastInsertID), nil

}

func (b *Book) GetStockById(bookID uint) (int, error) {

	var stock int
	err := db.Get(&stock, "SELECT stock FROM books WHERE id = $1", bookID)
	if err != nil {
		return 0, err
	}

	return stock, nil
}

func (b *Book) CheckBookBySKU(sku string) (bool, error) {

	var count int
	err := db.Get(&count, "SELECT COUNT(id) FROM books WHERE sku = $1", sku)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (b *Book) GetBookById(bookID uint) (*Book, error) {

	var book Book
	err := db.Get(&book, "SELECT id, category_id, isbn, sku, author, title, image, pages, language, description, stock, status, borrowed_count, published_at, base_price, created_at, updated_at FROM books WHERE id = $1", bookID)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (b *Book) GetDetailBookById(bookID uint) (*Book, *BookPhysicalDetails, error) {

	// Get book data
	book, err := b.GetBookById(bookID)
	if err != nil {
		return nil, nil, errors.New("Book not found")
	}

	// Get book physical details data
	var physicalDetails BookPhysicalDetails
	err = db.Get(&physicalDetails, "SELECT book_id, weight, height, width FROM book_physical_details WHERE book_id = $1", bookID)
	if err != nil {
		return nil, nil, err
	}

	return book, &physicalDetails, nil
}

func (b *Book) GetAllBooks() ([]Book, error) {

	var books []Book
	sqlStatement := `SELECT id, category_id, isbn, sku, author, title, image, pages, language, description, stock, status, borrowed_count, published_at, base_price, created_at, updated_at FROM books`
	err := db.Select(&books, sqlStatement)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *Book) UpdateBookStock(bookID uint, stock int) error {

	tx := db.MustBegin()

	_, err := tx.Exec(`UPDATE books SET stock = $1 WHERE id = $2`, stock, bookID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (b *Book) CheckBookById(bookID uint) error {

	count := 0
	err := db.Get(&count, "SELECT 1 FROM books WHERE id = $1", bookID)
	if err != nil {
		return err
	}

	return nil
}

func (b *Book) GetBookBasePrice(bookID uint) (float64, error) {

	var basePrice float64
	err := db.Get(&basePrice, "SELECT base_price FROM books WHERE id = $1", bookID)
	if err != nil {
		return 0, err
	}

	return basePrice, nil
}
